package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	gopatterns "github.com/pseudoelement/go-sandbox/go-patterns"
)

const WORKERS_COUNT int = 3_000

var sem *gopatterns.Semaphore = gopatterns.NewSemaphore(context.TODO(), 20)

/**
 * 3000 workers: avg 3834.75
 */
func TestConcurrentUpdatesUnlocked() {
	db := runDB()
	errChan := make(chan error, WORKERS_COUNT)
	wg := sync.WaitGroup{}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	now := time.Now()
	for range WORKERS_COUNT {
		wg.Add(1)
		go func() {
			if err := _sendUpdateQuery(db); err != nil {
				errChan <- err
			}
			wg.Done()
		}()
	}

	for err := range errChan {
		log.Println("[errChan] error:", err)
	}
	log.Println("Time: ", time.Since(now).Milliseconds())
}

/**
 * 3000 workers: avg 5049.25
 */
func TestConcurrentUpdatesRowLocked() {
	db := runDB()
	errChan := make(chan error, WORKERS_COUNT)
	wg := sync.WaitGroup{}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	now := time.Now()
	for range WORKERS_COUNT {
		wg.Add(1)
		go func() {
			if err := _sendUpdateQueryRowLocked(db); err != nil {
				errChan <- err
			}
			wg.Done()
		}()
	}

	for err := range errChan {
		log.Println("[errChan] error:", err)
	}
	log.Println("Time: ", time.Since(now).Milliseconds())
}

/**
 * 3000 workers: avg 3944
 */
func TestConcurrentUpdatesTableLocked() {
	db := runDB()
	errChan := make(chan error, WORKERS_COUNT)
	wg := sync.WaitGroup{}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	now := time.Now()
	for range WORKERS_COUNT {
		wg.Add(1)
		go func() {
			if err := _sendUpdateQueryAccessExclusive(db); err != nil {
				errChan <- err
			}
			wg.Done()
		}()
	}

	for err := range errChan {
		log.Println("[errChan] error:", err)
	}
	log.Println("Time: ", time.Since(now).Milliseconds())
}

/**
 * 3000 workers: avg 4129.5
 */
func TestConcurrentUpdatesMutexLocked() {
	db := runDB()
	errChan := make(chan error, WORKERS_COUNT)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	now := time.Now()
	for range WORKERS_COUNT {
		wg.Add(1)
		go func() {
			mu.Lock()
			err := _sendUpdateQuery(db)
			mu.Unlock()
			if err != nil {
				errChan <- err
			}
			wg.Done()
		}()
	}

	for err := range errChan {
		log.Println("[errChan] error:", err)
	}
	log.Println("Time: ", time.Since(now).Milliseconds())
}

func runDB() *sql.DB {
	err := godotenv.Load("./db/.env")
	if err != nil {
		panic("[godotenv.Load] err: " + err.Error())
	}

	host, port, password, user, dbname := os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME")
	if host == "" || port == "" || password == "" || user == "" || dbname == "" {
		panic("Invalid props in .env file")
	}

	log.Println(".env vars: ", host, port, password, user, dbname)

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(50)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Postgres successfully connected.")

	return db
}

func _sendUpdateQuery(db *sql.DB) error {
	id := 8
	row, err := NewQueryBuilder[SchemaPayment](db).
		Begin().
		Query("UPDATE payments SET amount = amount + 1 WHERE id = $1;", id).
		QueryRow("SELECT * from payments WHERE id = $1;", id).
		ResultRow(0)
	if err != nil {
		return err
	}

	log.Println("Tx done. Row is ", row)
	return nil
}

func _sendUpdateQueryRowLocked(db *sql.DB) error {
	id := 8
	row, err := NewQueryBuilder[SchemaPayment](db).
		Begin().
		Query("SELECT amount FROM payments WHERE id = $1 FOR UPDATE;", id).
		Query("UPDATE payments SET amount = amount + 1 WHERE id = $1;", id).
		QueryRow("SELECT * from payments WHERE id = $1;", id).
		ResultRow(0)
	if err != nil {
		return err
	}

	log.Println("Tx done. Row is ", row)
	return nil
}

func _sendUpdateQueryAccessExclusive(db *sql.DB) error {
	id := 8
	row, err := NewQueryBuilder[SchemaPayment](db).
		Begin().
		Query("LOCK TABLE payments IN ACCESS EXCLUSIVE MODE;").
		Query("UPDATE payments SET amount = amount + 1 WHERE id = $1;", id).
		QueryRow("SELECT * from payments WHERE id = $1;", id).
		ResultRow(0)
	if err != nil {
		return err
	}

	log.Println("Tx done. Row is ", row)
	return nil
}

// func sendUpdateQuery(db *sql.DB) error {
// 	tx, err := db.BeginTx(
// 		context.TODO(),
// 		&sql.TxOptions{Isolation: sql.LevelReadCommitted},
// 	)
// 	if err != nil {
// 		return fmt.Errorf("[sendUpdateQuery] BeginTx err: %w", err)
// 	}
// 	id := 2
// 	_, execErr := tx.Exec(`UPDATE payments SET amount = amount + 1 WHERE id = $1;`, id)
// 	if execErr != nil {
// 		_ = tx.Rollback()
// 		return fmt.Errorf("[sendUpdateQuery] tx.Exec err: %w", execErr)
// 	}

// 	var amount int
// 	err = tx.QueryRow("SELECT amount from payments WHERE id = $1;", id).Scan(&amount)
// 	if err != nil {
// 		_ = tx.Rollback()
// 		return fmt.Errorf("[sendUpdateQuery] tx.QueryRow err: %w", err)
// 	}
// 	if err := tx.Commit(); err != nil {
// 		_ = tx.Rollback()
// 		return fmt.Errorf("[sendUpdateQuery] tx.Commit err: %w", err)
// 	}

// 	log.Println("Tx done. Updated amount is ", amount)
// 	return nil
// }

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func TestConcurrentUpdatesUnlocked() {
	db := runDB()
	sendUpdateQueryUnlocked(db)
	// workersCount := 300
	// errChan := make(chan error, workersCount)
	// wg := sync.WaitGroup{}

	// go func() {
	// 	wg.Wait()
	// 	close(errChan)
	// }()

	// for range workersCount {
	// 	wg.Add(1)
	// 	go func() {
	// 		if err := sendUpdateQuery(db); err != nil {
	// 			errChan <- err
	// 		}
	// 		wg.Done()
	// 	}()
	// }

	// for err := range errChan {
	// 	log.Println("[errChan] error:", err)
	// }
}

func TestConcurrentUpdatesRowLocked() {}

func TestConcurrentUpdatesTableLocked() {}

func TestConcurrentUpdatesMutexLocked() {
	db := runDB()
	workersCount := 1000
	errChan := make(chan error, workersCount)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for range workersCount {
		wg.Add(1)
		go func() {
			mu.Lock()
			err := sendUpdateQueryUnlocked(db)
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

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Postgres successfully connected.")

	return db
}

func sendUpdateQueryUnlocked(db *sql.DB) error {
	id := 8
	res, err := NewQueryBuilder[SchemaPayment](db).
		Begin().
		Query("UPDATE payments SET amount = amount + 1 WHERE id = $1;", id).
		QueryRow("SELECT * from payments WHERE id = $1;", id).
		QueryRows("SELECT * from payments;").
		Results()

	if err != nil {
		log.Println("ERROR: ==>", err)
		return err
	}

	log.Println("Tx done. Results is ", res)
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

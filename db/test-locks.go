package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func TestConcurrentUpdatesUnlocked() {
	db := runDB()
	var paymentRow SchemaPayment
	err := db.QueryRow("SELECT * FROM payments WHERE id=$1", 3).Scan(
		&paymentRow.PaymentId,
		&paymentRow.PersonId,
		&paymentRow.CreatedAt,
		&paymentRow.Amount)
	if err != nil {
		panic("[TestConcurrentUpdatesUnlocked] QueryRow err: " + err.Error())
	}

	fmt.Printf("Payment row: %+v \n", paymentRow)
}

func TestConcurrentUpdatesRowLocked() {}

func TestConcurrentUpdatesTableLocked() {}

func TestConcurrentUpdatesMutexLocked() {}

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

func sendUpdateQuery(db *sql.DB) {
	tx, err := db.BeginTx(
		context.TODO(),
		&sql.TxOptions{Isolation: sql.LevelReadUncommitted},
	)
	if err != nil {
		log.Fatal("[sendConcurrentUpdates] BeginTx err: " + err.Error())
	}
	id := 37
	_, execErr := tx.Exec(`UPDATE users SET status = ? WHERE id = ?`, "paid", id)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

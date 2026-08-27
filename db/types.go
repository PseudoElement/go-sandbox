package db

import "time"

type SchemaPayment struct {
	PaymentId int       `json:"id"`
	PersonId  int       `json:"person_id"`
	CreatedAt time.Time `json:"created_at"`
	Amount    int       `json:"amount"`
}

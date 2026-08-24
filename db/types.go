package db

import "time"

type SchemaPayment struct {
	PaymentId int       `json:"payment_id"`
	PersonId  int       `json:"person_id"`
	CreatedAt time.Time `json:"created_at"`
	Amount    int       `json:"amount"`
}

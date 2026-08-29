package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrNotInitialized = errors.New("db.BeginTx should be called first.")
	ErrOutOfBound     = errors.New("idx is out of results size.")
)

type QueryBuilder[T any] struct {
	db               *sql.DB
	tx               *sql.Tx
	err              error
	resultsQueryRow  []T
	resultsQueryRows [][]T
	commited         bool
}

func NewQueryBuilder[T any](db *sql.DB) *QueryBuilder[T] {
	return &QueryBuilder[T]{
		db:               db,
		tx:               nil,
		err:              nil,
		resultsQueryRow:  make([]T, 0),
		resultsQueryRows: make([][]T, 0),
		commited:         false,
	}
}

func (qb *QueryBuilder[T]) ResultRows(idx uint8) ([]T, error) {
	if qb.err != nil {
		return []T{}, qb.err
	}
	if int(idx) > len(qb.resultsQueryRows)-1 {
		return []T{}, ErrOutOfBound
	}
	if !qb.commited {
		qb.err = qb._commit()
		qb.commited = true
	}
	return qb.resultsQueryRows[idx], qb.err
}

func (qb *QueryBuilder[T]) ResultRow(idx uint8) (T, error) {
	if qb.err != nil {
		return *new(T), qb.err
	}
	if int(idx) > len(qb.resultsQueryRow)-1 {
		return *new(T), ErrOutOfBound
	}
	if !qb.commited {
		qb.err = qb._commit()
		qb.commited = true
	}
	return qb.resultsQueryRow[idx], qb.err
}

func (qb *QueryBuilder[T]) Begin() *QueryBuilder[T] {
	tx, err := qb.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	qb.err = err
	qb.tx = tx
	return qb
}

func (qb *QueryBuilder[T]) Query(query string, args ...any) *QueryBuilder[T] {
	qb._checkTxNonNil()
	if qb.err != nil {
		return qb
	}
	_, execErr := qb.tx.Exec(query, args...)
	if execErr != nil {
		qb.err = fmt.Errorf("[sendUpdateQuery] tx.Exec err: %w", execErr)
	}
	return qb
}

func (qb *QueryBuilder[T]) QueryRows(query string, args ...any) *QueryBuilder[T] {
	qb._checkTxNonNil()
	if qb.err != nil {
		return qb
	}
	results := qb._queryRows(
		func(results []T) bool { return false },
		query,
		args...,
	)
	qb.resultsQueryRows = append(qb.resultsQueryRows, results)

	return qb
}

func (qb *QueryBuilder[T]) QueryRow(query string, args ...any) *QueryBuilder[T] {
	qb._checkTxNonNil()
	if qb.err != nil {
		return qb
	}
	results := qb._queryRows(
		func(results []T) bool { return len(results) > 0 },
		query,
		args...,
	)
	if len(results) > 0 {
		qb.resultsQueryRow = append(qb.resultsQueryRow, results[0])
	}

	return qb
}

func (qb *QueryBuilder[T]) _commit() error {
	qb._checkTxNonNil()
	if qb.err != nil {
		if !errors.Is(qb.err, ErrNotInitialized) {
			_ = qb.tx.Rollback()
		}
		return qb.err
	}
	if err := qb.tx.Commit(); err != nil {
		return fmt.Errorf("[sendUpdateQuery] tx.Commit err: %w", err)
	}
	return nil
}

func (qb *QueryBuilder[T]) _queryRows(stopFn func([]T) bool, query string, args ...any) []T {
	rows, err := qb.tx.Query(query, args...)
	if err != nil {
		qb.err = err
		return []T{}
	}
	defer rows.Close()

	columnNames, err := rows.Columns()
	if err != nil {
		qb.err = err
		return []T{}
	}
	var res T
	structPtr := reflect.New(reflect.TypeOf(res))
	structVal := structPtr.Elem()
	results := make([]T, 0)
	for rows.Next() {
		pointers := make([]any, len(columnNames))
		for i := 0; i < structVal.NumField(); i++ {
			pointers[i] = structVal.Field(i).Addr().Interface()
		}
		if err := rows.Scan(pointers...); err != nil {
			qb.err = err
			return []T{}
		}
		res := structVal.Interface().(T)
		results = append(results, res)

		if stopFn(results) {
			break
		}
	}

	return results
}

func (qb *QueryBuilder[T]) _checkTxNonNil() {
	if qb.err == nil && qb.tx == nil {
		qb.err = ErrNotInitialized
	}
}

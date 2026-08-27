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
)

type QueryBuilder[T any] struct {
	db      *sql.DB
	tx      *sql.Tx
	err     error
	res     T
	results []T
}

func NewQueryBuilder[T any](db *sql.DB) *QueryBuilder[T] {
	return &QueryBuilder[T]{db: db, tx: nil, err: nil, results: make([]T, 0)}
}

func (qb *QueryBuilder[T]) Results() ([]T, error) {
	err := qb._commit()
	return qb.results, err
}

func (qb *QueryBuilder[T]) Result() (T, error) {
	err := qb._commit()
	return qb.res, err
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
	qb.results = results

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
	qb.res = results[0]

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
		return nil
	}
	defer rows.Close()

	columnNames, err := rows.Columns()
	if err != nil {
		qb.err = err
		return nil
	}
	structPtr := reflect.New(reflect.TypeOf(qb.res))
	structVal := structPtr.Elem()
	results := make([]T, 0)
	for rows.Next() {
		pointers := make([]any, len(columnNames))
		for i := 0; i < structVal.NumField(); i++ {
			pointers[i] = structVal.Field(i).Addr().Interface()
		}
		if err := rows.Scan(pointers...); err != nil {
			qb.err = err
			return nil
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
	if qb.tx == nil {
		qb.err = ErrNotInitialized
	}
}

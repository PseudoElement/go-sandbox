package db

import (
	"errors"
	"log"
	"math"
	"slices"
	"time"

	"github.com/lib/pq"
)

func toErrorWithCode(err error) (*ErrorWithCode, bool) {
	if errWithCode, ok := errors.AsType[*ErrorWithCode](err); ok {
		return errWithCode, true
	}
	if pqErr, ok := errors.AsType[*pq.Error](err); ok {
		return &ErrorWithCode{Code: string(pqErr.Code), Message: pqErr.Error()}, true
	}
	return nil, false
}

func retry(fn func() error, options RetryOptions) error {
	defaultDelay := options.DefaultDelayMs
	if options.DefaultDelayMs == 0 {
		defaultDelay = 100
	}

	var err error = nil
	defer func() {
		if err != nil && options.FallbackFn != nil {
			options.FallbackFn()
		}
	}()

	for count := range options.RetryCount {
		err = fn()
		if err == nil {
			return nil
		}
		errWithCode, ok := toErrorWithCode(err)
		if !ok {
			return err
		}
		contains := slices.ContainsFunc(options.ErrorCodesForRetry, func(errCode string) bool {
			return errCode == errWithCode.Code
		})
		if !contains {
			return err
		}
		log.Println("RETRY_", count)
		mul := math.Pow(float64(count+1), 2)
		time.Sleep(time.Millisecond * time.Duration(defaultDelay*int(mul)))
	}

	return err
}

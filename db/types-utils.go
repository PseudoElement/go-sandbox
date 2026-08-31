package db

type RetryOptions struct {
	RetryCount         int
	DefaultDelayMs     int
	FallbackFn         func() error
	ErrorCodesForRetry []string
}

type ErrorWithCode struct {
	Code    string
	Message string
}

func (e *ErrorWithCode) Error() string {
	return e.Message
}

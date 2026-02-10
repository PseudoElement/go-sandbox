package main

import (
	"fmt"
	"math/big"
	"strings"
	"time"
)

type HexString string

func (as *HexString) Value() string {
	return string(*as)
}

func (as *HexString) Valid() bool {
	return as.has0x() && len([]byte(as.Value())) == 42
}

func (as *HexString) has0x() bool {
	return strings.HasPrefix(as.Value(), "0x")
}

func (as *HexString) ToDecimal() (*big.Int, error) {
	if !as.Valid() {
		return nil, fmt.Errorf("%s is invalid hex string", as.Value())
	}

	hexStr := as.Value()[2:]
	decimalValue := new(big.Int)

	_, success := decimalValue.SetString(hexStr, 16)
	if !success {
		return nil, fmt.Errorf("failed to parse hex string: %s", hexStr)
	}

	return decimalValue, nil
}

func makeCall(value string) *HexString {
	println("Start call...")
	time.Sleep(1 * time.Second)
	println("End call.")

	hexStr := HexString(value)

	return &hexStr
}

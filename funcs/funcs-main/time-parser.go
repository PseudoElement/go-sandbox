package main

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	InvalidTokenType = errors.New("invalid token type")
	UnknownTimeUnit  = errors.New("unknown time unit")
)

type TokenType = int8

var (
	MonthValues    = [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	DayValues      = [7]string{"Mon", "Tue", "Wen", "Thi", "Fri", "Sat", "Sun"}
	TimeZoneValues = [2]string{"MST", "UTC"}
)

const (
	Sec TokenType = iota + 1
	Min
	Hour
	Day
	Month
	Year
	Timezone
	Separator
)

type Token struct {
	FromIdx  int
	ToIdx    int
	Type     int8
	Value    string
	TypeFunc func(char rune) bool
}

func tokenizeLayout(layout string) ([]Token, error) {
	tokens := []Token{}

	var tokenValue string = ""
	token := Token{}
	var prevUtf8Code rune = 0
	for idx, currUtf8Code := range layout {
		if idx == 0 || sameAsPrevious(currUtf8Code, prevUtf8Code) {
			tokenValue += string(currUtf8Code)
			if idx == 0 {
				fn, err := defineTokenTypeFunc(currUtf8Code)
				if err != nil {
					return []Token{}, err
				}
				token.TypeFunc = fn
			}
		} else {
			token.ToIdx = idx - 1
			tokenType, err := defineTokenType(tokenValue)
			if err != nil {
				return []Token{}, err
			}
			token.Value = tokenValue
			token.Type = tokenType
			tokens = append(tokens, token)
			// next token generation
			tokenValue = string(currUtf8Code)
			token = Token{FromIdx: idx}
			fn, err := defineTokenTypeFunc(currUtf8Code)
			if err != nil {
				return []Token{}, err
			}
			token.TypeFunc = fn
		}

		// handle last token in the end of layoutString
		if utf8.RuneCountInString(layout)-1 == idx {
			token.ToIdx = idx
			tokenType, err := defineTokenType(tokenValue)
			if err != nil {
				return []Token{}, err
			}
			token.Value = tokenValue
			token.Type = tokenType
			tokens = append(tokens, token)
		}

		prevUtf8Code = currUtf8Code
	}

	return tokens, nil
}

func parseTimeString(layout string, timeString string) (time.Time, error) {
	tokens, err := tokenizeLayout(layout)
	if err != nil {
		return time.Time{}, err
	}

	m := make(map[TokenType]any, len(tokens))
	var tokenIdx int
	var timeUnitValue string
	for idx, char := range timeString {
		token := tokens[tokenIdx]
		sameTimeUnit := token.TypeFunc(char)
		if sameTimeUnit {
			timeUnitValue += string(char)
		} else {
			/* handle previous accumulated timeUnitValue */
			if token.Type != Separator {
				timeUnit, tokenType, err := defineTimeUnit(token, timeUnitValue)
				if err != nil {
					return time.Time{}, err
				}
				m[tokenType] = timeUnit
			}
			/* start handling next */
			tokenIdx++
			timeUnitValue = string(char)
		}

		// handle last token in the end of layoutString
		if utf8.RuneCountInString(timeString)-1 == idx {
			timeUnit, tokenType, err := defineTimeUnit(token, timeUnitValue)
			if err != nil {
				return time.Time{}, err
			}
			m[tokenType] = timeUnit
		}
	}

	log.Println("MAP", m)
	t := time.Date(
		m[Year].(int),
		m[Month].(time.Month),
		m[Day].(int),
		m[Hour].(int),
		m[Min].(int),
		m[Sec].(int),
		1000,
		m[Timezone].(*time.Location),
	)

	return t, nil
}

func defineTimeUnit(token Token, timeUnitValue string) (any, TokenType, error) {
	switch token.Type {
	case Year:
		year, err := strconv.Atoi(timeUnitValue)
		if err != nil {
			return nil, 0, err
		}
		return year, Year, nil
	case Sec:
		sec, err := strconv.Atoi(timeUnitValue)
		if err != nil {
			return nil, 0, err
		}
		return sec, Sec, nil
	case Min:
		min, err := strconv.Atoi(timeUnitValue)
		if err != nil {
			return nil, 0, err
		}
		return min, Min, nil
	case Hour:
		hour, err := strconv.Atoi(timeUnitValue)
		if err != nil {
			return nil, 0, err
		}
		return hour, Hour, nil
	case Day:
		utf8Code, _ := utf8.DecodeRune([]byte{timeUnitValue[0]})
		if isNumeric(utf8Code) {
			day, err := strconv.Atoi(timeUnitValue)
			if err != nil {
				return nil, 0, err
			}
			return day, Day, nil
		} else {
			return dayToNumeric(timeUnitValue), Day, nil
		}
	case Month:
		utf8Code, _ := utf8.DecodeRune([]byte{timeUnitValue[0]})
		if isNumeric(utf8Code) {
			monthVal, err := strconv.Atoi(timeUnitValue)
			if err != nil {
				return nil, 0, err
			}
			return time.Month(monthVal), Month, nil
		} else {
			return time.Month(monthToNumeric(timeUnitValue)), Month, nil
		}
	case Timezone:
		tz, err := time.LoadLocation(timeUnitValue)
		if err != nil {
			return nil, 0, err
		}
		return tz, Timezone, nil
	default:
		return nil, 0, UnknownTimeUnit
	}
}

func dayToNumeric(day string) int {
	for idx, val := range DayValues {
		if day == val {
			return idx + 1
		}
	}
	return 0
}

func monthToNumeric(month string) int {
	for idx, val := range MonthValues {
		if month == val {
			return idx + 1
		}
	}
	return 0
}

func sameAsPrevious(prevUtf8Code rune, currUtf8Code rune) bool {
	return isAlpha(currUtf8Code) == isAlpha(prevUtf8Code) &&
		isNumeric(currUtf8Code) == isNumeric(prevUtf8Code) &&
		isSeparator(currUtf8Code) == isSeparator(prevUtf8Code)
}

func isNumeric(char rune) bool {
	return char >= 48 && char <= 57
}

func isAlpha(char rune) bool {
	return (char >= 65 && char <= 90) || (char >= 97 && char <= 122)
}

func isSeparator(char rune) bool {
	return char == 32 || // " "
		char == 45 || // "-"
		char == 46 || // "."
		char == 58 || // ":"
		char == 44 || // ","
		char == 59 // ";"
}

func defineTokenType(tokenValue string) (TokenType, error) {
	_, err := strconv.Atoi(tokenValue)
	if err != nil {
		for _, month := range MonthValues {
			if month == tokenValue {
				return Month, nil
			}
		}
		for _, day := range DayValues {
			if day == tokenValue {
				return Day, nil
			}
		}
		for _, tz := range TimeZoneValues {
			if tz == tokenValue {
				return Timezone, nil
			}
		}
		for _, char := range tokenValue {
			if !isSeparator(char) {
				return 0, InvalidTokenType
			}
		}
		return Separator, nil
	}

	if strings.Contains(tokenValue, "3") || strings.Contains(tokenValue, "15") {
		return Hour, nil
	}
	if strings.Contains(tokenValue, "1") {
		return Month, nil
	}
	if strings.Contains(tokenValue, "2") && len(tokenValue) <= 2 {
		return Day, nil
	}
	if strings.Contains(tokenValue, "4") {
		return Min, nil
	}
	if strings.Contains(tokenValue, "5") {
		return Sec, nil
	}
	if strings.Contains(tokenValue, "6") {
		return Year, nil
	}
	if strings.Contains(tokenValue, "7") {
		return Timezone, nil
	}

	return 0, InvalidTokenType
}

func defineTokenTypeFunc(anyCharOfToken rune) (func(char rune) bool, error) {
	if isAlpha(anyCharOfToken) {
		return isAlpha, nil
	}
	if isNumeric(anyCharOfToken) {
		return isNumeric, nil
	}
	if isSeparator(anyCharOfToken) {
		return isSeparator, nil
	}
	return nil, InvalidTokenType
}

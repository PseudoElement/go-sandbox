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
	FromIdx int
	ToIdx   int
	Type    int8
}

func tokenizeLayout(layout string) ([]Token, error) {
	tokens := []Token{}

	var tokenValue string = ""
	token := Token{}
	var prevUtf8Code rune = 0
	for idx, currUtf8Code := range layout {
		if idx == 0 || sameAsPrevious(currUtf8Code, prevUtf8Code) {
			tokenValue += string(currUtf8Code)
		} else {
			token.ToIdx = idx - 1
			tokenType, err := defineTokenType(tokenValue)
			if err != nil {
				return []Token{}, err
			}
			// token.Value = tokenValue
			token.Type = tokenType
			tokens = append(tokens, token)
			// next token generation
			tokenValue = string(currUtf8Code)
			token = Token{FromIdx: idx}
		}

		// handle last token in the end of layoutString
		if utf8.RuneCountInString(layout)-1 == idx {
			token.ToIdx = idx
			tokenType, err := defineTokenType(tokenValue)
			if err != nil {
				return []Token{}, err
			}
			// token.Value = tokenValue
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

	m := make(map[TokenType]int, len(tokens))
	var timezone *time.Location
	var month time.Month
	for _, token := range tokens {
		if token.Type == Separator {
			continue
		}

		switch token.Type {
		case Year:
			year, err := strconv.Atoi(timeString[token.FromIdx : token.ToIdx+1])
			if err != nil {
				return time.Time{}, err
			}
			m[Year] = year
		case Sec:
			sec, err := strconv.Atoi(timeString[token.FromIdx : token.ToIdx+1])
			if err != nil {
				return time.Time{}, err
			}
			m[Sec] = sec
		case Min:
			min, err := strconv.Atoi(timeString[token.FromIdx : token.ToIdx+1])
			if err != nil {
				return time.Time{}, err
			}
			m[Min] = min
		case Hour:
			hour, err := strconv.Atoi(timeString[token.FromIdx : token.ToIdx+1])
			if err != nil {
				return time.Time{}, err
			}
			m[Hour] = hour
		case Day:
			utf8Code, _ := utf8.DecodeRune([]byte{timeString[token.FromIdx : token.ToIdx+1][0]})
			if isNumeric(utf8Code) {
				day, err := strconv.Atoi(timeString[token.FromIdx : token.ToIdx+1])
				if err != nil {
					return time.Time{}, err
				}
				m[Day] = day
			} else {
				m[Day] = dayToNumeric(timeString[token.FromIdx : token.ToIdx+1])
			}
		case Month:
			utf8Code, _ := utf8.DecodeRune([]byte{timeString[token.FromIdx : token.ToIdx+1][0]})
			if isNumeric(utf8Code) {
				monthVal, err := strconv.Atoi(timeString[token.FromIdx : token.ToIdx+1])
				if err != nil {
					return time.Time{}, err
				}
				month = time.Month(monthVal)
			} else {
				month = time.Month(monthToNumeric(timeString[token.FromIdx : token.ToIdx+1]))
			}
		case Timezone:
			tz, err := time.LoadLocation(timeString[token.FromIdx : token.ToIdx+1])
			if err != nil {
				return time.Time{}, err
			}
			timezone = tz
		}
	}
	log.Println("MAP", m)
	log.Println("Month", month)

	t := time.Date(m[Year], month, m[Day], m[Hour], m[Min], m[Sec], 1000, timezone)

	return t, nil
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

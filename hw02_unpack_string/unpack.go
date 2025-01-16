package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidString = errors.New("invalid string")
	ByteDelimiter    = []byte(`\`)[0]
)

func Unpack(s string) (string, error) {
	// Place your code here.
	var newStr strings.Builder
	var elCurrent, elNext int
	var elBack byte
	var continueIterValue, delimiter bool

	for i, r := range s {
		// условия пропусков
		if r >= 128 {
			continue
		}
		delimiter = checkDelimiterEl(r, delimiter)

		if checkDelimiterSet(i, s, r) {
			delimiter = true
			newStr.WriteString(string(r))
		}

		elBack = resetLastValue(i, s, elBack)
		continueIterValue, elBack = continueIter(elBack, r, delimiter)

		if continueIterValue {
			continue
		}
		// текущий элемент
		elCurrent = i
		// следующий
		if i+1 < len(s) {
			elNext = i + 1
			// если число составное = ошибка
			if verifyComNum(r, rune(s[elNext]), elBack) {
				return "", ErrInvalidString
			}
		} else if i+1 > len(s) {
			elNext = i
		}

		// если первый лемент числовой = ошибка
		if i == 0 && unicode.IsDigit(r) {
			return "", ErrInvalidString
		}

		// если ошибок не произошло - делаем строку новую
		if unicode.IsDigit(rune(s[elNext])) {
			in, _ := strconv.Atoi(string(s[elNext]))
			if elCurrent != elNext {
				newStr.WriteString(strings.Repeat(string(s[elCurrent]), in))
			}
		} else {
			if !unicode.IsDigit(r) {
				newStr.WriteString(string(r))
			} else if elBack == ByteDelimiter {
				newStr.WriteString(string(r))
			}
		}
	}
	return newStr.String(), nil
}

func verifyComNum(r, rNext rune, elBack byte) bool {
	if unicode.IsDigit(r) && unicode.IsDigit(rNext) && elBack != ByteDelimiter {
		return true
	}
	return false
}

func checkDelimiterEl(r rune, delimiter bool) bool {
	if byte(r) != ByteDelimiter {
		delimiter = false
	}
	return delimiter
}

func checkDelimiterSet(i int, s string, r rune) bool {
	if i > 0 && len(s) >= i+2 && byte(r) == ByteDelimiter && s[i+1] == ByteDelimiter && s[i+2] == ByteDelimiter {
		return true
	}
	return false
}

func resetLastValue(i int, s string, elBack byte) byte {
	if i > 0 && elBack != s[i-1] {
		elBack = 0
	}
	return elBack
}

func continueIter(elBack byte, r rune, delimiter bool) (bool, byte) {
	var res bool
	if byte(r) == ByteDelimiter && elBack != ByteDelimiter {
		elBack = ByteDelimiter
		res = true
	}
	if delimiter {
		res = true
	}
	return res, elBack
}

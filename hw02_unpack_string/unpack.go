package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	// Place your code here.
	var newStr string
	var elCurrent, elNext int
	var elBack string
	var continueIterValue, delimiter bool

	for i, r := range s {
		// условия пропусков
		if r >= 128 {
			continue
		}

		delimiter = checkDelimiterEl(r, delimiter)

		delimiter, newStr = checkDelimiterSet(i, s, newStr, delimiter, r)

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
			checkSteamNum := string(r) + string(s[elNext])
			fmt.Println("составное число: ", checkSteamNum)
			_, err := strconv.Atoi(checkSteamNum)
			if err == nil {
				if elBack != `\` {
					return "", ErrInvalidString
				}
			}
		} else if i+1 > len(s) {
			elNext = i
		}
		// предыдущий
		if i == 0 {
			// если первый лемент числовой = ошибка
			fmt.Println("первый элемент строки: ", string(r))
			_, err := strconv.Atoi(string(r))
			if err == nil {
				return "", ErrInvalidString
			}
		}

		// если ошибок не произошло - делаем строку новую
		in, err := strconv.Atoi(string(s[elNext]))
		if err == nil {
			if elCurrent != elNext {
				fmt.Println(in)
				newStr += strings.Repeat(string(s[elCurrent]), in)
			}
		} else {
			_, err := strconv.Atoi(string(r))
			if err != nil {
				newStr += string(r)
			}
			if elBack == `\` {
				newStr += string(r)
			}
		}

		fmt.Println("cur - ", elCurrent)
		fmt.Println("next -", elNext)
		fmt.Println("===============")
	}
	fmt.Printf("Новая строка:%s %T\n", newStr, newStr)
	return newStr, nil
}

func checkDelimiterEl(r rune, delimiter bool) bool {
	if string(r) != `\` {
		delimiter = false
	}
	return delimiter
}

func checkDelimiterSet(i int, s, newStr string, delimiter bool, r rune) (bool, string) {
	if i > 0 && len(s) >= i+2 && string(r) == `\` && string(s[i+1]) == `\` && string(s[i+2]) == `\` {
		delimiter = true
		newStr += string(r)
	}
	return delimiter, newStr
}

func resetLastValue(i int, s, elBack string) string {
	if i > 0 && elBack != string(s[i-1]) {
		elBack = ""
	}
	return elBack
}

func continueIter(elBack string, r rune, delimiter bool) (bool, string) {
	var res bool
	if string(r) == `\` && elBack != `\` {
		elBack = `\`
		res = true
	}
	if delimiter {
		res = true
	}
	return res, elBack
}

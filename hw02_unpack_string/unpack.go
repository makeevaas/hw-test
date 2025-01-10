package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")
var Sepatator = `\`

func Unpack(s string) (string, error) {
	// Place your code here.
	var newStr string
	var elCurrent, elNext int
	var elBack string
	for i, r := range s {
		// условия пропусков
		if r >= 128 {
			continue
		}

		if i > 0 && string(r) == Sepatator && len(s) >= i+1 && string(s[i+1]) == Sepatator {
			continue
		}

		if i > 0 && string(r) == Sepatator && len(s) >= i+1 && string(s[i+1]) != Sepatator && string(s[i-1]) != Sepatator {
			elBack = Sepatator
			continue
		}

		if i > 0 && elBack != string(s[i-1]) {
			elBack = ""
		}

		fmt.Printf("%s\n", elBack)
		// текущий элемент
		elCurrent = i
		// следующий
		if i+1 < len(s) {
			elNext = i + 1
			// если число составное = ошибка
			if elBack != Sepatator {
				checkSteamNum := string(r) + string(s[elNext])
				fmt.Println("составное число: ", checkSteamNum)
				_, err := strconv.Atoi(checkSteamNum)
				if err == nil {
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
			} else {
				if string(s[elNext]) == `\` {
					newStr += string(r)
				}
			}

		}

		fmt.Println("cur - ", elCurrent)
		fmt.Println("next -", elNext)
		fmt.Println("===============")
	}
	fmt.Printf("Новая строка:%s %T\n", newStr, newStr)
	return newStr, nil
}

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
	var separator bool
	for i, r := range s {
		// условия пропусков
		if r >= 128 {
			continue
		}
		if i > 0 && len(s) >= i+2 && string(r) == `\` && string(s[i+1]) == `\` && string(s[i+2]) == `\` {
			separator = true
			newStr += string(r)
		}
		if string(r) != `\` {
			separator = false
		}
		if i > 0 && elBack != string(s[i-1]) {
			elBack = ""
		}
		if string(r) == `\` && elBack != `\` {
			elBack = `\`
			continue
		}
		if separator {
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

		fmt.Printf("%v\n", separator)

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

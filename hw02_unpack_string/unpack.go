package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrInvalidString    = errors.New("invalid string")
	FirstElementOfArray = 0
)

func Unpack(input string) (string, error) {
	isPreviousDigit := false

	splitedStringArr := splitString(input)
	var sb strings.Builder

	for i := 0; i < len(splitedStringArr); i++ {
		char := splitedStringArr[i]   // Получаем текущий символ
		if i == FirstElementOfArray { // Первый элемент массива
			_, err := strconv.Atoi(char) // Преобразуем в целое число
			if err == nil {
				return "", ErrInvalidString // Обработка ошибки преобразования
			}
			if len(splitedStringArr) > 1 {
				continue
			}
			sb.WriteString(char)
		}
		// Обработка следующих элементов
		char = splitedStringArr[i]
		num, err2 := strconv.Atoi(char)
		if err2 != nil { // Если это не цифра и предыдущий цифра - continue.
			// Если это не цифра и предыдущий не цифра - пишем предыдущую букву
			if isPreviousDigit {
				if i != len(splitedStringArr)-1 {
					isPreviousDigit = false
					continue
				}
				sb.WriteString(splitedStringArr[i])
			}
			sb.WriteString(splitedStringArr[i-1])
			if i == len(splitedStringArr)-1 {
				sb.WriteString(splitedStringArr[i])
			}
		}
		// Если это цифра пишем букву num раз
		if isPreviousDigit {
			return "", ErrInvalidString
		}
		for j := 0; j < num; j++ {
			sb.WriteString(splitedStringArr[i-1])
		}
		isPreviousDigit = true
	}

	return sb.String(), nil
}

func splitString(input string) []string {
	result := make([]string, len(input))
	for i, r := range input {
		result[i] = string(r)
	}
	return result
}

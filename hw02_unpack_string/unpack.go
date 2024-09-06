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

	if len(input) == 0 {
		return "", nil
	}

	splitedStringArr := splitString(input)
	var sb strings.Builder

	_, err := strconv.Atoi(splitedStringArr[FirstElementOfArray]) // Преобразуем в целое число
	if err == nil {                                               // Обработка ошибки преобразования
		return "", ErrInvalidString
	}
	if len(splitedStringArr) == 1 { // Если массив состояит только из 1 буквы - пишем
		sb.WriteString(splitedStringArr[FirstElementOfArray])
	}

	for i := 1; i < len(splitedStringArr); i++ {
		char := splitedStringArr[i]
		num, err2 := strconv.Atoi(char)
		if err2 != nil { // Если это не цифра и предыдущий цифра - continue.
			// Если это не цифра и предыдущий не цифра - пишем предыдущую букву
			if isPreviousDigit {
				if i != len(splitedStringArr)-1 {
					isPreviousDigit = false
					continue
				}
				sb.WriteString(splitedStringArr[i])
				continue
			}
			sb.WriteString(splitedStringArr[i-1])
			if i == len(splitedStringArr)-1 {
				sb.WriteString(splitedStringArr[i])
			}
			continue
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
	// Создаём срез, который будет содержать руны
	var result []string

	// Проходим по строке как по рунным литералам
	for _, r := range input {
		// Добавляем каждую руну в результат как строку
		result = append(result, string(r))
	}
	return result
}

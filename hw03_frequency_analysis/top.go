package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(input string) []string {
	if len(input) == 0 {
		return []string{}
	}

	// Создаем мапу (слово, количество повторений)
	wordCount := make(map[string]int)
	words := strings.Fields(input)

	//Заполнили мапу
	for _, word := range words {
		if word == "" {
			continue
		}
		wordCount[word]++
	}

	// Создаем слайс уникальных значений из мапы и сортируем по убыванию
	setCount := make([]int, 0, len(wordCount))
	countSet := make(map[int]bool) // мапа
	for _, count := range wordCount {
		if !countSet[count] {
			countSet[count] = true
			setCount = append(setCount, count)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(setCount)))

	// Проходим по setCount и выбираем блоки согласно count и сразу их сортируем без учета регистра
	var result []string
	for _, count := range setCount {
		var listWords []string
		for word, cnt := range wordCount {
			if cnt == count {
				listWords = append(listWords, word)
			}
		}

		//Лексиграфическая сортировка слов, которые повторяется n раз
		sort.Slice(listWords, func(i, j int) bool {
			return listWords[i] < listWords[j]
		})

		result = append(result, listWords...)
		if len(result) >= 10 {
			return result[:10]
		}
	}

	return result
}

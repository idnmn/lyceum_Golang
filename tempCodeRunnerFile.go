package main

import (
	"fmt"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {
	// Преобразуем строки в нижний регистр и удаляем пробелы
	str1 = strings.ToLower(strings.ReplaceAll(str1, " ", ""))
	str2 = strings.ToLower(strings.ReplaceAll(str2, " ", ""))

	// Сортируем строки и сравниваем их
	return strings.Compare(str1, str2) == 0
}

func main() {
	fmt.Println(AreAnagrams("listen", "silent")) // true
	fmt.Println(AreAnagrams("hello", "world"))   // false
}

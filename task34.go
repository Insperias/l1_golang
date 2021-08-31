package main

import "fmt"

func uniqStrings(str string) bool {
	uniq := make(map[rune]bool)

	for _, ch := range str { //проходим по символам
		if _, ok := uniq[ch]; !ok { //если символа не было, то добавляем его
			uniq[ch] = true
		} else {
			return false //иначе строка не состоит из уникальных символов
		}
	}
	return true
}

func main() {
	testStr1 := "abcd"
	testStr2 := "abbc"
	fmt.Println(uniqStrings(testStr1))
	fmt.Println(uniqStrings(testStr2))
}

package main

import (
	"fmt"
	"strings"
)

func ReverseStringWords(s ...string) string {
	for i, j := 0, len(s)-1; i < j; { //проходим, пока не дойдем до середины
		s[i], s[j] = s[j], s[i] //меняем слова с конца и сначала
		i++
		j--
	}
	return strings.Join(s, " ") //объединяем все в строку
}

func main() {
	str := "snow dog sun"
	fmt.Println(ReverseStringWords(strings.Fields(str)...))
}

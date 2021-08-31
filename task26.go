package main

import "fmt"

//функция переворачивает строку даже с нечитаемыми символами
func ReverseRune(s string) string {
	res := make([]byte, len(s))
	prevPos, resPos := 0, len(s)
	for pos := range s { //проходим по символам
		resPos -= pos - prevPos            //вычисляем, с какого байта нужно записать символ в результат
		copy(res[resPos:], s[prevPos:pos]) //копируем символ в массив байт из строки
		prevPos = pos
	}
	copy(res[0:], s[prevPos:]) //копируем последний символ
	return string(res)
}

//работает только с символами, которые присутствуют в юникоде
func Reverse(in string) (out string) {
	for _, r := range in {
		out = string(r) + out
	}
	return
}

func main() {
	for _, s := range []string{
		"Ångström",
		"Hello, 世界",
		"\xff\xfe\xfd", // invalid UTF-8
	} {
		fmt.Printf("%q\n", ReverseRune(s))
	}
}

package main

import "fmt"

func main() {
	var x int64 = 129
	fmt.Println("Введите номер бита:")
	var i uint8
	fmt.Scanf("%v", &i)
	fmt.Println("Установить 1 или 0: ")
	var choice byte
	fmt.Scanf("%v", &choice)
	if choice == 1 {
		x |= (1 << i) //установка i-го бита в 1 (поразрядное или)
	} else {
		x &^= (1 << i) //установка i-го бита в 0 (сброс бита И НЕ)
	}
	fmt.Println(x)
}

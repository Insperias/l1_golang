package main

import (
	"fmt"
	"reflect"
)

func printType(t interface{}) {
	switch t.(type) { //для определения типа используем переключатель типа (специальная форма)
	case int:
		fmt.Println("its int")
	case string:
		fmt.Println("its string")
	case bool:
		fmt.Println("its bool")
	case chan int:
		fmt.Println("its int channel")
	case chan string:
		fmt.Println("its string channel")
	case chan bool:
		fmt.Println("its bool channel")
	default:
		fmt.Println("incompatible type")
	}

}

func printTypeReflection(t interface{}) {
	xType := reflect.TypeOf(t) //используем рефлексию для определения типа
	fmt.Println(xType)
}

func main() {
	out := make(chan int)
	printType(out)
	printTypeReflection(out)

}

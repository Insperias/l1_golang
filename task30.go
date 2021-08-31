package main

import (
	"fmt"
	"reflect"
)

// func delWithOrder(slice []int, i int) []int {
// 	return append(slice[:i], slice[i+1:]...)
// }

func delWithoutOrder(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1] //перемещаем последний элемент на i-ю позицию
	return slice[:len(slice)-1]    //обрезаем срез
}

func delWithOrder(a interface{}, i int) []interface{} {
	slice := make([]interface{}, 0)
	av := reflect.ValueOf(a) //получаем значение

	for ind := 0; ind < i; ind++ { //создаем слайс из элементов, стоящих перед i-ым
		slice = append(slice, av.Index(ind).Interface())
	}

	for ind := i + 1; ind < av.Len(); ind++ { //добавляем к слайсу элементы, стоящие после i-го
		slice = append(slice, av.Index(ind).Interface())
	}

	return slice

}

func DelEleInSlice(arr interface{}, index int) {
	vField := reflect.ValueOf(arr)                                      //получаем значение (а точнее указатель на значение)
	value := vField.Elem()                                              //получаем значение
	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array { //если это слайс или массив, значит можно удалить элемент
		result := reflect.AppendSlice(value.Slice(0, index), value.Slice(index+1, value.Len())) //создаем новый слайс
		value.Set(result)
	}
}

func main() {
	// arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []string{"simple", "foo", "bar", "hello"}
	// arr2 := []int{1, 2, 3, 4, 5}

	fmt.Println(delWithOrder(arr2, 1))
	DelEleInSlice(&arr2, 1)
	fmt.Println(arr2)

}

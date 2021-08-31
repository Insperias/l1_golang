package main

import (
	"fmt"
	"reflect"
)

func Hash(a interface{}, b interface{}) []interface{} {
	set := make([]interface{}, 0)      //результат пересечения
	hash := make(map[interface{}]bool) //map для элементов, входящих в первый массив
	av := reflect.ValueOf(a)           //первый массив
	bv := reflect.ValueOf(b)           //второй массив

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		hash[el] = true //записываем, что конкретный элемент есть в массиве
	}

	for i := 0; i < bv.Len(); i++ {
		el := bv.Index(i).Interface()
		if _, found := hash[el]; found { //если элемент есть во втором массиве
			set = append(set, el) //записываем его в результат
		}
	}

	return set
}

// func Intersection(a, b []int) []int {
// 	m := make(map[int]uint8)
// 	for _, k := range a { //для k-го элемента, где k - элемент а
// 		m[k] |= (1 << 0) //устанавливаем 0 бит в 1
// 	}
// 	for _, k := range b { //для k-го элемента, где k - элемент b
// 		m[k] |= (1 << 1) //устанавливаем 1 бит в 1
// 	}

// 	var intersect []int
// 	for k, v := range m {
// 		a := v&(1<<0) != 0 //0 бит установлен в 1 или нет
// 		b := v&(1<<1) != 0 //1 бит установлен в 1 или нет
// 		if a && b { //если 0 и 1 бит == 1
// 			intersect = append(intersect, k) //то элемент есть в обоих массивах
// 		}
// 	}
// 	return intersect
// }

func main() {
	arr1 := []string{"foo", "bar", "foo"}
	arr2 := []string{"foo", "hello", "world", "sample", "string"}
	fmt.Println(Hash(arr1, arr2))

	// intArr1 := []int{1, 5, 3, 2, 7, 4, 5}
	// intArr2 := []int{5, 7, 1, 4, 8, 11, 13, 47}
	// fmt.Println(Intersection(intArr1, intArr2))

}

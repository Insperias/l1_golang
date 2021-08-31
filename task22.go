package main

import (
	"fmt"
	"math/rand"
)

func qSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// выбрать опорный элемент
	pivotIndex := rand.Int() % len(a)

	// переместить опорный элемент вправо
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// переместить элементы меньше опорного влево
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// установить опорный элемент после последнего меньшего элемента
	a[left], a[right] = a[right], a[left]

	// вызвать данный метод для части справа опорного элемента и слева опорного
	qSort(a[:left])
	qSort(a[left+1:])

	return a
}

func main() {
	nItems := 20
	a := make([]int, nItems)
	for i := range a {
		a[i] = rand.Intn(10)
	}
	qSort(a)
	fmt.Println(a)
}

// func qsort_pass(arr []int, done chan int) []int{
//     if len(arr) < 2 {
//         done <- len(arr) //в канал записать длину массива
//         return arr //вернуть массив если состоит из 1 элемента
//     }
//     pivot := arr[0] //выбрать опорный элемент
//     i, j := 1, len(arr)-1 // установить начало и конец массива
//     for i != j { //пока не пройдем массив
//         for arr[i] < pivot && i!=j{ //элементы меньше опорного слева
//             i++
//         }
//         for arr[j] >= pivot && i!=j{ //элементы больше опорного справа
//             j--
//         }
//         if arr[i] > arr[j] { //если элемент слева больше элемента справа
//             arr[i], arr[j] = arr[j], arr[i] //меняем
//         }
//     }
//     if arr[j] >= pivot { //
//         j--
//     }
//     arr[0], arr[j] = arr[j], arr[0] //поместить на позицию j опорный элемент
//     done <- 1; //отправить значение 1 в канал (сортировка выполнена)

//     go qsort_pass(arr[:j], done) //запустить сортировку для левой части
//     go qsort_pass(arr[j+1:], done) //запустить сортировку для правой части
//     return arr
// }

// func qsort(arr []int) []int {
//     done := make(chan int)
//     defer func() {
//         close(done)
//     }()

//     go qsort_pass(arr[:], done) //запустить сортировку

//     rslt := len(arr)
//     for rslt > 0 { //как только значение достигнет 0, то все горутины завершили работу
//         rslt -= <-done; //и массив отсортирован
//     }
//     return arr
// }

// func main() {
//     rand.Seed(time.Now().UTC().UnixNano())
//     arr_rand := make([]int, 20)
//     for i := range arr_rand {
//         arr_rand[i] = rand.Intn(10)
//     }
//     fmt.Println(arr_rand)
//     qsort(arr_rand)
//     fmt.Println(arr_rand)
// }

package main

import "fmt"

func binarySearch(a []int, search int) (result int, searchCount int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1 // элемент не найден
	case a[mid] > search:
		result, searchCount = binarySearch(a[:mid], search)
	case a[mid] < search:
		result, searchCount = binarySearch(a[mid+1:], search)
		if result >= 0 { // если что-либо кроме результата -1 "не найдено"
			result += mid + 1
		}
	default: // a[mid] == search
		result = mid // элемент найден
	}
	searchCount++
	return
}

func main() {
	arr := []int{1, 2, 3, 7, 10, 14, 15, 18, 20}
	res, count := binarySearch(arr, 3)
	fmt.Println(arr[res], count)

}

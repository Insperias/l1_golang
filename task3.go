package main

import "fmt"

func main() {
	var sum int
	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan int) //создаем канал для передачи значений квадратов чисел

	go func(out chan<- int, arr ...int) {
		for _, el := range arr {
			out <- el * el
		}
		close(out) //закрываем канал после того, как передали туда все значения
	}(ch, arr...)

	for sqr := range ch { //считываем данные из канала и суммируем
		sum += sqr
	}

	fmt.Println(sum)
}

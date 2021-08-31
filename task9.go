package main

import (
	"fmt"
	"os"
)

//возвращаем канал, в который записываются числа из массива
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n //записываем число в канал
		}
		close(out) //закрываем канал после окончания записи
	}()
	return out
}

//возвращаем канал, в который записывается результат 2*х, где
//где х - считанное число из первого канала
func mul2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for x := range in {
			out <- 2 * x
		}
		close(out)
	}()
	return out
}

//выводим данные в stdout из канала
func write(in <-chan int) {
	for x := range in {
		fmt.Fprintln(os.Stdout, x)
	}
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	g := gen(arr...) //передаем данные не как список, а как последовательность чисел
	write(mul2(g))

}

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func randNumsGenerator(n int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //создаем объект Rand

	out := make(chan int) //канал для отправки значений
	go func() {
		for i := 0; i < n; i++ {
			out <- r.Intn(n) //отправляем рандомные числа (между 0 и n) в канал
		}
		close(out)
	}()
	return out
}

func eversNums(nums <-chan int) <-chan int {
	out := make(chan int) //канал для записи четных чисел

	go func() {
		for num := range nums { //считываем значения
			if num%2 == 0 {
				out <- num //отправляем во второй канал только четные
			}
		}
		close(out)
	}()
	return out //возвращаем канал
}

func main() {
	n := 1000
	rnds := randNumsGenerator(n)
	for x := range eversNums(rnds) {
		fmt.Fprintln(os.Stdout, x)
	}

}

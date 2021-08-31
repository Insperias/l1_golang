package main

import (
	"fmt"
	"time"
)

func main() {
	nSec := 1
	ch := make(chan int)
	timer := time.NewTimer(time.Duration(nSec) * time.Second) //создаем таймер на n секунд
	go func(ch chan<- int) {
		number := 0
		for {
			ch <- number //записываем значения в канал
			number++
		}
	}(ch)
LOOP: //метка для выхода из цикла
	for {
		select {
		case <-timer.C: //если таймер закончился, то здесь будет находиться текущее время
			fmt.Println("timer.C timeout happend")
			close(ch)  // закрываем канал
			break LOOP //выходим по метке из цикла
		case value := <-ch:
			fmt.Println(value)
		}
	}
}

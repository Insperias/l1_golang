package main

import (
	"fmt"
	"sync"
)

func main() {
	var counters = map[int]int{}
	mu := &sync.Mutex{} //создаем объект Mutex

	for i := 0; i < 5; i++ {
		go func(counters map[int]int, th int, mu *sync.Mutex) {
			for j := 0; j < 5; j++ {
				mu.Lock() //блокируем mutex для закрытия доступа к данным в данной части
				counters[th*10+j]++
				mu.Unlock() //снимаем блокировку и открываем доступ к данным
			}
		}(counters, i, mu)
	}
	fmt.Scanln()
	mu.Lock() //...
	fmt.Println("counters result", counters)
	mu.Unlock() //...
}

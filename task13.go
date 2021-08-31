package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}

/*
	В результате получим deadlock, т.к. переменная WaitGroup будет бесконечно ждать окончание горутин, когда они уже закончили работу.
	Это произошло потому, что переменная wg передается в анонимную функцию без адреса (передается копия переменной). Поэтому в WaitGroup
	счетчик останется на 5, потому что он уменьшается только в копии переменной, а не в основной.
*/

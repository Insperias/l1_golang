package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	out := os.Stdout
	arr := []int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{} //создаем объект WaitGroup для ожидания окончания работы всех горутин
	for _, el := range arr {
		wg.Add(1)
		go func(number int, out io.Writer, wg *sync.WaitGroup) {
			defer wg.Done() //убираем один элемент из WaitGroup после выхода из горутины
			fmt.Fprintln(out, number*number)
		}(el, out, wg)
	}

	wg.Wait() //ожидаем завершение выполнения всех горутин
}

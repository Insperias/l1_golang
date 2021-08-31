package main

import (
	"fmt"
	"os"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ch := make(chan int)

	go func(in chan<- int, arr ...int) {
		for _, el := range arr {
			ch <- el
		}
		close(in)
	}(ch, arr...)

	for el := range ch {
		fmt.Fprintln(os.Stdout, el)
	}
}

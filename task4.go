package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

//принимает канал с данными и канал, отвечающий за сигнал завершения
func worker(elems <-chan interface{}, die chan bool) {
	for {
		select {
		case el := <-elems:
			fmt.Fprintln(os.Stdout, reflect.ValueOf(el)) //получаемые значения могут быть любыми
		case <-die:
			fmt.Fprintln(os.Stdout, "ENDED")
			die <- true
			return
		}
	}
}

func main() {
	const nWorkers = 3
	elements := make(chan interface{})
	die := make(chan bool)

	for w := 0; w < nWorkers; w++ {
		go worker(elements, die) //создаем n воркеров
	}

	sc := bufio.NewScanner(os.Stdin) //создаем объект сканнера для считывания данных из потока ввода
	for sc.Scan() {
		element := sc.Text() //считываем данные и возвращаем их в строковом виде
		if element == "c" {  //символ конца ввода

			die <- true // отправляем сигнал завершения
			break       // и выходим
		}
		elements <- element
	}
	<-die //ждем, когда все горутины завершат работу

}

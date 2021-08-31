package main

import "fmt"

func Generator() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				close(ch)
				return
			}
		}
	}()
	return ch
}

func main() {
	number := Generator()
	fmt.Println(<-number)
	fmt.Println(<-number)
	number <- 0 // передаем сигнал завершения для остановки горутины
	fmt.Println(<-number)
	// …
}

// func main() {
// 	cancelCh := make(chan struct{}) // канал отмены для завершения горутины
// 	dataCh := make(chan int)
// 	go func(cancelCh chan struct{}, dataCh chan int) {
// 		val := 0
// 		for {
// 			select {
// 			case <-cancelCh:
// 				return
// 			case dataCh <- val: // если принимаем значение из канала отмены, значит завершаем горутину
// 				val++
// 			}
// 		}
// 	}(cancelCh, dataCh)
// 	for curVal := range dataCh {
// 		fmt.Println("read", curVal)
// 		if curVal > 3 {
// 			fmt.Println("send cancel")
// 			cancelCh <- struct{}{} // отправляем в канал отмены в качестве значения пустую структуру (которая не занимает место)
// 			break
// 		}
// 	}
// }

// import (
//     "context"
//     "fmt"
//     "time"
// )

// func main() {
//можно завершить горутину при помощи контекста
//     forever := make(chan struct{})
//     ctx, cancel := context.WithCancel(context.Background())

//функция, принимающая контекст, запускает горутину и ожидает ее возврата или отмены контекста.
//Оператор select помогает определить, что случится первым, и завершить работу функции.

//     go func(ctx context.Context) {
//         for {
//             select {
//             case <-ctx.Done():  // контекст
//                 forever <- struct{}{}
//                 return
//             default:
//                 fmt.Println("for loop")
//             }

//             time.Sleep(500 * time.Millisecond)
//         }
//     }(ctx)

//     go func() {
//         time.Sleep(3 * time.Second)
//         cancel() //вызываем функцию отмены для завершения контекста
//     }()

//     <-forever
//     fmt.Println("finish")
// }

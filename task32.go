package main

import (
	"fmt"
	"time"
)

func sleepv1(sec float64) {
	<-time.After(time.Second * time.Duration(sec))
}

func main() {
	sleepv1(0.5)
	fmt.Println("slept")
}

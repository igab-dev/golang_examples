package main

import (
	"fmt"
	"time"
)

func firstRotine(times int, ch chan int) {

	for i := 0; i < times; i++ {
		result := 666
		ch <- result
	}
}

func secondRotine(times int, ch2 chan int) {

	for i := 0; i < times; i++ {
		result := 777
		ch2 <- result
	}
}

func main() {

	ch := make(chan int)
	go firstRotine(10, ch)
	time.Sleep(time.Second)
	returnFisrtRotine := <-ch
	fmt.Println(returnFisrtRotine)

	ch2 := make(chan int)
	go secondRotine(10, ch2)
	time.Sleep(time.Second)
	reuturnSecondRotine := <-ch2
	fmt.Println(reuturnSecondRotine)

}

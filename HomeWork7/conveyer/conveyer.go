package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0; x <= 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// возведение в квадрат
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				close(squares)
				return
			}
			squares <- x * x
		}
	}()

	// печать
	for sq := range squares {
		fmt.Println(sq)
	}
}

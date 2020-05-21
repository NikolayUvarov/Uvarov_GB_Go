package main

import (
	"fmt"
)

var rub = 0

const rate float32 = 67

func main() {
	fmt.Print("USD from RUB calc\nEnter roubles:")
	n, err := fmt.Scanln(&rub)
	if err == nil && n == 1 {
		fmt.Printf("Result: %d roubles is %.2f USD", rub, float32(rub)/rate)
	} else {
		fmt.Println("Error in parameters: %v %v", n, err)
	}
}

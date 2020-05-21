package main

import (
	"calculator"
	"fmt"
)

func main() {
	input := ""
	for {
		fmt.Print("Go Calc> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "exit" {
			break
		}
		if input == "exit" {
			calculator.Help()
		} else {
			if res, err := calculator.Calculate(input); err == nil {
				fmt.Printf("Результат: %v\n", res)
			} else {
				fmt.Println("Не удалось произвести вычисление, наберите help для справки.")
			}
		}
	}
}

package main

import (
	"fmt"
	"math"
)

var rub float64 = 0
var percent float64 = 0

const years int = 5

func main() {
	fmt.Print("Cложный процент за ", years, " лет.\nВведите сумму вклада и проценты за год через пробел:")
	n, err := fmt.Scanln(&rub, &percent)
	if err != nil || n != 2 {
		fmt.Errorf("Error")
	}

	fmt.Printf("Cумма за %d лет: %.2f", years, rub*math.Pow((1+percent/100), float64(years)))

}

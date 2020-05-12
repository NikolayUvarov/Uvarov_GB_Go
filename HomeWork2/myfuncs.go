package main

import (
	"fmt"
	"math"
	"math/big"
)

const a float64 = 3
const b float64 = 4

var c float64 = 0

func isEven(i int64) bool {
	return i%2 == 0
}
func isDivisibleBy3(i int64) bool {
	return i%3 == 0
}
func fibb100() {
	f0 := big.NewInt(0)
	f1 := big.NewInt(1)
	fmt.Printf(" %25d", f0)
	for i := 2; i <= 100; i++ {
		fmt.Printf(" %25d", f1)
		if i%10 == 0 {
			fmt.Println()
		}
		f0, f1 = f1, new(big.Int).Add(f0, f1)
	}
}
func eratosfen(n int64) {
	fmt.Println("Eratosfen to ", n)
	s := make([]int64, n)
	max := int64(math.Sqrt(float64(n))) + 1
	var i, j int64
	for i = 2; i <= max; i++ {
		if s[i] == 0 {
			for j = i * i; j < n; j += i {
				s[j] = 1
			}
		}
	}

	for i = 2; i < n; i++ {
		if s[i] == 0 {
			fmt.Printf("%d ", i)
		}
	}

}

func main() {
	fmt.Println("isEven 2", isEven(2))
	fmt.Println("isEven 3", isEven(20001))
	fmt.Println("isDivBy3 6", isDivisibleBy3(6))
	fmt.Println("isDivBy3 7", isDivisibleBy3(7))
	fmt.Println("isDivBy3 112", isDivisibleBy3(112))
	fmt.Println("isDivBy3 111", isDivisibleBy3(111))
	fibb100()
	eratosfen(100)
}

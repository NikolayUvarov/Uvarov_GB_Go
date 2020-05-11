package main

import (
	"fmt"
	"math"
)

const a float64 = 3
const b float64 = 4

var c float64 = 0

func main() {
	fmt.Println("Triangle: a =", a, "b =", b)
	c = math.Sqrt(a*a + b*b)
	fmt.Println("Периметр:", a+b+c)
	fmt.Println("Площадь:", a*b/2)
	fmt.Println("Длина гипотенузы:", c)
}

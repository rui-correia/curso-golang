package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtraćão =>", a-b)
	fmt.Println("Divisao =>", a/b)
	fmt.Println("Multiplicacao =>", a*b)
	fmt.Println("Modulo =>", a%b)

	//bitwise
	fmt.Println("E =>", a&b)   // 11 & 10 = 10
	fmt.Println("Ou =>", a|b)  // 11 & 10 = 11
	fmt.Println("XoR =>", a^b) // 11 & 10 = 01

	c := 3.0
	d := 2.0

	//outras operacões usando math...
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("O menos valor =>", math.Min(c, d))
	fmt.Println("Exponenciacao =>", math.Pow(c, d))
}

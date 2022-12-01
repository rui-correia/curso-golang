package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {

	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "function"
	default:
		return "Não sei."

	}
}

func main() {

	fmt.Println(tipo(9.2))
	fmt.Println(tipo(9))
	fmt.Println(tipo("OPA"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))

}

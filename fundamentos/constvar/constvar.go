package main

import (
	"fmt"
	//alias para importacão
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo float64 (inferido pelo compilador)

	//forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferencia é", area)

	//criando constantes em blocos
	const (
		a = 1
		b = 2
	)
	//criando variaveis em blocos
	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)
	//criando varias variaveis na mesma linha
	var e, f bool = true, false
	fmt.Println(e, f)
	//criando varias variaveis na mesma linha de forma simplificada
	g, h, i := 2, false, "epa"
	fmt.Println(g, h, i)
}

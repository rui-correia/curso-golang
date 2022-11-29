package main

import "fmt"

func main() {
	i := 1
	//Go não tem aritmética de ponteiros
	var ponteiro *int = nil

	ponteiro = &i //pegando endereço da variável i e atribuindo para ponteiro
	*ponteiro++   //Desreferenciando o ponteiro (pegando valor)
	i++
	fmt.Println(ponteiro, *ponteiro, i)
}

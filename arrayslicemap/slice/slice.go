package main

import (
	"fmt"
	"reflect"
)

func main() {
	array1 := [3]int{1, 2, 3} //array
	slice1 := []int{1, 2, 3}  //slice

	fmt.Println(array1, slice1)
	fmt.Println(reflect.TypeOf(array1), reflect.TypeOf(slice1))

	array2 := [5]int{1, 2, 3, 4, 5}

	// Slice não é um Array. Slice define um pedaço de um array
	slice2 := array2[1:3]
	fmt.Println(array2, slice2)

	slice3 := array2[:2] // Novo slice, mas aponta para o mesmo array
	fmt.Println(array2, slice3)

	slice4 := slice2[:1]
	fmt.Println(slice2, slice4)
}

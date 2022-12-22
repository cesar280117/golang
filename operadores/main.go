package main

import (
	"fmt"
)

func main() {
	// var edades []int
	// edades = append(edades, 10, 20, 40, 60)

	// jugadores := make([]string, 3)
	// fmt.Println(jugadores)

	// personas := make(map[int]string)
	// personas[0] = "cesar"
	// personas[1] = "Lizbeth"
	// personas[3] = "otro"

	// myName, exists := personas[0]
	// fmt.Println(myName, exists)
	// fmt.Println(personas)

	// usuarios:=map[int]string{
	// 	1:"Pedro",
	// 	2:"pablo",
	// }
	// for _, v := range personas {
	// 	// fmt.Println(v)
	// }

	// for i := 0; i < 100; i++ {
	// 	fmt.Println(i)
	// }

	numeros := []int{
		1, 2, 3, 80, 25, 60, 54, 55, 76, 33,
	}

	fmt.Println("antes de incrementar ", numeros)

	for i, v := range numeros {
		numeros[i] = v + 20
	}

	fmt.Println("despues de incrementar ", numeros)

}

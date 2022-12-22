package main

import (
	"fmt";
	"strconv"
)

func main() {
	//los valors int son positivos o negativos
	var myIntVar int
	myIntVar = 25
	fmt.Println("my variable es de valor ", myIntVar)

	//los valores solo son positivos
	var myUiInt uint
	myUiInt = 90
	fmt.Print("el valor de mi variable es de ", myUiInt)

	nombre := "cesar"
	fmt.Printf("\nhola mi nombre es: %s \n", nombre)


	var becado bool
	becado = true
	fmt.Println("estoy becado?", becado)
	fmt.Println("la direccion de memoria de mi variable es", &nombre)

	const pi =3.1416
	fmt.Println(pi)


	edad:=20
	fmt.Printf("type: %T, valor: %d\n",edad,edad)
	edadString:=fmt.Sprintf("%d",edad)
	fmt.Printf("type: %T, valor: %s",edadString,edadString)

	age,_ := strconv.ParseInt("23",0,64)
	fmt.Printf("type:%T, value:%d y el error \n",age,age)
	
	peso,_:=strconv.ParseFloat("20.90",64)
	fmt.Printf("type:%T, value:%f",peso,peso)
}

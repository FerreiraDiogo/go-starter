package main

import (
	"fmt"
	"math"
)

//variaveis globais
var cc, python, java bool

//variaveis globais com incializadores.
var ruby, golang , elixir = 1,true, "WTF?!"

func main(){
	fmt.Println("sapo")	
	fmt.Println(math.Pi)
	fmt.Println(add(2,2))
	a,b := swapStrings("o sapo", "pulou")
	fmt.Print(a,b)
	c,d,e,f := evenMoreReturnValues()
	println(c,d,e,f)
	println(nakedReturnFunction(17))

	var i int
	println(i,cc,java, python)
	println(ruby, golang, elixir)

	println(manyWaysToDeclareVariables())

	println(typeCastingInAction())

}


func add( a, b int) int{
	//isso é um comentário de uma linha
	/*
	muitas
	linhas
	de coomentario
	*/

	return a + b
}

func swapStrings (a,b string) (string, string){
	return b,a
}

func evenMoreReturnValues()(int,int,int,int){
	return 1,2,3,4
}

func nakedReturnFunction( sum int)(x, y int){
	x = sum * 4/9
	y = sum - x
	return
}

func manyWaysToDeclareVariables() (string){
	var firstVariableType string
	var firstText = "You can assign a variable without declaring its type by initializing it."
	secondType := "You can use the quick assignment operator( := ) to declare a variable without the var keyword. This is just a shortcu for declaring and initializing a variable."
	lastText := "Quick assignment is only available within functions or block scopes"
	
	return firstText + secondType + lastText + firstVariableType
}

func typeCastingInAction()(int, float64){

	i := 2
	f := 6.22

	return int(f), float64(i)

}
package main

import (
	"fmt"
)

func main(){

	//o laço for...range é só um enhanced for. 
	//Ele possui um indice e um valor. 
	
	a :=[]int{1,2,3,4,5}
	fmt.Println("for value in range...")
	for i, v := range a {

		fmt.Printf("%d = i, %d = v\n", i,v)
	}

	//Pode-se rejeitar qualquer um dos dois valores o atribuindo a uma variavel nomeada _
	b :=[]int{1,2,3}
	for _, v := range b {
		fmt.Println(v)
	}

	//pode-se também excluir totalmente o valor dado ao range
	c :=[]int{1,2,3}
	for i := range c {
		fmt.Println(c[i])
	}
}
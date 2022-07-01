package main

import (
	"fmt"
)

type Vertex struct{
	X int
	Y int
	// a,b int
}

func main(){
	//Structs are a collection of fields. They remind me of the JS objects
	fmt.Println(Vertex{1,2})

	v:= Vertex{3,4}
	//We can access those fields through with dot operator...
	fmt.Println(v.X)

	v.X = 5
	fmt.Println(v.X)

	//or with the bloody pointers!

	i := &Vertex{5,6}
	fmt.Println(i.X)

	fmt.Println(i)
	fmt.Println(*i)
	fmt.Println((*i).X)



}
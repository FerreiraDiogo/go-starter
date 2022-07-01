package main

import (
	"fmt"
	"math"
	"runtime"
)

func main(){

	fmt.Println("Now to go's conditional clauses.")
	isTrue := true
	if isTrue {
		fmt.Println("Within the if conditional")
	}

	if v:= math.Pow(2,2); v < 10{
		fmt.Println(v)
	}

	if(isTrue){
		println("We can use parenthesis for if clauses, but they aren't mandatory")
	}

	fmt.Println("Now, the switch statements. They work just as in c-like syntaxes, except that the brak keyword is provided out of the box")
	os := runtime.GOOS
	switch (os) {
		case "linux": fmt.Println("Linux")
		case "darwin": fmt.Println("OS X")
		default : fmt.Printf("Within the first switch. My Operational System is %s\n",os)
	}

	fmt.Println("\nParenthesis are also optional and we can also declare a quick variable within the switch declaration")
	switch myOs := runtime.GOOS; myOs{ //I dont like this syntax. Although its shorter i think its less readable
		case "linux": fmt.Println("Linux")
		case "darwin": fmt.Println("OS X")
		default : fmt.Printf("Within the first switch. My Operational System is %s\n",os)
	}




}
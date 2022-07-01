/*
Whatahell one has in mind to try and learn a programming language
that does not manage its own pointers?!
*/

package main

import "fmt"

func main(){

	i:= 42

	p := &i

	z :=0

	fmt.Printf("Variavel i: %v\n", i)
	fmt.Printf("Variavel p: %v\n", p)
	fmt.Printf("*p: %v. It equals to i\n", *p)
	fmt.Printf("&i: %v. It equals to p.\n", &i)

	z= *p/2
	*p = z

	fmt.Printf("z: %v, i: %v, p: %v", z, i, p )

}
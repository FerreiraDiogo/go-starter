package main

import (
	"fmt"
)


func main(){

	fmt.Println(sqrt(4))

}


func sqrt(x float64) float64{
	z := 1.0
	control := z
	for i :=1; i < 10; i++ {
		z -= (z*z - x) / (2*z)
		if(z == control){
			return z
		}
		control = z
		fmt.Println(z)
		

	}
	return z
}
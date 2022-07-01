package main

import "fmt"

func main(){

	fmt.Println("Go has only one looping construct: the for loop.")
	fmt.Println("Go1s for loop control variable must be initialized with the quick assignment operator")

	for i:=0 ; i <11; i++ {

		fmt.Printf("Basic for loop, printing the value of the control variable: %v \n",i)
	}

	fmt.Println("\n\n Now the continued for. Works just as a conventional while...do loop:")
	sum := 1
	for ;sum < 100;{  //while(sum < 100)
		sum += sum
		fmt.Printf("Within the continued for loop. The sum variable value is %v \n", sum)
	}

	fmt.Println("\n\n For the sake o clarity, both semicolons can be ommited: ")
	sum2 :=1
	for sum2 < 1000 {
		sum2 += sum2
		fmt.Printf("Sum2 value is %v\n", sum2)
	}



	
}
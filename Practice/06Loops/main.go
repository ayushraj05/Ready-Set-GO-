package main

import "fmt"

// we only have for loop in go for looping

func main(){
	// we will implement the for loop while style

	i:= 0;
	for i<=5 {
		fmt.Print(i," ");
		i++;
	}

	// for infinite loop
	// for{

	// }

	// clasic for loop

	for i:= 1; i< 7; i++{
		fmt.Print(i," ")
	}
	fmt.Print(i)
}

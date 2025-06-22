package main

import "fmt"

// if the variable name start with capital letter then the variable is public


func main() {
	var username string = "Ayush"
	fmt.Println(username)
	fmt.Printf("variable is of the type %T \n",username)

	userage := 69 // we can declare variables like this as well :)
	// but only inside a function not golbaly :)

	// there is type inference as well

	fmt.Print("Hello!! we can easily use the print statment \n")

	fmt.Println(userage)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input :)\n"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the Rating : ")

	// comma ok syntak || err ok syntax

	input , _ := reader.ReadString('\n')

	fmt.Println("Thanks for the rating : ",input)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Pizza store!!")
	fmt.Println("Rate our pizza out of 5 :-")

	reader := bufio.NewReader(os.Stdin)

	input , _ := reader.ReadString('\n')

	fmt.Println("Thanks for the rating ", input)

	numRating,err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println("The Rating you gave is : ",numRating+1)
	}
}

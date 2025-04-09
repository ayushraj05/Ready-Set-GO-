package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is to deal with time in go lang :)")
	presentTime := time.Now()

	fmt.Println(presentTime.Format("02-01-2006 Mon 15:04:05"))
}

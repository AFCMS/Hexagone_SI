package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing number argument")
		return
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("invalid number")
		return
	}

	charLen := len(strconv.Itoa(number))

	for i := 0; i <= number; i++ {
		fmt.Printf("%0*d\n", charLen, i)
	}
}

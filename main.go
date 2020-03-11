package main

import (
	"fmt"
	"github.com/nikoheikkila/roman-encoder/encoder"
	"os"
	"strconv"
)

func usage() string {
	return "Usage: ./roman <integer>"
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage())
		return
	}

	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("Error: Invalid argument %s given", os.Args[1])
		return
	}

	roman, err := encoder.Encode(n)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(roman)
}

package main

import (
	"fmt"
	"strconv"
)

func BinToDec() {
	var binary string
	fmt.Print("Enter Binary Number:")
	fmt.Scanln(&binary)
	output, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Output %d", output)
}

func DecToBin() {
	var int int64
	fmt.Print("Enter Int Number:")
	fmt.Scanln(&int)
	output := strconv.FormatInt(int, 2)

	fmt.Printf("Output %s", output)
}

func main() {
	DecToBin()
	BinToDec()
}

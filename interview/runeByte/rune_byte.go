package main

import "fmt"

func main() {
	s := "lorem"

	fmt.Printf("%T\n", s) // string
	fmt.Printf("%T\n", s[0]) // byte alias uint8
	for _, v := range s {
		fmt.Printf("%T\n", v) // int32 alias rune
	}
}

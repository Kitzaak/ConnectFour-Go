package main

import (
	"connect_four/board"
	"fmt"
)

func Hello() string {
	return "Hello, world!"
}

func main() {
	fmt.Println(Hello())
	fmt.Println(board.DisplayEmpty())
}

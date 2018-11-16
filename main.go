package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	tokens := tokenize("def a() 1 end")
	for _, token := range tokens {
		fmt.Println(token)
	}
}

package main

import (
	"fmt"
	"gocy/commandsl"
)

func main() {
	// initation()
	// commandsl.CalculateG()
	commandsl.ForWhile()
}

func initation() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}
}

package main

import (
	"fmt"
	"gocy/class"
)

func main() {
	//initation()
	//commandsl.CalculateG()
	//commandsl.ForWhile()
	//commandsl.LogicalT()
	//commandsl.YearCheck(2002)
	//commandsl.SwitchTest2()
	//commandsl.ForL(10)
	//commandsl.Infinity(10)
	//commandsl.Variables()
	//commandsl.Randomly(50)
	class.Piggy()
}

func initation() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}
}

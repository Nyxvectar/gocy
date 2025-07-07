package main

import (
	"fmt"
	"gocy/class"
	"gocy/commandsl"
)

func main() {
	initation(false)
	switchCmd(false)
	switchClass(false)
}

func switchClass(whether bool) {
	if whether == true {
		class.Piggy()
		class.Consts()
		class.TurnToTure("1")
		class.TurnToFalse("1")
	}
}

func switchCmd(whether bool) {
	if whether == true {
		commandsl.CalculateG()
		commandsl.ForWhile()
		commandsl.LogicalT()
		commandsl.YearCheck(2002)
		commandsl.SwitchTest2()
		commandsl.ForL(10)
		commandsl.Infinity(10)
		commandsl.Variables()
		commandsl.Randomly(50)
	}
}

func initation(whether bool) {
	if whether == true {
		s := "gopher"
		fmt.Printf("Hello and welcome, %s!\n", s)
		for i := 1; i <= 5; i++ {
			fmt.Println("i =", 100/i)
		}
	}
}

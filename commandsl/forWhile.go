package commandsl

import (
	"fmt"
	"os"
	"strings"
)

var walkOutside = true
var takeTheBlueHill = false
var command = "Walk outside"

func ForWhile() {
	boolTest(true)
	elseIfTest(1)
}

func elseIfTest(order int) {
	//The variables name is in the front and the class
	//of it should be put behind its name.
	//var judgeResult = strings.Contains(order, "")
	if order >= 0 {
		//If func can be put into a if, and it is same
		//for the other case like while, case, etc.
		if order <= 100 {
			if order == 1 {
				println("The num is 1")
			} else if order == 2 {
				println("The num is 2")
			} else {
				println("The num is %v", order)
			}
			//Just for fun, do not ask why not directly get %v
		}
	} else {
		print("exit")
		os.Exit(0)
	}
}

func boolTest(whether bool) {
	fmt.Println("You found that you was in a cavern")
	var exit = strings.Contains(command, "outside")
	//The check is sensitive with upper or lower case.
	//Contains can check the "" rather than "outside"
	//It seems that it means a form of boolean true.
	fmt.Println("You left the cave:", exit)
	if whether == true {
		if walkOutside == true {
			if takeTheBlueHill == true {
				print("Got it.")
			} else {
				print("Failed")
			}
		}
	}
}

// a possible list of usually syns are
// > < >= <= == != etc.

func SwitchTest1() {
	println("There is a cavern entrance here and a path to the east")
	var commanda = "go inside"
	// This usage(1) would directly compare the value they have
	// An expected output is "You find yourself in a dimly lit cavern"
	switch commanda {
	case "go east":
		print("You head further up the hountain")
	case "enter the cave", "go inside":
		println("You find yourself in a dimly lit cavern")
	case "read sign":
		println("The sign reads 'No Minors'.")
	default:
		println("Did not quite get that.")
	}
}

func SwitchTest2() {
	// The switch test 2's form is longer, and if there is no
	// necessary, we prefer to use the form 1.
	var room = "lake"
	switch {
	case room == "cave":
		println("You find yourself in a dimly lit cavern")
	case room == "lake":
		println("The ice seems solid enough.")
		fallthrough
	// The fallthrough can fall through the next branch and run.
	// But we need to know that it would only run the next branch
	// rather than the whole chains of the all case sentences below.
	case room == "underwater":
		println("The ice seems solid enough.(Twice)")
		fallthrough
	case room == "lake":
		println("The ice seems solid enough.(Third)")
	}
}

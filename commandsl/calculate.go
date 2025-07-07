package commandsl

import (
	"fmt"
)

// global variables
var a = 2
var b = 3
var dateG = "Aug 28, 2020"

func CalculateG() {
	basicR()
	variablesAndConst()
}

func variablesAndConst() {
	const lightSpeed, distance = 299792, 800000000 //kms, ms
	fmt.Printf("\nTime is %v", distance/1000/lightSpeed)
	// the sequence of / does no matter due to it falling.,
}

func basicR() {
	fmt.Println("Hello and this is a test")
	fmt.Println(50 * 30) // expected 1500
	fmt.Println(50 % 30) // expected 20
	fmt.Println(50 / 30) // expected 1
	var c = a + b/a
	var d = a * 2
	fmt.Println(c, d)
	fmt.Printf("I joined GitHub on %v CE, and now is %v", dateG, 2025)
	// the number behind letters would automatically add a space
}

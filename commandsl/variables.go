package commandsl

var x = 10

func Variables() {
	{
		var y = 20
		println(y)
	}
	var y = 30
	println(x, y)
	// The expexted y output is 30 but not 20
	// because the {} declares a domain of each
	// variable.
}

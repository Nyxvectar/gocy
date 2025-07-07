package commandsl

var aa = true
var bb = false

func LogicalT() {
	if aa || bb {
		// Expected is true because there had a true
		// within aa and bb
		print("aa || bb is ", aa || bb)
	} else if aa && bb {
		print("aa && bb is ", aa && bb)
	} else {
		print("both || and && could not run")
	}
}

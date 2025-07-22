/**
 * Author:  Nyxvectar Yan
 * Repo:    gocy
 * Created: 07/22/2025
 */

package block

func PointerTrain() {
	var NilP *int
	var temp1, temp2 = 1, 2
	var result1, result2 *int
	// '&' at here means 'Get the address'
	// '*' at here means 'Get the Object'
	result1, result2 = &temp1, &temp2
	var a, b = *result1, *result2
	{
		println("The plus answer is", a+b)
		println("Nil output is like:", NilP)
		// The nil one expected a output like 0x0
	}

}

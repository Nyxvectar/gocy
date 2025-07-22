/**
 * Author:  Nyxvectar Yan
 * Repo:    gocy
 * Created: 07/22/2025
 */

package block

import "fmt"

func Struc() {
	var i, n = 0, 0
	{
		// This is a array.
		// The length of it can not be edited.
		// an expected output is '2'
		var arrayName = [2]int{1, 2}
		temp := arrayName[1]
		fmt.Println(temp)
		{
			var Ett [10]int
			for i < 10 {
				Ett[n] = i
				fmt.Printf("Element[%v] = %v\n", n, Ett[i])
				n++
				i++
			}
		}
	}
}

package basictypes

import "fmt"

func Array() {
	var array [2]string
	array[0] = "oi"
	array[1] = "tchay"

	fmt.Println("Receba array in golang", array)
}

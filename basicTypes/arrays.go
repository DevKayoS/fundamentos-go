package basictypes

import "fmt"

func Array() {
	var array [2]string
	array[0] = "oi"
	array[1] = "tchay"

	fmt.Println("Receba array in golang", array)
}

func Slice() {
	var receba []string
	receba = append(receba, "gracas", "a deus", "pai")
	receba = append(receba, "e nada krai", "teste", "toma")

	fmt.Println("Tamanho do slice: ", len(receba))
	lengthSlice := len(receba)

	fmt.Println("Dividindo deveria vir ate a metade", receba[lengthSlice/2:])

	receba = receba[:lengthSlice/2]

	fmt.Println("Receba como fazer um array sem ter que padronizar sei la porra", receba)
}

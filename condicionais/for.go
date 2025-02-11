package condicionais

import "fmt"

func For() {
	sum := 0
	// for normal
	for i := 0; i < 10; i++ {
		// acontece a logica aqui 10 vezes
		sum++
	}

	// fazendo for com slice
	var slice = []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(slice); i++ {
		fmt.Println("Esse foi o valor do indice", slice[i])
	}

}

func While() {
	sum := 0
	for sum < 10 {
		sum++
		fmt.Println("aprendendo o basico de novo", sum)
	}
}

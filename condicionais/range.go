package condicionais

import "fmt"

func Range() {
	//slice
	nums := []int{1, 2, 3, 4, 5}
	// key e o indice
	// value literalmente e o valor
	for key, value := range nums {
		fmt.Println(key, value)
	}

	// slice com string
	names := []string{"kayo", "julia", "Pedrao", "Carlos"}
	for chave, valor := range names {
		fmt.Println(chave, valor)
	}

	// map de string
	users := map[string]string{
		"nome":     "kayo",
		"idade":    "19",
		"endereco": "rua bandeira das esmeraldas",
	}

	for _, value := range users {
		fmt.Println(value)
	}

}

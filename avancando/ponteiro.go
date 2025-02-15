package avancando

import "fmt"

type Pessoa struct {
	Name string
}

func Ponteiro() {
	p1 := Pessoa{
		Name: "Kayo",
	}

	p2 := Pessoa{
		Name: "Jubs",
	}

	//impressionante essa parada realmente funcionar
	var p3 *Pessoa = &p1
	p3.Name = "Isso daqui"

	fmt.Println("Pessoa 1", p1)
	fmt.Println("Pessoa 2", p2)
	fmt.Println("Pessoa 3", p3)

}

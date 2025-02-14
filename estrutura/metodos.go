package estrutura

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) Apresentar() {
	fmt.Printf("Ola muito prazer! Me chamo %s e tenho %d anos \n", p.Nome, p.Idade)
}

func Metodos() {
	pessoa1 := Pessoa{
		Nome:  "Kayo",
		Idade: 19,
	}

	pessoa1.Apresentar()
}

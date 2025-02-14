package estrutura

import "fmt"

type Cliente struct {
	Nome     string
	Idade    int
	Endereco []string
	Email    string
	Trabalho Trabalho
}

type Trabalho struct {
	Titulo      string
	Senioridade string
	Empresa     string
}

func StructClient() {
	cliente1 := Cliente{
		Nome:  "Kayo",
		Idade: 19,
		Email: "kayo@example.com",
	}

	cliente1.Endereco = append(cliente1.Endereco, "Rua bandeira")
	cliente1.Endereco = append(cliente1.Endereco, "Rua Esmeraldas")
	cliente1.Endereco = append(cliente1.Endereco, "Rua francisco")

	fmt.Println("Como sera que e retornado essa struct", cliente1)
	fmt.Println("Retornando o nome do Cliente", cliente1.Nome)
	fmt.Println("Retornando a idade do Cliente", cliente1.Idade)
	fmt.Println("Retornando a idade do Endereco", cliente1.Endereco)
	fmt.Println("Retornando a idade do Email", cliente1.Email)

	for k, v := range cliente1.Endereco {
		fmt.Println(k+1, v)
	}

	cliente2 := Cliente{
		Nome:  "Ze da manga",
		Idade: 18,
		Email: "mango@joe.com",
		Trabalho: Trabalho{
			Titulo:      "Analista de infraestrutra",
			Senioridade: "Junior",
			Empresa:     "Amazon",
		},
	}

	cliente2.Endereco = append(cliente2.Endereco, "Rua das Flores")

	fmt.Println("Rapaz do ceu que trem grande", cliente2)
}

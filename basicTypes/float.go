package basictypes

import "fmt"

func FloatTeste() {
	fmt.Println("Isso aqui e um float pls aprenda: ")
	areaCirc := calcRaioCirculo(2.2)

	areaCircFormated := fmt.Sprintf("%.2f", areaCirc)
	fmt.Println("calculo da area do circulo", areaCircFormated)
}

func calcRaioCirculo(raio float64) float64 {
	pi := 3.14

	area := (raio * raio) * pi
	return area
}

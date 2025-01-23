package condicionais

import (
	"fmt"
	"time"
)

func Switch() {
	fmt.Println("Quando que é sexta?")
	today := time.Now().Weekday()

	switch time.Friday {
	case today + 0:
		fmt.Println("SOmente apenas hoje o dia da maldade SEXTOU")
	case today + 1:
		fmt.Println("amanha vai ser o puro")
	case today + 2:
		fmt.Println("Daqui dois dias")
	default:
		fmt.Println("Seloco num compensa ta longe dms")
	}
}

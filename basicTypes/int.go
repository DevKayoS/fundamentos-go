package basictypes

import (
	"fmt"
	"fundamentosGolang/utils"
)

func Int() {
	idade := 19

	texto := fmt.Sprintf("Idade: %d", idade)

	utils.GerarMensagemTerminal(texto)
}

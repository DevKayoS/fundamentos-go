package main

import (
	"fundamentosGolang/utils"
	"log/slog"
)

func main() {
	slog.Info("[INICIANDO PROGRAMA]")
	texto := "nossa lobo mau "
	texto += "e essas orelhas grandonas"
	utils.GerarMensagemTerminal(texto)
}

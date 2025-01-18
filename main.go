package main

import (
	basictypes "fundamentosGolang/basicTypes"
	"log/slog"
)

func main() {
	slog.Info("[INICIANDO PROGRAMA]")
	basictypes.Int()
	basictypes.FloatTeste()

	slog.Info("[FINALIZANDO PROGRAMA]")
}

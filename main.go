package main

import (
	basictypes "fundamentosGolang/basicTypes"
	"log/slog"
)

func main() {
	slog.Info("[INICIANDO PROGRAMA]")
	// basictypes.Int()
	// basictypes.FloatTeste()
	// basictypes.Array()
	// basictypes.Slice()
	basictypes.Map("Julia")
	slog.Info("[FINALIZANDO PROGRAMA]")
}

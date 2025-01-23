package main

import (
	"fundamentosGolang/condicionais"
	"log/slog"
)

func main() {
	slog.Info("[INICIANDO PROGRAMA]")
	// basictypes.Int()
	// basictypes.FloatTeste()
	// basictypes.Array()
	// basictypes.Slice()
	// basictypes.Map("Julia")
	// condicionais.Serase()
	condicionais.Switch()
	slog.Info("[FINALIZANDO PROGRAMA]")
}

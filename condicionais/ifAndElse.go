package condicionais

import (
	"errors"
	"fmt"
	"log/slog"
)

func Serase() {
	nota := 32

	if nota >= 10 {
		fmt.Println("Toma bobo passou")
	} else {
		fmt.Println("Toma bobo reprovou")
	}

	if err := returnError(); err != nil {
		fmt.Println(err.Error())
	}

	var players = map[string]int{
		"Kayo": 10,
	}

	if value, ok := players["Kayo"]; ok {
		fmt.Println("Tem um jogador:", value, ok)
	}
}

func returnError() error {
	slog.Error("Isso gerou um log de erro")
	return errors.New("Isso aqui deveria ser um erro")
}

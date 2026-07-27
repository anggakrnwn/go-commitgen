package main

import (
	"fmt"

	"github.com/anggakrnwn/go-commitgen/internal/infrastructure/config"
)

func main() {
	cfg := config.Connection()

	fmt.Println("[TEST CONFIG]: berhasil dimuat!")
	fmt.Println("Default model:", cfg.DefaultModel)
	if cfg.GeminiAPIKey != "" {
		fmt.Println("model api key terdeteksi!")
	} else {
		fmt.Println("model kosong!")
	}

}

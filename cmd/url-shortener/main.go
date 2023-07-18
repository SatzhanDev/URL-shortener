package main

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {
	// TODO: init config: cleanenv
	cfg := config.MustLoad()
	fmt.Println(cfg)
	// TODO: init logger: slog
	// TODO: storage: sqlite
	// TODO: init router: chi, "chi render"
	// TODO: run server:
}

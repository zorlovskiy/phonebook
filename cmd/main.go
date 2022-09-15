package main

import (
	"fmt"

	"github.com/caarlos0/env/v6"

	"github.com/zorlovskiy/phonebook/config"
	"github.com/zorlovskiy/phonebook/internal/app"
)

func main() {

	// чтение конфига псмотри либу caarlos/env
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)
	app.Run(&cfg)

}

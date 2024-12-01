package main

import (
	"liatdong/internal/config"
)

func main() {
	viperConfig := config.NewViperConfig()
	app := config.NewFiberApp(viperConfig)
	log := config.NewLogger(viperConfig)
	pool, err := config.NewPgxConnection(viperConfig)
	if err != nil {
		// gimana ya?
	}

	config.Bootstrap(&config.ConfigInstance{
		Pool: pool,
		App:  app,
		Log:  log,
		// Validate: validate,
		Config: viperConfig,
	})

	log.Println("Starting server on :8080...")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf(err.Error())
	}
}

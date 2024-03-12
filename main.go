package main

import (
	"github.com/Schierke/personal-site/config"
	"github.com/Schierke/personal-site/handler"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func loadAppConfig() (*config.AppConfig, error) {
	// Load app config from file
	configFile := ".env"

	cfg, err := config.LoadAppConfig(configFile)

	if err != nil {
		return nil, err
	}

	return &cfg, err
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	server := handler.NewServer(handler.WithPort(5000))

	err := server.Start()

	if err != nil {
		log.Fatal().Err(err).Msg("can't start the server")
	}

}

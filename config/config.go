package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type LocalConfig struct {
	SpotifyKey    string `env:"SPOTIFY_KEY"`
	SpotifySecret string `env:"SPOTIFY_SECRET"`
	SpotifyURL    string `env:"SPOTIFY_URL"`
}

func GetConfig() (*LocalConfig, error) {
	var localConfig LocalConfig
	//will use godotenv to load the file
	err := godotenv.Load("./local.env")

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	//from the loaded file we will parse the environment variable into localconfig instance
	if err := env.Parse(&localConfig); err != nil {
		log.Fatalf("Error reading the environment variable %v", err)
		return nil, err
	}
	log.Printf("%+v\n", localConfig)
	return &localConfig, nil
}

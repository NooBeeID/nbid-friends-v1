package config

import "github.com/joho/godotenv"

func LoadConfig(filename string) {
	godotenv.Load(filename)
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() map[string]string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error al cargar el archivo .env", err)
	}

	var variables = map[string]string{"CLIENT_ID": os.Getenv("CLIENT_ID"), "CLIENT_SECRET": os.Getenv("CLIENT_SECRET")}
	return variables
}

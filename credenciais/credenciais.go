package credenciais

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Credenciais(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env")
	}

	return os.Getenv(key)
}

package main

import (
	"os"

	"github.com/exercise/cmd/api/handlers"
	"github.com/exercise/internal/pkg/pokemon"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// main Init
func main() {
	loadEnv()
	port := os.Getenv("port")

	r := gin.Default()

	clientPokemon := pokemon.New()
	handler := handlers.New(clientPokemon)
	handler.API(r)

	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}

// loadEnv load the environment variables
func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

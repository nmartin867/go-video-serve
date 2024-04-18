package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"github.com/nmartin/go-video-serve/config"
)

func loadEnv() {
	env := os.Getenv("GIN_MODE")
	if "" == env {
		env = "development"
	}
	if err := godotenv.Load(".env." + env + ".local"); err != nil {

	}
	if err := godotenv.Load(); err != nil {

	}
}

func main() {
	config.Init()
}

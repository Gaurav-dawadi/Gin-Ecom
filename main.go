package main

import (
	"go-practice/application"
	"go-practice/utils/logger"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		logger.LogFatal("Couldnot load environment file")
	}
}

func main() {
	logger.LogInfo("Before loading env")
	loadEnv()
	logger.LogInfo("After loading env")
	application.ApplicationInitialize.RunApplication(application.ApplicationInitialize{})
}

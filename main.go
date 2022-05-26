package main

import (
	"fmt"
	"go-practice/application"
	"log"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Couldnot load environment file")
	}
}

func main() {
	fmt.Println("Before loading env")
	loadEnv()
	fmt.Println("After loading env")
	application.RunApplication()
}

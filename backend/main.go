package main

import (
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Warn(fmt.Sprintf("Did not load .env file: %v", err))
	}
	formatter := &log.TextFormatter{FullTimestamp: true}
	log.SetFormatter(formatter)
	log.SetLevel(log.DebugLevel)
}

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("POSTGRES_NONROOT_USER"),
		os.Getenv("POSTGRES_NONROOT_PASSWORD"),
		"app_database")

	a.Run(":8080")
}

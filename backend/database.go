package main

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func connectDB(a *App, pgURI string) {
	var err error
	log.Info("Attempting to connect to database")
	for i := 1; i < 6; i++ {
		a.DB, err = sql.Open("postgres", pgURI)
		if err != nil {
			log.Error("Error with opening connection with database")
			panic(err)
		}

		err = a.DB.Ping()
		if err == nil {
			break
		}
		retryWait := i * i
		log.Info(fmt.Sprintf("Connection attempt %d unsuccessful, retrying in %d seconds...", i, retryWait))
		time.Sleep(time.Duration(retryWait) * time.Second)
	}
	if a.DB == nil {
		log.Fatal("Could not connect to database")
	}
	log.Info("Successfully connected to database")
}

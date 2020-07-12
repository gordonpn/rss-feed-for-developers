package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/whiteshtef/clockwork"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.WithFields(log.Fields{"Error": err}).Warn("Problem with loading .env file")
	}
	log.SetLevel(log.DebugLevel)
}

func main() {
	_, present := os.LookupEnv("DEV")
	if present {
		job()
	} else {
		job()
		scheduler := clockwork.NewScheduler()
		oneMinute := 60000
		scheduler.SetPollingInterval(15 * oneMinute)
		scheduler.Schedule().Every(12).Hours().Do(job)
		scheduler.Run()
	}
}

func job() {
	signalHealthCheck(START)
	db := Database{}
	fetchedDevToPosts := fetchDevToPosts()
	subs := getSubreddits()
	fetchedRedditPosts := fetchRedditListings(subs)
	err := db.connect()
	checkAndPanic("Error with connecting to database", err)
	err = db.insert(fetchedDevToPosts)
	checkAndPanic("Error with inserted Dev.to posts", err)
	err = db.insert(fetchedRedditPosts)
	checkAndPanic("Error with inserting Reddit posts", err)
	signalHealthCheck(SUCCESS)
}

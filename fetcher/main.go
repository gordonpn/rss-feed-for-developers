package main

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
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
}

package main

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	_ = fetchDevToPosts()
	fetchRedditListings()
}

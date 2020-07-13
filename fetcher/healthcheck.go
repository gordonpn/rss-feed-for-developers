package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

type Status string

const (
	SUCCESS Status = ""
	START   Status = "/start"
	FAIL    Status = "/fail"
)

func signalHealthCheck(action Status) {
	_, present := os.LookupEnv("DEV")
	if present {
		return
	}
	_, present = os.LookupEnv("FETCHER_HC_UUID")
	if !present {
		log.Fatal("FETCHER_HC_UUID is not present and is needed, because running in production")
	}
	healthCheckURL := "https://hc-ping.com"
	url := fmt.Sprintf("%s/%s%s", healthCheckURL, os.Getenv("FETCHER_HC_UUID"), action)
	res, err := httpClient.Get(url)
	if err != nil {
		log.Error("Error with signaling health check")
		log.Error(err)
	}
	if res != nil {
		defer res.Body.Close()
	}
}

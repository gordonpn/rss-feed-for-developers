package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func checkAndPanic(message string, err error) {
	if err != nil {
		signalHealthCheck(FAIL)
		log.Errorf("%s", message)
		log.Panic(err)
	}
}

func getJSON(url string, target interface{}, headers map[string]string) error {
	req, err := http.NewRequest("GET", url, nil)
	checkAndPanic("Error with creating new GET request", err)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	res, err := httpClient.Do(req)
	checkAndPanic("Error with performing GET request", err)
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}

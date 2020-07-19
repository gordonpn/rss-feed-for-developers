package main

import (
	"encoding/json"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"net/http"
	"regexp"
	"time"
	"unicode"

	log "github.com/sirupsen/logrus"
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

func cleanString(unclean string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, unclean)
	re := regexp.MustCompile("[^a-zA-Z0-9 ]+")
	clean := re.ReplaceAllLiteralString(result, "")
	return clean
}

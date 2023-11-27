package main

import (
	"crypto/rand"
	"fmt"
	"strings"
)

var googleDomains = map[string]string{
	"com": "https://google.com/search?q=",
	"za":  "https://www.google.co.za/search?q=",
}

type SearchResult struct {
	ResultRank  int
	ResultURL   string
	ResultTitle string
	ResultDesc  string
}

// userAgent is a browser that makes the request
var userAgents = []string{}

func randomUserAgent() string {
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}
func buildGoogleUrls(searchTerm, countryCode, pages, count int) {
	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if googleBase, found := googleDomains[countryCode]; found {
		for i := 0; i < pages; i++ {
			start := i * count
			scrapeURL := fmt.Sprintf(googleBase, searchTerm, count, languageCode, start)

		}
	}

}
func GoogleScrape(searchTerm, countryCode, languageCode string, pages, count) ([]SearchResult, error) {
	results := []SearchResult{}
	resultCounter := 0
	googlePages, err := buildGoogleUrls(searchTerm, countryCode,languageCode, pages, count)
}
func main() {
	res, err := GoogleScrape("Andrew Huberman", "en", "com", 1, 30)
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}

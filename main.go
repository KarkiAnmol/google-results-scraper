package main

import "crypto/rand"

var googleDomains = map[string]string{}

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

func main() {

}

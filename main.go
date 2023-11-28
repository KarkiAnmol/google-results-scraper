package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
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
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}
func buildGoogleUrls(searchTerm, countryCode,languageCode string, pages, count int)([]string,error) {
	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if googleBase, found := googleDomains[countryCode]; found {
		for i := 0; i < pages; i++ {
			start := i * count
			scrapeURL := fmt.Sprintf("%s%s&num=%d&hl=%s&start=%d&filter=0",googleBase, searchTerm, count, languageCode, start)

		}
	}else{
		err =fmt.Errorf("country (%s) is currently not supported",countryCode)
		return nil,err

	}
	return toScrape,nil

}
func googleResultParsing(response *http.Response,rank int)([]SearchResult,error) {
	doc,err:=goquery.NewDocumentFromResponse(response)
	if err!=nil{
		return nil,err
	}
	result:=[]SearchResult{}
	sel:=doc.Find("div.g")
	rank++
	for i:= range sel.Nodes{
		item := sel.Eq(i)
		linkTag := item.Find("a")
		link,_:=linkTag.Attr("href")
		titleTag:=item.Find("h3.r")
		descTag:= item.Find("span.st")
		desc:=descTag.Text()
		title:=titleTag.Text()
		link =strings.Trim(link," ")
		if link != " " && link != "#" && !strings.HasPrefix(link,"/"){
			result:=SearchResult{
				rank,
				link,
				title,
				desc,
			}
			results = append(results,result)
			rank++
		}
		return results,err
	}
}
func getScrapeClient(proxyString interface{}) *http.Client{
	switch v:= proxyString.(type){
	case string:
		proxyURL,_:=url.Parse(v)
		return &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	
	default:
		return &http.Client{}
}
}
func GoogleScrape(searchTerm, countryCode, languageCode string,proxyString interface{}, pages, count,backoff) ([]SearchResult, error) {
	results := []SearchResult{}
	resultCounter := 0
	googlePages, err := buildGoogleUrls(searchTerm, countryCode,languageCode, pages, count)
	if err!=nil{
		return nil,err
	}
	for _,page:= range googlePages{
		res,err:=scrapeClientRequest(page,proxyString)
		if err!=nil{
			return nil,err
		}
		data,err :=googleResultParsing(res,resultCounter)
		if err!=nil{
			return nil,err
		}
		resultCounter += len(data)
		for _,result:= range data{
			results = append(results,result)

		}
		time.Sleep(time.Duration(backoff) *time.Second)
	}
	return googlePages,nil
}
func scrapeClientRequest(searchURL string,proxyString interface{})(*http.Response,error){
	baseClient:= getScrapeClient(proxyString)
	req,_=http.NewRequest("GET",searchURL,nil)
	req.Header().set("User-Agent",randomUserAgent())
	res,err:= baseClient.Do(req)
	if res.StatusCode !=200{
		err:= fmt.Errorf("scraper recieved a non-200 status code suggesting a ban")
		return nil,err
	}
	if err!=nil{
		return nil,err
	}
	return res,nil

}
func main() {
	res, err := GoogleScrape("Andrew Huberman", "com", "en", nil,1, 30,10)
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}

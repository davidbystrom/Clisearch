package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Scrape(query string) {
	// Request the HTML page.
	res, err := http.Get("https://google.com/search?q=" + query)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	numLinks := 1

	// Find the review items
	doc.Find("a").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		// For each item found, get the title
		url, ok := s.Attr("href")
		if ok {
			if strings.Contains(url, "/url") {
				fmt.Printf("Result %d: %s\n", numLinks, strings.Split(strings.Replace(url, "/url?q=", "", 1), "&")[0])
				numLinks++
			}
		}
		if numLinks == 11 {
			return false
		}
		return true
	})
}

func main() {
	queryArray := os.Args[1:]
	query := strings.Join(queryArray, "+")
	Scrape(query)
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

//ExampleScrape is
func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://tw.manhuagui.com/comic/7620")
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

	fmt.Println("start")
	s := doc.Find("script")
	fmt.Println(s)

	// script := s.Nodes[1].FirstChild.Data
	fmt.Println("script===>")
	// fmt.Println(script)
	// vm := otto.New()

	prefix := "https://tw.manhuagui.com"

	doc.Find("#chapter-list-1 ul li a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		fmt.Printf(s.Text())
		fmt.Println()
		href, hasvalue := s.Attr("href")
		fmt.Println(href, hasvalue)
		fmt.Println()

		url := prefix + href
		fmt.Println(url)

		// GetImage(url)

	})
}

func main() {
	ExampleScrape()
}

//GetImage is
func GetImage(url string) {
	res, err := http.Get(url)
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

	fmt.Println(doc)

	fmt.Println("mangaFile個數")
	fmt.Println(doc.Find("#imgLoading"))

	fmt.Println(len(doc.Find("#imgLoading").Nodes))
	fmt.Println("mangaFile Src==>")
	fmt.Println(doc.Find("#mangaFile").Attr("src"))

}

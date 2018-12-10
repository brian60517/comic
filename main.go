package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/robertkrimen/otto"
)

func main() {
	ExampleScrape()

	// DownloadJpg()
}

//ExampleScrape is
func ExampleScrape() {
	// Request the HTML page.
	// res, err := http.Get("https://tw.manhuagui.com/comic/7620")
	res, err := http.Get("https://tw.manhuagui.com/comic/7620/354279.html")
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
	s := doc.Find("script").Each(func(i int, sf *goquery.Selection) {
		if strings.Contains(sf.Text(), "window") {
			fmt.Println(sf.Text())
			vm := otto.New()
			_, err = vm.Run(`
				var args = 'abc';
				SMH.imgData = function(args){
					args = JSON.stringify(args);
				}
			`)

			newScript := strings.Replace(sf.Text(), `window["\x65\x76\x61\x6c"]`, "", 1)
			fmt.Println()
			fmt.Println()
			fmt.Println(newScript)
			srcValue, scriptErr := vm.Run(newScript)
			if scriptErr != nil {
				log.Fatal(scriptErr)
			}

			fmt.Println()
			fmt.Println("srcValue===>")
			fmt.Println(srcValue)

			vm.Run("SMH.imgData.preInit()")

			// 2018/12/10 18:17:36 TypeError: 'splic' is not a function

			if value, err := vm.Get("args"); err == nil {
				fmt.Println("args結果===>")
				fmt.Println(value)
				// goData, _ := value.Export()
			}
		}

	})
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

//DownloadJpg is
func DownloadJpg() {
	url := "http://i.imgur.com/m1UIjW1.jpg"
	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create("/AAA/A001.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success!")
}

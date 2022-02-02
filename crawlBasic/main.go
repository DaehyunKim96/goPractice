package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

var baseURL string = "https://kr.indeed.com/%EC%B7%A8%EC%97%85?q=golang&limit=5"

func main() {
	pages := getPages
	fmt.Println(pages())
}
func getPages() int {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkStatusCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Find("a"))
	})
	return 0
}
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func checkStatusCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status", res.StatusCode)
	}
}

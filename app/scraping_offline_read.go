package main

import (
	"io/ioutil"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

func main() {
	fileInfos, _ := ioutil.ReadFile("github_goquery.html")
	stringReader := strings.NewReader(string(fileInfos))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	doc.Find("table > tbody > tr > td.content > span > a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		fmt.Println(url)
	})
}
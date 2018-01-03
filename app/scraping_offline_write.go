package main

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	doc, err := goquery.NewDocument("https://github.com/PuerkitoBio/goquery")
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	res, err := doc.Find("body").Html()
	if err != nil {
		fmt.Print("dom get failed")
	}
	ioutil.WriteFile("github_goquery.html", []byte(res), os.ModePerm)
}
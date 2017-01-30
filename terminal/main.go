package main

import (
	"fmt"
	"github.com/moovweb/gokogiri"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print(page)
	// parse the web page
	doc, err := gokogiri.ParseHtml(page)
	if err != nil {
		panic(err)
	}
	defer doc.Free()
	fmt.Print(doc)
}

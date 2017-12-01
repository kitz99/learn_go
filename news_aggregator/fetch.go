package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// SitemapIndex is the abdtract parsed form of xml sitemap
type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"` // A slice
}

// News structure
type News struct {
	Title    []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Location []string `xml:"url>loc"`
}

func main() {
	var s SitemapIndex
	var n News

	response, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	xml.Unmarshal(bytes, &s)

	// fmt.Println(s.Locations)
	for _, location := range s.Locations { // range returns index, object
		response, _ := http.Get(location)
		bytes, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		xml.Unmarshal(bytes, &n)
	}
}

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SitemapIndex is the abdtract parsed form of xml sitemap
type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"` // A slice

}

// News structure
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

func main() {
	var s SitemapIndex
	var n News
	news_map := make(map[string]NewsMap)

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
		for idx, _ := range n.Titles {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	for title, data := range news_map {
		fmt.Println("\n\n\n")
		fmt.Println(title)
		fmt.Println(data.Keyword)
		fmt.Println(data.Location)
	}
}

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Location represents a description for an element from Sitemap
type Location struct {
	Loc string `xml:"loc"`
}

// SitemapIndex is the abdtract parsed form of xml sitemap
type SitemapIndex struct {
	Locations []Location `xml:"sitemap"` // A slice
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	response, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)
}

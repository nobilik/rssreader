package rssreader

import "time"

type RssItem struct {
	Title       string
	Source      string
	SourceURL   string
	Link        string
	PublishDate time.Time
	Description string
}

type feed struct {
	// Version string    `xml:"version,attr"`
	Title string `xml:"channel>title"`
	Link  string `xml:"channel>link"`
	Items []item `xml:"channel>item"`
}

type item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PublishDate string `xml:"pubDate"`
	Description string `xml:"description"`
}

package rssreader

import "time"

type RssItem struct {
	Title       string    `json:"title,omitempty"`
	Source      string    `json:"source,omitempty"`
	SourceURL   string    `json:"source_url,omitempty"`
	Link        string    `json:"link,omitempty"`
	PublishDate time.Time `json:"publish_date,omitempty"`
	Description string    `json:"description,omitempty"`
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

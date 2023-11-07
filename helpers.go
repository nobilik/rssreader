package rssreader

import (
	"strings"
	"time"
)

// parseTime tries to parse time from string
// maybe is too much
func parseTime(s string) (time.Time, error) {
	dateFormats := []string{
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
	}

	// Tue, 24 Oct 2023 23:00:18 GMT
	var t time.Time
	var err error

	for _, layout := range dateFormats {
		t, err = time.Parse(layout, s)
		if err == nil {
			return t, nil
		}
	}

	return t, err
}

// remaps items to output format
func mapItems(parsed *feed) []RssItem {
	var result []RssItem
	for _, item := range parsed.Items {
		r := RssItem{
			Source:      strings.TrimSpace(parsed.Title),
			SourceURL:   parsed.Link,
			Title:       strings.TrimSpace(item.Title),
			Link:        item.Link,
			Description: strings.TrimSpace(item.Description),
		}
		pd, err := parseTime(item.PublishDate)
		if err == nil {
			r.PublishDate = pd
		}

		result = append(result, r)
	}
	return result
}

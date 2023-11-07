# rssreader

Package `github.com/nobilik/rssreader` provides a simple and concurrent RSS feed parser for fetching and parsing RSS feeds from multiple URLs. It allows you to retrieve RSS items from various sources and provides easy access to the parsed data.

## Installation

To use `rssreader`, add it to your Go project's dependencies:

```bash
go get github.com/nobilik/rssreader
```

After adding it to your project, run go mod tidy to ensure that the package is included in your go.mod file.

To parse RSS feeds from multiple URLs, use the Parse function. Provide the feed URLs as variadic arguments, and you'll receive an array of RssItem and an array of errors:

```go
items, errors := rssreader.Parse("url1", "url2", ...)
```

The RssItem struct represents an RSS item with the following fields:

- Title: The title of the RSS item.
- Source: The source of the RSS feed.
- SourceURL: The URL of the RSS feed.
- Link: The link associated with the RSS item.
- PublishDate: The publication date of the RSS item (if available).
- Description: A description of the RSS item.

You can access these fields to extract information from the parsed RSS items.

### Example

```go
package main

import (
	"fmt"
	"github.com/nobilik/rssreader"
)

func main() {
	feedURLs := []string{
		"https://example.com/rss-feed-1",
		"https://example.com/rss-feed-2",
	}

	items, errors := rssreader.Parse(feedURLs...)

	for _, item := range items {
		fmt.Println("Title:", item.Title)
		fmt.Println("Source:", item.Source)
		fmt.Println("Source URL:", item.SourceURL)
		fmt.Println("Link:", item.Link)
		fmt.Println("Publish Date:", item.PublishDate)
		fmt.Println("Description:", item.Description)
	}

	for _, err := range errors {
		fmt.Println("Error:", err)
	}
}
```

### License

    This package is licensed under the MIT License.

## This package is created for testing purposes and not maintained

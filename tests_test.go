package rssreader

import (
	"testing"
)

// var tURLS = []string{
// 	// "http://rss.cnn.com/rss/cnn_topstories.rss",
// 	// "http://feeds.bbci.co.uk/news/rss.xml",
// 	// "https://techcrunch.com/rssfeeds/",
// 	// "https://arstechnica.com/rss-feeds/",
// 	// "http://feeds.bbci.co.uk/news/world/rss.xml",
// 	"http://www.chinaminingmagazine.com/rss/current.xml",
// }

func TestParse(t *testing.T) {
	feedURLs := []string{
		"http://feeds.bbci.co.uk/news/world/rss.xml",
		"http://rss.cnn.com/rss/cnn_topstories.rss",
	}

	items, errors := Parse(feedURLs...)

	if len(errors) > 0 {
		t.Errorf("Errors encountered while parsing feeds: %v", errors)
	}

	if len(items) == 0 {
		t.Errorf("No items were parsed from the feeds")
	}
}

func TestParseURL(t *testing.T) {
	testCases := []struct {
		url       string
		expectErr bool
	}{
		{"http://rss.cnn.com/rss/cnn_topstories.rss", false},
		{"https://example.com/nonexistentfeed", true},
	}

	for _, tc := range testCases {
		items, err := parseURL(tc.url)

		if tc.expectErr && err == nil {
			t.Errorf("Expected an error for URL: %s", tc.url)
		}

		if !tc.expectErr && err != nil {
			t.Errorf("Unexpected error for URL: %s - %v", tc.url, err)
		}

		if len(items) == 0 && !tc.expectErr {
			t.Errorf("No items were parsed for URL: %s", tc.url)
		}
	}
}

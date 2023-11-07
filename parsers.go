package rssreader

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/html/charset"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

// parses all provided urls
// we use waitgroup to wait for all goroutines to finish.
// and we don't care of result that we have from single goroutine.
// so we don't use channels
func Parse(feedURLs ...string) ([]RssItem, []error) {
	var wg sync.WaitGroup
	results := make(chan []RssItem)
	errorsChan := make(chan error)

	wg.Add(len(feedURLs))

	for _, url := range feedURLs {
		go func(url string) {
			defer wg.Done()
			items, err := parseURL(url)
			if err != nil {
				// for more info it can be modified to know which feed falls
				errorsChan <- err
				return
			}
			results <- items
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
		close(errorsChan)
	}()

	allItems := []RssItem{}
	var errors []error

	go func() {
		for err := range errorsChan {
			errors = append(errors, err)
		}
	}()

	for items := range results {
		allItems = append(allItems, items...)
	}

	return allItems, errors
}

func parseURL(furl string) ([]RssItem, error) {
	parsed, err := parseFeed(furl)
	if err != nil {
		return []RssItem{}, err
	}

	return mapItems(&parsed), nil
}

func parseFeed(feedURL string) (feed, error) {
	var result feed
	data, err := fetch(feedURL)
	if err != nil {
		return feed{}, err
	}

	err = xml.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("url: %s, error: %s", feedURL, err)
	}

	if len(result.Items) == 0 {
		parseItems(data, &result)
	}

	return result, nil
}

// I'm not so clear with rss variations
// on one chinese rss I saw items data outside channel "http://www.chinaminingmagazine.com/rss/current.xml"
// so maybe it needs some improvement
func parseItems(data []byte, result *feed) {
	r := bytes.NewReader(data)
	d := xml.NewDecoder(r)
	// if is not utf-8
	d.CharsetReader = charset.NewReaderLabel
	for {
		token, err := d.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		switch se := token.(type) {
		case xml.StartElement:
			if se.Name.Local == "item" {
				var itemResult item
				if err := d.DecodeElement(&itemResult, &se); err != nil {
					return
				}
				result.Items = append(result.Items, itemResult)
			}
		}
	}
}

// reads data from url
func fetch(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/xml")
	response, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("url: %s, error: %s", url, err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("url: %s, status code error: %d %s", url, response.StatusCode, response.Status)
	}
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("url: %s, error: %s", url, err)
	}

	return contents, nil
}

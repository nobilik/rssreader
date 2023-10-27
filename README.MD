add "github.com/nobilik/rssreader" to yuor import
go mod tidy
use rssreader.Parse("url1", "url2"...)

as result you'll have an array of RssItems

```
type RssItem struct {
    Title string `json:"title,omitempty"`
    Source string `json:"source,omitempty"`
    SourceURL string `json:"source_url,omitempty"`
    Link string `json:"link,omitempty"`
    PublishDate time.Time `json:"publish_date,omitempty"`
    Description string `json:"description,omitempty"`
}
```

and an array of errors

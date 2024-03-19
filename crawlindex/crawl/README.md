# Package [cloudeng.io/glean/crawlindex/crawl](https://pkg.go.dev/cloudeng.io/glean/crawlindex/crawl?tab=doc)

```go
import cloudeng.io/glean/crawlindex/crawl
```


## Types
### Type Crawler
```go
type Crawler struct {
	// contains filtered or unexported fields
}
```
Crawler represents a crawler instance that contains global configuration
information.

### Functions

```go
func New(resources Resources) *Crawler
```
New creates a new Crawler instance.



### Methods

```go
func (c *Crawler) Run(ctx context.Context, fv *Flags, datasource string) error
```




### Type Flags
```go
type Flags struct {
	config.FileFlags
	Outlinks bool `subcmd:"outlinks,false,display extracted outlinks"`
	Progress bool `subcmd:"progress,true,'display progress of downloads'"`
}
```
Flags represents the flags that are used to control the crawl.


### Type Resources
```go
type Resources struct {
	Extractors      map[content.Type]outlinks.Extractor
	PopulateCrawlFS func(ctx context.Context, cfg config.CrawlService, factories map[string]crawlcmd.FSFactory) error
	NewContentFS    func(ctx context.Context, cfg crawlcmd.CrawlCacheConfig) (content.FS, error)
}
```
Resources represents the resources that are used by the crawler.






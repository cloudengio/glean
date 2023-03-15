# Package [cloudeng.io/glean/crawlindex/crawl](https://pkg.go.dev/cloudeng.io/glean/crawlindex/crawl?tab=doc)

```go
import cloudeng.io/glean/crawlindex/crawl
```


## Types
### Type Crawler
```go
type Crawler struct {
	GleanConfig gleancfg.Glean
	Extractors  func() map[content.Type]outlinks.Extractor
	FSForCrawl  func(config.Crawl) map[string]file.FSFactory
}
```
Crawler represents a crawler instance that contains global configuration
information.

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






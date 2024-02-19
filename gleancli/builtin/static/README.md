# Package [cloudeng.io/glean/gleancli/builtin/static](https://pkg.go.dev/cloudeng.io/glean/gleancli/builtin/static?tab=doc)

```go
import cloudeng.io/glean/gleancli/builtin/static
```

Package static provides the set of builtin resources for a given instance of
the Glean CLI.

## Functions
### Func APIExtensions
```go
func APIExtensions(parents ...string) []gleancfg.Extension
```
APIExtensions returns the builtin API related commands.

### Func Extractors
```go
func Extractors() map[content.Type]outlinks.Extractor
```
Extractors represents the set of available outlink extractors.

### Func MustDocumentConverters
```go
func MustDocumentConverters() *content.Registry[converters.Document]
```
MustDocumentConverters returns the available converters, it will panic on
encountering an error.

### Func MustUserConverters
```go
func MustUserConverters() *content.Registry[converters.User]
```
MustUserConverters returns the available converters, it will panic on
encountering an error.



## Types
### Type CrawlProcessors
```go
type CrawlProcessors struct {
	Extractors map[content.Type]outlinks.Extractor
}
```
CrawlProcessors represents the set of available Extractors for crawling.

### Functions

```go
func MustCrawlProcessors() CrawlProcessors
```




### Type IndexProcessors
```go
type IndexProcessors struct {
	DocumentConverters *content.Registry[converters.Document]
	UserConverters     *content.Registry[converters.User]
}
```
IndexProcessor represents the set of available converters for indexing.

### Functions

```go
func MustIndexProcessors() IndexProcessors
```








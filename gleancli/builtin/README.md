# Package [cloudeng.io/glean/gleancli/builtin](https://pkg.go.dev/cloudeng.io/glean/gleancli/builtin?tab=doc)

```go
import cloudeng.io/glean/gleancli/builtin
```

Package builtin provides the set of builtin resources for a given instance
of the Glean CLI.

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

### Func FSForCrawl
```go
func FSForCrawl(cfg config.Crawl) map[string]file.FSFactory
```
FSForCrawl returns a map of filesystem schemes to filesystem factories for
use when creating crawling requests.

### Func MustConverters
```go
func MustConverters() *content.Registry[converters.T]
```
    MustConverters returns the available converters, it will panic

on encountering an error.





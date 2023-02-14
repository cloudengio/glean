# Package [cloudeng.io/glean/gleancli/builtin](https://pkg.go.dev/cloudeng.io/glean/gleancli/builtin?tab=doc)

```go
import cloudeng.io/glean/gleancli/builtin
```

Package builtin provides the set of builtin resources for a given instance
of the Glean CLI.

## Functions
### Func Converters
```go
func Converters(cfg config.Converters) (*content.Registry[converters.T], error)
```
Converters returns the converters for the given configuration.

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





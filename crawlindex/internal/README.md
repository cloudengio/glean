# Package [cloudeng.io/glean/crawlindex/internal](https://pkg.go.dev/cloudeng.io/glean/crawlindex/internal?tab=doc)

```go
import cloudeng.io/glean/crawlindex/internal
```


## Functions
### Func ParseGleanError
```go
func ParseGleanError(r *http.Response, err error) error
```
ParseGleanError parses the error returned by the glean API and returns a
more useful error message.





# Package [cloudeng.io/glean/crawlindex/internal](https://pkg.go.dev/cloudeng.io/glean/crawlindex/internal?tab=doc)

```go
import cloudeng.io/glean/crawlindex/internal
```


## Functions
### Func NewClient
```go
func NewClient(ctx context.Context, domain string, token *apitokens.T) (context.Context, *gleanclientsdk.APIClient)
```

### Func NewIndexingClient
```go
func NewIndexingClient(ctx context.Context, domain string, token *apitokens.T) (context.Context, *gleansdk.APIClient)
```

### Func ParseGleanError
```go
func ParseGleanError(r *http.Response, err error) error
```
ParseGleanError parses the error returned by the glean API and returns a
more useful error message.





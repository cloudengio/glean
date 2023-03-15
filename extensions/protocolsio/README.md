# Package [cloudeng.io/glean/extensions/protocolsio](https://pkg.go.dev/cloudeng.io/glean/extensions/protocolsio?tab=doc)

```go
import cloudeng.io/glean/extensions/protocolsio
```


## Functions
### Func Extension
```go
func Extension(parents ...string) gleancfg.Extension
```

### Func NewConverter
```go
func NewConverter() converters.T
```



## Types
### Type CommonFlags
```go
type CommonFlags struct {
	config.FileFlags
}
```


### Type CrawlFlags
```go
type CrawlFlags struct {
	protocolsiocmd.CrawlFlags
	CommonFlags
}
```


### Type GetFlags
```go
type GetFlags struct {
	protocolsiocmd.GetFlags
	config.DatasourceName
	CommonFlags
}
```


### Type ScanFlags
```go
type ScanFlags struct {
	protocolsiocmd.ScanFlags
	CommonFlags
}
```






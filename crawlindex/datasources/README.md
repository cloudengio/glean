# Package [cloudeng.io/glean/crawlindex/datasources](https://pkg.go.dev/cloudeng.io/glean/crawlindex/datasources?tab=doc)

```go
import cloudeng.io/glean/crawlindex/datasources
```


## Types
### Type Flags
```go
type Flags struct {
	config.FileFlags
}
```


### Type T
```go
type T struct {
	GleanConfig gleancfg.Glean
}
```

### Methods

```go
func (d *T) Download(ctx context.Context, instance, datasource string) error
```


```go
func (d *T) Register(ctx context.Context, fv *Flags, datasource string) error
```


```go
func (d *T) ShowConfig(ctx context.Context, filename string) error
```








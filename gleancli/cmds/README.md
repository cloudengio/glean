# Package [cloudeng.io/glean/gleancli/cmds](https://pkg.go.dev/cloudeng.io/glean/gleancli/cmds?tab=doc)

```go
import cloudeng.io/glean/gleancli/cmds
```


## Functions
### Func MustNew
```go
func MustNew(options Options) *subcmd.CommandSetYAML
```



## Types
### Type BulkFlags
```go
type BulkFlags struct {
	index.BulkFlags
	crawlindex.AuthFileFlag
}
```


### Type Crawl
```go
type Crawl struct {
	Options
}
```

### Methods

```go
func (cmd *Crawl) Run(ctx context.Context, values interface{}, args []string) error
```




### Type Datasources
```go
type Datasources struct {
	Options
	// contains filtered or unexported fields
}
```

### Methods

```go
func (ds Datasources) Download(ctx context.Context, values any, args []string) error
```


```go
func (ds Datasources) ExplainConfig(_ context.Context, _ interface{}, _ []string) error
```


```go
func (ds Datasources) Register(ctx context.Context, values interface{}, args []string) error
```


```go
func (ds Datasources) ShowConfig(ctx context.Context, _ interface{}, args []string) error
```




### Type DeleteAllFlags
```go
type DeleteAllFlags struct {
	index.DeleteAllFlags
	crawlindex.AuthFileFlag
}
```


### Type DeleteFlags
```go
type DeleteFlags struct {
	index.DeleteFlags
	crawlindex.AuthFileFlag
}
```


### Type DownloadFlags
```go
type DownloadFlags struct {
	crawlindex.AuthFileFlag
}
```


### Type Index
```go
type Index struct {
	Options
	// contains filtered or unexported fields
}
```


### Type Options
```go
type Options struct {
	extensions.StaticResources

	extensions.DynamicResources

	Extensions    []extensions.Extension
	APIExtensions []extensions.Extension

	InitContext func(ctx context.Context) (context.Context, error)
}
```


### Type ProcessNowFlags
```go
type ProcessNowFlags struct {
	index.ProcessNowFlags
	crawlindex.AuthFileFlag
}
```


### Type QueryFlags
```go
type QueryFlags struct {
	index.QueryFlags
	crawlindex.AuthFileFlag
}
```


### Type RegisterFlags
```go
type RegisterFlags struct {
	config.FileFlags
	crawlindex.AuthFileFlag
}
```


### Type StatsFlags
```go
type StatsFlags struct {
	index.StatsFlags
	crawlindex.AuthFileFlag
}
```






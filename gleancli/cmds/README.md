# Package [cloudeng.io/glean/gleancli/cmds](https://pkg.go.dev/cloudeng.io/glean/gleancli/cmds?tab=doc)

```go
import cloudeng.io/glean/gleancli/cmds
```


## Functions
### Func MustNew
```go
func MustNew(opts ...Option) *subcmd.CommandSetYAML
```



## Types
### Type Crawl
```go
type Crawl struct {
	// contains filtered or unexported fields
}
```

### Methods

```go
func (cmd *Crawl) Run(ctx context.Context, values interface{}, args []string) error
```




### Type Datasources
```go
type Datasources struct{}
```

### Methods

```go
func (ds Datasources) Download(ctx context.Context, values interface{}, args []string) error
```


```go
func (ds Datasources) Register(ctx context.Context, values interface{}, args []string) error
```


```go
func (ds Datasources) ShowConfig(ctx context.Context, values interface{}, args []string) error
```




### Type GlobalFlags
```go
type GlobalFlags struct {
	gleancfg.GleanFlags
}
```


### Type Index
```go
type Index struct {
	// contains filtered or unexported fields
}
```


### Type Option
```go
type Option func(o *options)
```

### Functions

```go
func WithAPIExtensions(extensions []gleancfg.Extension) Option
```


```go
func WithConverters(converters *content.Registry[converters.T]) Option
```


```go
func WithFSForCrawl(fsForCrawl func(cfg config.Crawl) map[string]file.FSFactory) Option
```


```go
func WithGleanConfig(cfg gleancfg.Glean) Option
```


```go
func WithOutlinkExtractors(extractors func() map[content.Type]outlinks.Extractor) Option
```








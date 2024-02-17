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
	// contains filtered or unexported fields
}
```

### Methods

```go
func (ds Datasources) Download(ctx context.Context, _ interface{}, args []string) error
```


```go
func (ds Datasources) ExplainConfig(ctx context.Context, _ interface{}, _ []string) error
```


```go
func (ds Datasources) Register(ctx context.Context, values interface{}, args []string) error
```


```go
func (ds Datasources) ShowConfig(ctx context.Context, _ interface{}, args []string) error
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
	Options
}
```


### Type Option
```go
type Option func(o *Options)
```


### Type Options
```go
type Options struct {
	CrawlProcessors static.CrawlProcessors
	IndexProcessors static.IndexProcessors
	CreateFS        func(name string, cfg yaml.Node, factories map[string]crawlcmd.FSFactory) error
	CreateContentFS func(ctx context.Context, path string, cfg yaml.Node) (content.FS, error)
	APIExtensions   []gleancfg.Extension
}
```






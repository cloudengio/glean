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
}
```


### Type DeleteFlags
```go
type DeleteFlags struct {
	index.DeleteFlags
}
```


### Type DownloadFlags
```go
type DownloadFlags struct {
	config.FileFlags
}
```


### Type Index
```go
type Index struct {
	Options
	// contains filtered or unexported fields
}
```


### Type LocalKeysAndLogging
```go
type LocalKeysAndLogging struct {
	LocalKeysAndLoggingFlags
}
```

### Methods

```go
func (lf *LocalKeysAndLogging) InitContext(ctx context.Context) (context.Context, error)
```


```go
func (lf *LocalKeysAndLogging) RegisterGlobalFlags(cmdset *subcmd.CommandSetYAML)
```




### Type LocalKeysAndLoggingFlags
```go
type LocalKeysAndLoggingFlags struct {
	cmdutil.LoggingFlags
	LocalKeyFile string `subcmd:"local-key-file,,'local cleartext file to use for retrieving keys'"`
}
```


### Type Options
```go
type Options struct {
	extensions.StaticResources

	extensions.DynamicResources

	Extensions    []extensions.Extension
	APIExtensions []extensions.Extension

	PlatformSpecific PlatformSpecificConfig
}
```


### Type PlatformSpecificConfig
```go
type PlatformSpecificConfig interface {
	RegisterGlobalFlags(*subcmd.CommandSetYAML)
	InitContext(context.Context) (context.Context, error)
}
```


### Type ProcessNowFlags
```go
type ProcessNowFlags struct {
	index.ProcessNowFlags
}
```


### Type QueryFlags
```go
type QueryFlags struct {
	index.QueryFlags
}
```


### Type RegisterFlags
```go
type RegisterFlags struct {
	config.FileFlags
}
```


### Type StatsFlags
```go
type StatsFlags struct {
	index.StatsFlags
}
```






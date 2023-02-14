# Package [cloudeng.io/glean/gleancli/cmds](https://pkg.go.dev/cloudeng.io/glean/gleancli/cmds?tab=doc)

```go
import cloudeng.io/glean/gleancli/cmds
```


## Variables
### GlobalConfig
```go
GlobalConfig config.Glean

```



## Functions
### Func MustNew
```go
func MustNew() *subcmd.CommandSetYAML
```



## Types
### Type Crawl
```go
type Crawl struct{}
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
	config.GleanFlags
}
```


### Type Index
```go
type Index struct{}
```






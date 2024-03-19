# Package [cloudeng.io/glean/gleancli/runner/runnercmd](https://pkg.go.dev/cloudeng.io/glean/gleancli/runner/runnercmd?tab=doc)

```go
import cloudeng.io/glean/gleancli/runner/runnercmd
```


## Functions
### Func NewCmdSet
```go
func NewCmdSet(spec string, cmds []CommandSpec, globalFlagValues any) *subcmd.CommandSetYAML
```

### Func TemplateFuncs
```go
func TemplateFuncs(authFiles map[string]string) template.FuncMap
```
TemplateFuncs returns a template.FuncMap that contains an AuthFile function
that provides access to the supplied authFiles map.



## Types
### Type CommandSpec
```go
type CommandSpec struct {
	Name       string
	FlagValues any
	Runner     subcmd.Runner
}
```
CommandSpec represents a simple top-level command that can be added to a
subcmd.CommandSetYAML.


### Type CrawlFlags
```go
type CrawlFlags struct {
	ProcessDownloads bool `subcmd:"process-downloads,true,'process downloaded files'"`
}
```


### Type T
```go
type T struct {
	Datasources          []string
	CrawlCommands        map[string][]string
	ProcessCommands      map[string][]string
	IndexCommands        map[string][]string
	TestCacheCommands    map[string][]string
	AuthFiles            map[string]string
	GlobalExecOpts       []cmdexec.Option
	DatasourceConfigFile string
}
```

### Methods

```go
func (c *T) Crawl(ctx context.Context, values interface{}, args []string) error
```


```go
func (c *T) CrawlAll(ctx context.Context, values interface{}, _ []string) error
```


```go
func (c *T) CrawlIndex(ctx context.Context, _ interface{}, args []string) error
```


```go
func (c *T) Index(ctx context.Context, _ interface{}, args []string) error
```


```go
func (c *T) IndexAll(ctx context.Context, _ interface{}, _ []string) error
```


```go
func (c *T) NewRunner(datasource string) *cmdexec.Runner
```


```go
func (c *T) NewRunnerOpts(datasource string) []cmdexec.Option
```


```go
func (c *T) RunCommands(ctx context.Context, datasource string, cmdsets ...map[string][]string) error
```


```go
func (c *T) ShowAll(_ context.Context, _ interface{}, _ []string) error
```


```go
func (c *T) Spec() (string, []CommandSpec)
```


```go
func (c *T) TestCache(ctx context.Context, _ interface{}, args []string) error
```


```go
func (c *T) TestCacheAll(ctx context.Context, _ interface{}, _ []string) error
```




### Type TemplateVars
```go
type TemplateVars struct {
	DatasourceName       string
	DatasourceConfigFile string
}
```
TemplateVars represents the variables that can be accessed from templates.






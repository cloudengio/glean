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


### Type CrawlIndexFlags
```go
type CrawlIndexFlags struct {
	CrawlFlags
	IndexFlags
}
```


### Type IndexFlags
```go
type IndexFlags struct {
	IndexingDryRun bool `subcmd:"indexing-dry-run,false,'dry run only'"`
}
```


### Type T
```go
type T struct {
	Datasources          []string
	CrawlCommands        map[string][]string
	ProcessCommands      map[string][]string
	IndexCommands        map[string][]string
	IndexStatsCommands   map[string][]string
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
func (c *T) CrawlIndex(ctx context.Context, values interface{}, args []string) error
```


```go
func (c *T) CrawlIndexAll(ctx context.Context, values interface{}, _ []string) error
```


```go
func (c *T) Index(ctx context.Context, values interface{}, args []string) error
```


```go
func (c *T) IndexAll(ctx context.Context, values interface{}, _ []string) error
```


```go
func (c *T) IndexingStats(ctx context.Context, values interface{}, args []string) error
```


```go
func (c *T) NewRunner(datasource string, flags any) *cmdexec.Runner
```


```go
func (c *T) NewRunnerOpts(datasource string, flags any) []cmdexec.Option
```


```go
func (c *T) RunCommands(ctx context.Context, datasource string, flags any, cmdsets ...map[string][]string) error
```
RunCommands runs the supplied commands for the specified datasource.
Flags represents the flags made available as template variables.


```go
func (c *T) ShowAll(_ context.Context, _ interface{}, _ []string) error
```


```go
func (c *T) Spec() (string, []CommandSpec)
```


```go
func (c *T) TestCache(ctx context.Context, values interface{}, args []string) error
```


```go
func (c *T) TestCacheAll(ctx context.Context, values interface{}, _ []string) error
```




### Type TemplateVars
```go
type TemplateVars struct {
	DatasourceName       string
	DatasourceConfigFile string
	Flags                any
}
```
TemplateVars represents the variables that can be accessed from templates.






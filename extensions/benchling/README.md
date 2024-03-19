# Package [cloudeng.io/glean/extensions/benchling](https://pkg.go.dev/cloudeng.io/glean/extensions/benchling?tab=doc)

```go
import cloudeng.io/glean/extensions/benchling
```


## Variables
### ExtensionSpec
```go
ExtensionSpec = extensions.ExtensionSpec{
	Name:       cmdName,
	CmdSpec:    cmdSpec,
	AuthCfg:    extensions.APIKey{},
	ServiceCfg: benchlingcmd.Service{},
	AddFunc:    AddExtension,
}

```



## Functions
### Func AddExtension
```go
func AddExtension(extension extensions.Extension, cmdSet *subcmd.CommandSetYAML, parents []string) error
```

### Func NewDocumentConverter
```go
func NewDocumentConverter() converters.Document
```

### Func NewUserConverter
```go
func NewUserConverter() converters.User
```



## Types
### Type CommonFlags
```go
type CommonFlags struct {
	config.FileFlags
	AuthFile string `subcmd:"benchling-auth,$HOME/.benchling.yaml,'benchling.io auth config file'"`
}
```


### Type CrawlFlags
```go
type CrawlFlags struct {
	CommonFlags
	benchlingcmd.CrawlFlags
	Entities string `subcmd:"entities,'users,entries,folders,projects','comma separated list of entities to crawl: users, entries, folders, projects'"`
}
```


### Type IndexFlags
```go
type IndexFlags struct {
	CommonFlags
	benchlingcmd.IndexFlags
}
```






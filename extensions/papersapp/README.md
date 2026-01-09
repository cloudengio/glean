# Package [cloudeng.io/glean/extensions/papersapp](https://pkg.go.dev/cloudeng.io/glean/extensions/papersapp?tab=doc)

```go
import cloudeng.io/glean/extensions/papersapp
```


## Variables
### ExtensionSpec
```go
ExtensionSpec = extensions.ExtensionSpec{
	Name:       cmdName,
	CmdSpec:    cmdSpec,
	ServiceCfg: papersappcmd.Service{},
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
	CommonFlags
	papersappcmd.CrawlFlags
}
```


### Type ScanFlags
```go
type ScanFlags struct {
	CommonFlags
	papersappcmd.ScanFlags
}
```






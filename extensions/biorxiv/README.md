# Package [cloudeng.io/glean/extensions/biorxiv](https://pkg.go.dev/cloudeng.io/glean/extensions/biorxiv?tab=doc)

```go
import cloudeng.io/glean/extensions/biorxiv
```

Package bioxrv provides gleancli extensions for bioRxiv and medRxiv.

## Variables
### ExtensionSpec
```go
ExtensionSpec = extensions.ExtensionSpec{
	Name:       cmdName,
	CmdSpec:    cmdSpec,
	AuthCfg:    struct{}{},
	ServiceCfg: biorxivcmd.Service{},
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
	biorxivcmd.CrawlFlags
}
```


### Type IndexFlags
```go
type IndexFlags struct {
	CommonFlags
	biorxivcmd.IndexFlags
}
```


### Type LookupFlags
```go
type LookupFlags struct {
	CommonFlags
	biorxivcmd.LookupFlags
}
```


### Type ScanFlags
```go
type ScanFlags struct {
	CommonFlags
	biorxivcmd.ScanFlags
}
```






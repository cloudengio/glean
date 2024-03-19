# Package [cloudeng.io/glean/extensions/protocolsio](https://pkg.go.dev/cloudeng.io/glean/extensions/protocolsio?tab=doc)

```go
import cloudeng.io/glean/extensions/protocolsio
```


## Variables
### ExtensionSpec
```go
ExtensionSpec = extensions.ExtensionSpec{
	Name:       cmdName,
	CmdSpec:    cmdSpec,
	AuthCfg:    extensions.APIKey{},
	ServiceCfg: protocolsiocmd.Service{},
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
	AuthFile string `subcmd:"protocolsio-auth,$HOME/.protocolsio.yaml,'protocols.io auth config file'"`
}
```


### Type CrawlFlags
```go
type CrawlFlags struct {
	protocolsiocmd.CrawlFlags
	CommonFlags
}
```


### Type GetFlags
```go
type GetFlags struct {
	protocolsiocmd.GetFlags
	config.DatasourceName
	CommonFlags
}
```


### Type ScanFlags
```go
type ScanFlags struct {
	protocolsiocmd.ScanFlags
	CommonFlags
}
```






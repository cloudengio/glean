# Package [cloudeng.io/glean/extensions/testcmd](https://pkg.go.dev/cloudeng.io/glean/extensions/testcmd?tab=doc)

```go
import cloudeng.io/glean/extensions/testcmd
```


## Variables
### ExtensionSpec
```go
ExtensionSpec = extensions.ExtensionSpec{
	Name:       cmdName,
	CmdSpec:    cmdSpec,
	AuthCfg:    struct{}{},
	ServiceCfg: struct{}{},
	AddFunc:    AddExtension,
}

```



## Functions
### Func AddExtension
```go
func AddExtension(extension extensions.Extension, cmdSet *subcmd.CommandSetYAML, parents []string) error
```



## Types
### Type CacheFlags
```go
type CacheFlags struct {
	CommonFlags
	Items int `subcmd:"show-n,100,'show this many items from the cache'"`
}
```


### Type CommonFlags
```go
type CommonFlags struct {
	config.FileFlags
}
```






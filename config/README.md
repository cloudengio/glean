# Package [cloudeng.io/glean/config](https://pkg.go.dev/cloudeng.io/glean/config?tab=doc)

```go
import cloudeng.io/glean/config
```

Package extensions contains a set of subcmd extensions that can be used by
gleancli implementations. Each such extension is a submodule so that its
dependencies are imported only by gleancli implementations that use it.

## Types
### Type DatasourceConfig
```go
type DatasourceConfig struct {
	GleanInstance                   string `yaml:"glean_instance"`
	gleansdk.CustomDatasourceConfig `yaml:",inline"`
}
```
DatasourceConfig represents the configuration of the datasource with Glean's
API.


### Type Extension
```go
type Extension interface {
	subcmd.Extension
	SetGleanConfig(*Glean)
	GleanConfig() *Glean
}
```

### Functions

```go
func NewExtension(name, spec string, appendFn func(cmdSet *subcmd.CommandSetYAML) error) Extension
```




### Type Glean
```go
type Glean []struct {
	Name string `yaml:"name"`
	Auth struct {
		BearerToken string `yaml:"token"`
	}
	API struct {
		Domain string `yaml:"domain"`
	}
}
```

### Methods

```go
func (c Glean) NewAPIClient(ctx context.Context, name string) (context.Context, *gleansdk.APIClient, error)
```


```go
func (c Glean) String() string
```




### Type GleanFlags
```go
type GleanFlags struct {
	Config string `subcmd:"config,$HOME/.glean.yaml,'glean config file'"`
}
```






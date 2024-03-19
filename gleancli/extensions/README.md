# Package [cloudeng.io/glean/gleancli/extensions](https://pkg.go.dev/cloudeng.io/glean/gleancli/extensions?tab=doc)

```go
import cloudeng.io/glean/gleancli/extensions
```

Package extensions defines an extension mechanism for gleancli
implementations and utilities for creating, managing and implementing
extensions.

## Functions
### Func FirstAPICrawl
```go
func FirstAPICrawl(crawls apicrawlcmd.Crawls) (apicrawlcmd.Crawl[yaml.Node], bool)
```



## Types
### Type APIKey
```go
type APIKey struct {
	APIKey string `yaml:"api_key" cmd:"API key in apitokens format (ie. scheme://<value>)"`
}
```
APIKey represents an API key in apitokens format.

### Methods

```go
func (k *APIKey) ParseAndRead(ctx context.Context, filename string, readers *apitokens.Readers) (*apitokens.T, error)
```
ParseAndRead reads the API key from the specified file and returns it in
apitokens format.




### Type DynamicResources
```go
type DynamicResources struct {
	PopulateCrawlFS func(ctx context.Context, cfg config.CrawlService, factories map[string]crawlcmd.FSFactory) error

	NewContentFS func(ctx context.Context, cfg crawlcmd.CrawlCacheConfig) (content.FS, error)

	NewCheckpointOp func(ctx context.Context, cfg crawlcmd.CrawlCacheConfig) (checkpoint.Operation, error)

	NewOperationsFS func(ctx context.Context, cfg crawlcmd.CrawlCacheConfig) (operations.FS, error)
}
```
DynamicResources provides a set of functions that can be used to create the
various resources required by commands and command extensions.


### Type Extension
```go
type Extension interface {
	subcmd.Extension
	SetOptions(ExtensionOptions)
	Options() ExtensionOptions
	AuthConfigType() any
	ServiceConfigType() any
}
```
Extension is implemented by all extensions.

### Functions

```go
func NewExtension(spec ExtensionSpec, parents []string) Extension
```
NewExtension creates a new extension from the supplied spec.


```go
func NewExtensions(parents []string, specs ...ExtensionSpec) []Extension
```
NewExtensions creates all of the command extensions specified by specs.




### Type ExtensionOptions
```go
type ExtensionOptions struct {
	StaticResources

	DynamicResources
}
```
ExtensionsOptions are the options that are passed to each extension.

### Methods

```go
func (eo ExtensionOptions) ResourcesForDatasource(ctx context.Context, configFile, authFile, datasource string) (config.Datasource, apicrawlcmd.Resources, error)
```




### Type ExtensionSpec
```go
type ExtensionSpec struct {
	Name       string // Name of the extension
	CmdSpec    string // YAML subcmd spec
	AuthCfg    any    // used for describing the auth configuration
	ServiceCfg any    // used for describing the service configuration
	// called to add the command to its parent cmdSet.
	AddFunc func(extension Extension, cmdSet *subcmd.CommandSetYAML, parents []string) error
}
```
ExtensionSpec defines an extension.


### Type StaticResources
```go
type StaticResources struct {
	DocumentConverters *content.Registry[converters.Document]
	UserConverters     *content.Registry[converters.User]
	Extractors         map[content.Type]outlinks.Extractor

	// TokenReaders provides a set of readers for reading API tokens.
	TokenReaders *apitokens.Readers
}
```
StaticResources provides a set of resources that are typically required
by commands and command extensions that are statically configured per
application instance.






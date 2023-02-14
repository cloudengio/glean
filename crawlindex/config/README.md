# Package [cloudeng.io/glean/crawlindex/config](https://pkg.go.dev/cloudeng.io/glean/crawlindex/config?tab=doc)

```go
import cloudeng.io/glean/crawlindex/config
```


## Types
### Type API
```go
type API struct {
}
```
API is the configuration for an API based crawl. It is currently only a
placeholder.


### Type BulkIndex
```go
type BulkIndex struct {
	ForceDeletion  bool `yaml:"force_deletion"`  // Glean's force deletion flag
	ForceRestart   bool `yaml:"force_restart"`   // Glean's force restart flag
	ReaddirEntries int  `yaml:"readdir_entries"` // number of entries per Readdir call.
}
```
BulkIndex represents the configuration info for a Glean builk index
operation.


### Type Cache
```go
type Cache struct {
	Path string // Path is the location of the cache for this datasource.
}
```
Cache represents a cache configuration.


### Type Conversion
```go
type Conversion struct {
	Type       content.Type
	Converter  *Converter
	Datasource *gleansdk.CustomDatasourceConfig
}
```
Conversion represents the configuration for a single content conversion
operation.


### Type Converter
```go
type Converter struct {
	FromContentType []content.Type `yaml:"from_content_types,flow"` // Content types that this converter can handle.

	ViewURLRewrites []string `yaml:"view_url_rewrites"` // Rewrite rules for viewurls specified as textutil.RewriteRules

	AllowAnonymousAccess bool `yaml:"allow_anonymous_access"` // allow anonymous access to the converted documents.

	// Default author to use if none can be obtained from the document itself.
	DefaultAuthor User `yaml:"default_author"`

	CustomConfig yaml.Node `yaml:"custom"`
}
```
Converter represents the ability to convert from a set of content types to
to a Glean document ("glean/document")


### Type Converters
```go
type Converters []Converter
```
Converters represents a set of converters.


### Type Crawl
```go
type Crawl struct {
	crawlcmd.Config `yaml:",inline"`
	AWS             awsconfig.AWSFlags `yaml:"aws,omitempty"`
}
```
Crawl represents a single crawl that contributes data to a datasource.


### Type Datasource
```go
type Datasource struct {
	Datasource string // Datasource name.

	Crawls []Crawl // File/download orientend Crawls that obtain data for this datasource.
	APIS   []API   // API based 'crawls' that obtain data for this datasource.

	// Bulk index configuration for this datasource.
	*BulkIndex `yaml:"bulk_index,omitempty"`
	// Incremental index configuration for this datasource.
	*IncrementalIndex `yaml:"incremental_index,omitempty"`

	Cache                  // Cache configuration for this datasource.
	Converters []Converter // Converters (from download.Result to Glean document) configuration.

	// The Glean datasource configuration in YAML as opposed to JSON
	// format.
	DatasourceConfig `yaml:"datasource_config"`
}
```
Datasource represents a single datasource or corpus to be crawled and
indexed.

### Functions

```go
func DatasourceForName(filename string, name string) (Datasource, error)
```
DatasourceForName returns the datasource configuration for the named
datasource read from the specified config file.



### Methods

```go
func (d Datasource) ConfigForContentType() map[content.Type]Conversion
```
ConfigForContentType returns a map from content type to all of the
configuration information that pertains to that content type.




### Type DatasourceConfig
```go
type DatasourceConfig struct {
	GleanInstance                   string `yaml:"glean_instance"`
	gleansdk.CustomDatasourceConfig `yaml:",inline"`
}
```
DatasourceConfig represents the configuration of the datasource with Glean's
API.


### Type Datasources
```go
type Datasources []Datasource
```
Datasources represents a list of named datasources.

### Methods

```go
func (d Datasources) ConfigForName(name string) (Datasource, bool)
```
ConfigForName for returns the configuration for the named datasource.




### Type FileFlags
```go
type FileFlags struct {
	ConfigFile string `subcmd:"datasource-configs,,datasource config file"`
}
```
FileFlags represents a command line flag for the datasource config file.


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


### Type IncrementalIndex
```go
type IncrementalIndex struct {
	DeletionDelay time.Duration `yaml:"deletion_delay"` // Documents that have not been updated within deletion delay will be removed the Glean index.
}
```
IncrementalIndex represents the configuration info for incremental, document
at a time, indexing.


### Type User
```go
type User struct {
	Email  string `yaml:"email"`
	UserID string `yaml:"user_id"`
	Name   string `yaml:"name"`
}
```
User represents a user in the system/datasource beinfg indexed.






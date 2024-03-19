# Package [cloudeng.io/glean/gleancli/builtin/static](https://pkg.go.dev/cloudeng.io/glean/gleancli/builtin/static?tab=doc)

```go
import cloudeng.io/glean/gleancli/builtin/static
```

Package static provides the set of builtin resources for a given instance of
the Glean CLI.

## Functions
### Func APIExtensions
```go
func APIExtensions(parents ...string) []extensions.Extension
```
APIExtensions returns the builtin API related commands.

### Func Extensions
```go
func Extensions(parents ...string) []extensions.Extension
```
Extensions returns any top-level commands.

### Func LinkExtractors
```go
func LinkExtractors() map[content.Type]outlinks.Extractor
```
LinkExtractors represents the set of available outlink extractors.

### Func MustDocumentConverters
```go
func MustDocumentConverters() *content.Registry[converters.Document]
```
MustDocumentConverters returns the available converters, it will panic on
encountering an error.

### Func MustUserConverters
```go
func MustUserConverters() *content.Registry[converters.User]
```
MustUserConverters returns the available converters, it will panic on
encountering an error.

### Func New
```go
func New() extensions.StaticResources
```

### Func NewHTML
```go
func NewHTML() converters.Document
```
NewHTML returns a new install of HTML.

### Func TokenReaders
```go
func TokenReaders() *apitokens.Readers
```



## Types
### Type HTML
```go
type HTML struct {
	// contains filtered or unexported fields
}
```
HTML represents an html to glean document converter. It differs from
converters.HTML only in that it sets the object type to statichtml.

### Methods

```go
func (cnv *HTML) Convert(ctx context.Context, datasource string, cfg config.Conversion, ctype content.Type, data []byte) (gleansdk.DocumentDefinition, error)
```


```go
func (cnv *HTML) Type() content.Type
```








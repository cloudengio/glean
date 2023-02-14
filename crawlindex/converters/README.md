# Package [cloudeng.io/glean/crawlindex/converters](https://pkg.go.dev/cloudeng.io/glean/crawlindex/converters?tab=doc)

```go
import cloudeng.io/glean/crawlindex/converters
```


## Constants
### GleanContentType
```go
GleanContentType = content.Type("glean/document")

```



## Functions
### Func CreateRegistry
```go
func CreateRegistry(converters ...T) (*content.Registry[T], error)
```



## Types
### Type HTML
```go
type HTML struct{}
```
HTML represents an html to glean document converter.

### Methods

```go
func (cnv *HTML) GleanDocument(datasource string, cfg config.Conversion, doc crawlindex.Document) (*gleansdk.DocumentDefinition, error)
```


```go
func (cnv *HTML) Type() content.Type
```




### Type T
```go
type T interface {
	Type() content.Type
	GleanDocument(datasource string, cfg config.Conversion, dl crawlindex.Document) (*gleansdk.DocumentDefinition, error)
}
```

### Functions

```go
func NewHTML() T
```
NewHTML returns a new install of HTML.








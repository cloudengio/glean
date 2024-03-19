# Package [cloudeng.io/glean/crawlindex/converters](https://pkg.go.dev/cloudeng.io/glean/crawlindex/converters?tab=doc)

```go
import cloudeng.io/glean/crawlindex/converters
```


## Constants
### GleanDocumentType, GleanUserType
```go
GleanDocumentType = content.Type("glean/document")
GleanUserType = content.Type("glean/user")

```



## Functions
### Func CreateDocumentRegistry
```go
func CreateDocumentRegistry(converters ...Document) (*content.Registry[Document], error)
```

### Func CreateUserRegistry
```go
func CreateUserRegistry(converters ...User) (*content.Registry[User], error)
```

### Func IgnoreContentType
```go
func IgnoreContentType(ctype content.Type) error
```

### Func IsIgnoreContentType
```go
func IsIgnoreContentType(err error) bool
```



## Types
### Type Document
```go
type Document interface {
	Type() content.Type
	Convert(ctx context.Context, datasource string, cfg config.Conversion, ctype content.Type, data []byte) (gleansdk.DocumentDefinition, error)
}
```

### Functions

```go
func NewHTML() Document
```
NewHTML returns a new install of HTML.




### Type HTML
```go
type HTML struct{}
```
HTML represents an html to glean document converter.

### Methods

```go
func (cnv *HTML) Convert(_ context.Context, datasource string, cfg config.Conversion, ctype content.Type, data []byte) (gleansdk.DocumentDefinition, error)
```


```go
func (cnv *HTML) Type() content.Type
```




### Type User
```go
type User interface {
	Type() content.Type
	Convert(ctx context.Context, datasource string, cfg config.Conversion, ctype content.Type, data []byte) (gleansdk.DatasourceUserDefinition, error)
}
```






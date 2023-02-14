# Package [cloudeng.io/glean/crawlindex](https://pkg.go.dev/cloudeng.io/glean/crawlindex?tab=doc)

```go
import cloudeng.io/glean/crawlindex
```


## Functions
### Func WriteDocument
```go
func WriteDocument(v Document, dir, file string) error
```
WriteDocument writes the provided document to the specified file.



## Types
### Type Document
```go
type Document struct {
	Time     time.Time
	Type     content.Type
	Download *download.Result
}
```
Document represents a downloaded document, whether it be a file, a http get
or the result of a Rest API call.

### Functions

```go
func NewDownloadDocument(d *download.Result, ctype content.Type) Document
```


```go
func ReadDocument(dir, file string) (Document, error)
```
ReadDocument reads a Document from the specified file.



### Methods

```go
func (d Document) WriteToFile(dir, file string) error
```








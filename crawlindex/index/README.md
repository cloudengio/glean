# Package [cloudeng.io/glean/crawlindex/index](https://pkg.go.dev/cloudeng.io/glean/crawlindex/index?tab=doc)

```go
import cloudeng.io/glean/crawlindex/index
```


## Types
### Type BulkFlags
```go
type BulkFlags struct {
	config.FileFlags
	UploadID      string `subcmd:"upload-id,upload,id to use for this bulk upload"`
	ForceRestart  bool   `subcmd:"force-restart,false,restart the bulk upload"`
	ForceDeletion bool   `subcmd:"force-sync-deletion,false,synchronously delete stale documents on upload of last bulk indexing batch"`
}
```
BulkFlags represents the flags to the bulk indexing command.


### Type BulkIndexOption
```go
type BulkIndexOption func(o *bulkOptions)
```
BulkIndexOption represents an option to NewBulkIndexer.

### Functions

```go
func WithBulkID(id string) BulkIndexOption
```
WithBulkID specifies a custom id to use for the bulk upload. If one is not
specified the current date and time are used.


```go
func WithForceDelete(forceDelete bool) BulkIndexOption
```
WithForceDelete disables the deletion of too many documents test.


```go
func WithForceRestart(forceRestart bool) BulkIndexOption
```
WithForceRestart sets the force restart options.




### Type Indexer
```go
type Indexer struct {
	GleanConfig config.Glean
	Config      config.Datasource
	Converters  *content.Registry[converters.T]
}
```
Indexer represents a Glean indexer.

### Methods

```go
func (idx *Indexer) Bulk(ctx context.Context, fv *BulkFlags, datasource string) error
```
Bulk indexes a datasource in bulk mode.


```go
func (idx *Indexer) Stats(ctx context.Context, fv *StatsFlags, datasource string) error
```




### Type Request
```go
type Request struct {
	Documents []*gleansdk.DocumentDefinition
	LastPage  bool
}
```
Request contains the documents to be indexed and a flag indicating if these
are the last documents in a bulk index operation. The indexer will stop when
it receives a request with LastPage set. If LastPage flag is never set,
the indexer will assume that the indexing operation is complete when its
input channel is closed.


### Type StatsFlags
```go
type StatsFlags struct {
	config.FileFlags
}
```
StatsFlags represents the flags to the bulk indexing command.






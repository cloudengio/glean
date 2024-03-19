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
	DryRun        bool   `subcmd:"dry-run,false,'process but not do index documents'"`
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
func WithDryRun(v bool) BulkIndexOption
```


```go
func WithForceDelete(forceDelete bool) BulkIndexOption
```
WithForceDelete disables the deletion of too many documents test.


```go
func WithForceRestart(forceRestart bool) BulkIndexOption
```
WithForceRestart sets the force restart options.


```go
func WithReqSizes(doc, emp int) BulkIndexOption
```


```go
func WithUsers(v bool) BulkIndexOption
```




### Type DeleteAllFlags
```go
type DeleteAllFlags struct {
	config.FileFlags
}
```
DeleteAllFlags represents the flags to the indexing delete-all command.


### Type DeleteFlags
```go
type DeleteFlags struct {
	config.FileFlags
	NumDocuments int    `subcmd:"num-documents,100,number of documents to return"`
	Type         string `subcmd:"doc-type,,type of object to delete"`
}
```
DeleteFlags represents the flags to the indexing delete command.


### Type Indexer
```go
type Indexer struct {
	// contains filtered or unexported fields
}
```
Indexer represents a Glean indexer.

### Functions

```go
func New(ctx context.Context, fv config.FileFlags, datasource string, resources Resources) (*Indexer, error)
```



### Methods

```go
func (idx *Indexer) Bulk(ctx context.Context, fv *BulkFlags) error
```
Bulk indexes a datasource in bulk mode.


```go
func (idx *Indexer) Delete(ctx context.Context, fv *DeleteFlags, query string) error
```


```go
func (idx *Indexer) DeleteAll(ctx context.Context, _ *DeleteAllFlags) error
```


```go
func (idx *Indexer) ProcessNow(ctx context.Context, _ *ProcessNowFlags) error
```


```go
func (idx *Indexer) Query(ctx context.Context, fv *QueryFlags, datasource string, query string) error
```


```go
func (idx *Indexer) Stats(ctx context.Context, _ *StatsFlags) error
```




### Type ProcessNowFlags
```go
type ProcessNowFlags struct {
	config.FileFlags
}
```
ProcessNowFlags represents the flags to the indexing process now command.


### Type QueryFlags
```go
type QueryFlags struct {
	config.FileFlags
	NumDocuments       int  `subcmd:"num-documents,10,number of documents to return"`
	ShowIndexingStatus bool `subcmd:"show-indexing-status,true,show indexing status for all returned documents"`
}
```
QueryFlags represents the flags to the indexing query command.


### Type Request
```go
type Request struct {
	Documents []gleansdk.DocumentDefinition
	Users     []gleansdk.DatasourceUserDefinition
	LastPage  bool
}
```
Request contains the entities to be indexed and a flag indicating if these
are the last entities in a bulk index operation. The indexer will stop when
it receives a request with LastPage set. If LastPage flag is never set,
the indexer will assume that the indexing operation is complete when its
input channel is closed.


### Type Resources
```go
type Resources struct {
	IndexingToken      *apitokens.T
	ClientToken        *apitokens.T
	DocumentConverters *content.Registry[converters.Document]
	UserConverters     *content.Registry[converters.User]
	NewOperationsFS    func(ctx context.Context, cfg crawlcmd.CrawlCacheConfig) (operations.FS, error)
}
```
Resources represents the resources needed by an indexer.


### Type StatsFlags
```go
type StatsFlags struct {
	config.FileFlags
}
```
StatsFlags represents the flags to the bulk indexing command.






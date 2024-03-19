# Package [cloudeng.io/glean/gleancli/builtin/dynamic](https://pkg.go.dev/cloudeng.io/glean/gleancli/builtin/dynamic?tab=doc)

```go
import cloudeng.io/glean/gleancli/builtin/dynamic
```

Package dynamic provides dynamically instantiated resources for a given
instance of the Glean CLI. These resources are generally returned as
'factories'.

## Functions
### Func New
```go
func New() extensions.DynamicResources
```
New returns a set of dynamic resources for use with the Glean CLI.

### Func NewCheckpointOp
```go
func NewCheckpointOp(ctx context.Context, cfg crawlcmd.CrawlCacheConfig) (checkpoint.Operation, error)
```
NewCheckpointOp creates a checkpoint.Operation appropriate for the given
crawl checkpoint configuration.

### Func NewContentFS
```go
func NewContentFS(ctx context.Context, cfg crawlcmd.CrawlCacheConfig) (content.FS, error)
```
NewContentFS creates a content.FS appropriate for the given crawl downloads
cache configuration.

### Func NewOperationsFS
```go
func NewOperationsFS(ctx context.Context, cfg crawlcmd.CrawlCacheConfig) (operations.FS, error)
```
NewOperationsFS returns an operations.FS appropriate for the given path and
configuration.

### Func PopulateCrawlFS
```go
func PopulateCrawlFS(_ context.Context, cfg config.CrawlService, factories map[string]crawlcmd.FSFactory) error
```
PopulateCrawlFS fills a map of URI schemes to crawlcmd.FS factories





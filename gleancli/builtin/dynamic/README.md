# Package [cloudeng.io/glean/gleancli/builtin/dynamic](https://pkg.go.dev/cloudeng.io/glean/gleancli/builtin/dynamic?tab=doc)

```go
import cloudeng.io/glean/gleancli/builtin/dynamic
```

Package dynamic provides dynamically instantiated resources for a given
instance of the Glean CLI. These resources are generally returned as
'factories'.

## Functions
### Func ContentFS
```go
func ContentFS(ctx context.Context, path string, cfg yaml.Node) (content.FS, error)
```
ContentFS returns a factory to create a content.FS for the given crawl cache
configuration.

### Func PopulateFS
```go
func PopulateFS(name string, cfg yaml.Node, factories map[string]crawlcmd.FSFactory) error
```
PopulateFS fills a map of URI schemes to crawlcmd.FS factories





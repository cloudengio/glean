# Package [cloudeng.io/glean/gleancli/runner/runnerglobals](https://pkg.go.dev/cloudeng.io/glean/gleancli/runner/runnerglobals?tab=doc)

```go
import cloudeng.io/glean/gleancli/runner/runnerglobals
```


## Types
### Type Flags
```go
type Flags struct {
	WorkingDir           string `subcmd:"working-dir,.,'working directory'"`
	DatasourceConfigFile string `subcmd:"datasource-configs,connectors.yaml,'datasource configuration file'"`
	Executable           string `subcmd:"executable,gleancli,'comma separated list of commands to run for gleancli, eg. go run .'"`
	DryRun               bool   `subcmd:"dry-run,false,'do not actually run the commands'"`
	Verbose              bool   `subcmd:"verbose,false,'print verbose output'"`
}
```

### Methods

```go
func (f *Flags) CmdExecOptions() []cmdexec.Option
```








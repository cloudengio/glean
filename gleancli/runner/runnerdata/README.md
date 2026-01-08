# Package [cloudeng.io/glean/gleancli/runner/runnerdata](https://pkg.go.dev/cloudeng.io/glean/gleancli/runner/runnerdata?tab=doc)

```go
import cloudeng.io/glean/gleancli/runner/runnerdata
```


## Variables
### Datasources, AuthFiles, CrawlCommands, ProcessCommands, BulkIndexCommand, IndexCommands, IndexStatsCommand, IndexStatsCommands, TestCacheCommands
```go
Datasources = []string{
	"biorxiv.org",
	"protocols.io",
	"papersapp.com",
	"benchling.com",
}
AuthFiles = map[string]string{
	"biorxiv.org":   "$HOME/.biorxiv.yaml",
	"protocols.io":  "$HOME/.protocolsio.yaml",
	"papersapp.com": "$HOME/.papersapp.yaml",
	"benchling.com": "$HOME/.benchling.yaml",
	"glean.com":     "$HOME/.glean.yaml",
}
CrawlCommands = map[string][]string{
	"biorxiv.org": {"api", "biorxiv", "crawl",
		"--datasource-configs={{.DatasourceConfigFile}}",
		"{{.DatasourceName}}"},
	"protocols.io": {"api", "protocols.io", "crawl",
		`--protocolsio-auth={{AuthFile "protocols.io"}}`,
		"--datasource-configs={{.DatasourceConfigFile}}",
		"{{.DatasourceName}}"},
	"papersapp.com": {"api", "papersapp", "crawl",
		`--papersapp-auth={{AuthFile "papersapp.com"}}`,
		"--datasource-configs={{.DatasourceConfigFile}}",
		"{{.DatasourceName}}"},
	"benchling.com": {"api", "benchling", "crawl",
		`--benchling-auth={{AuthFile "benchling.com"}}`,
		"--datasource-configs={{.DatasourceConfigFile}}",
		"{{.DatasourceName}}"},
	"files": {"crawl", "run",
		"--datasource-configs={{.DatasourceConfigFile}}",
		"{{.DatasourceName}}"},
}
ProcessCommands = map[string][]string{
	"biorxiv.org":   {},
	"protocols.io":  {},
	"papersapp.com": {},
	"benchling.com": {"api", "benchling", "create-indexable-documents",
		`--benchling-auth={{AuthFile "benchling.com"}}`,
		"--datasource-configs={{.DatasourceConfigFile}}",
		"{{.DatasourceName}}"},
	"files": {},
}
BulkIndexCommand = []string{
	"index", "bulk",
	"--dry-run={{.Flags.IndexingDryRun}}",
	`--glean-auth={{AuthFile "glean.com"}}`,
	"--datasource-configs={{.DatasourceConfigFile}}",
	"{{.DatasourceName}}",
}
IndexCommands = map[string][]string{
	"biorxiv.org":   BulkIndexCommand,
	"papersapp.com": BulkIndexCommand,
	"benchling.com": BulkIndexCommand,
	"protocols.io":  BulkIndexCommand,
	"files":         BulkIndexCommand,
}
IndexStatsCommand = []string{
	"index", "stats",
	`--glean-auth={{AuthFile "glean.com"}}`,
	"--datasource-configs={{.DatasourceConfigFile}}",
	"{{.DatasourceName}}",
}
IndexStatsCommands = map[string][]string{
	"biorxiv.org":   IndexStatsCommand,
	"papersapp.com": IndexStatsCommand,
	"benchling.com": IndexStatsCommand,
	"protocols.io":  IndexStatsCommand,
	"files":         IndexStatsCommand,
}
TestCacheCommands = map[string][]string{
	"biorxiv.org": {"test", "cache",
		"--datasource-configs={{.DatasourceConfigFile}}",
		"{{.DatasourceName}}"},
	"protocols.io": {"test", "cache",
		"--datasource-configs={{.DatasourceConfigFile}}",
		"{{.DatasourceName}}"},
	"papersapp.com": {"test", "cache",
		"--datasource-configs={{.DatasourceConfigFile}}",
		"{{.DatasourceName}}"},
	"benchling.com": {"test", "cache",
		"--datasource-configs={{.DatasourceConfigFile}}",
		"{{.DatasourceName}}"},
}

```





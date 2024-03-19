# Package [cloudeng.io/glean/crawlindex](https://pkg.go.dev/cloudeng.io/glean/crawlindex?tab=doc)

```go
import cloudeng.io/glean/crawlindex
```


## Types
### Type Auth
```go
type Auth []struct {
	Domain string `yaml:"domain" cmd:"domain name of the glean instance"`
	Auth   struct {
		BearerToken       string `yaml:"indexing_token" cmd:"indexing token for the glean instance"`
		ClientBearerToken string `yaml:"client_token" cmd:"client bearer token for the glean instance"`
	}
}
```

### Methods

```go
func (a Auth) TokensForDomain(domain string) (indexingToken, clientToken *apitokens.T)
```




### Type AuthFileFlag
```go
type AuthFileFlag struct {
	AuthFile string `subcmd:"glean-auth,$HOME/.glean-auth.yaml,'file containing authentication tokens for glean instances'"`
}
```






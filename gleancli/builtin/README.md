# Package [cloudeng.io/glean/gleancli/builtin](https://pkg.go.dev/cloudeng.io/glean/gleancli/builtin?tab=doc)

```go
import cloudeng.io/glean/gleancli/builtin
```

Package builtin provides the set of builtin resources for a given instance
of the Glean CLI. These resources fall into two categories:
 1. static, these can be instantiated once at init time.
 2. dynamic, these are instantiate dynamically, usually by other packages
    when the application is running. They are always returned as 'factories'
    that can be used to create the actual resources.



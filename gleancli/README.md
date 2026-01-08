# [cloudeng.io/glean/gleancli](https://pkg.go.dev/cloudeng.io/glean/gleancli?tab=doc)


Usage of `gleancli`

    command line for working with the Glean API

    datasources - register or list Glean SDK data sources
          crawl - crawl a web site or filesystem
            api - API releated commands
          index - crawl a web site or filesystem
           test - test configuration for the glean cli

global flags: [--keychain-account= --keychain-item= --keychain-plugin=
--keychain-type=all --local-key-file= --log-file= --log-format=json
--log-level=0 --log-source-code=false]

    -keychain-account string
      account that the keychain item belongs to
    -keychain-item string
      name of the macOS keychain item to use for storing/retrieving tokens
    -keychain-plugin string
      path to the plugin binary
    -keychain-type value
      the type of keychain plugin to use: file, data-protection, icloud or all
      (default all)
    -local-key-file string
      local cleartext file to use for retrieving keys
    -log-file string
      log file path. If not specified logs are written to stderr, if set to -
      logs are written to stdout
    -log-format string
      log format: text or json (default "json")
    -log-level int
      logging level: 0=error, 1=warn, 2=info, 3=debug
    -log-source-code
      include source code file and line number in logs


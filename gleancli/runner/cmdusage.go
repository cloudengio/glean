// Usage of runner
//
//	         crawl - crawl a datasource, optionally processing the results
//	     crawl-all - crawl all configured datasources concurrently
//	         index - index a datasource
//	     index-all - index all configured datasources concurrently
//	   crawl-index - crawl and index a datasource
//	      show-all - show all commands
//	    test-cache - test the cache configuration for a datasource and display the first n items therein
//	test-cache-all - test the cache configuration for all datasources
//
// global flags: [--datasource-configs=connectors.yaml --dry-run=false --executable=gleancli --verbose=false --working-dir=.]
//
//	-datasource-configs string
//	  datasource configuration file (default "connectors.yaml")
//	-dry-run
//	  do not actually run the commands
//	-executable string
//	  comma separated list of commands to run for gleancli, eg. go run .
//	  (default "gleancli")
//	-verbose
//	  print verbose output
//	-working-dir string
//	  working directory (default ".")
package main

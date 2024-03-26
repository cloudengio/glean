// Copyright 2024 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package runnercmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"text/template"

	"cloudeng.io/cmdutil/cmdexec"
	"cloudeng.io/cmdutil/subcmd"
	"cloudeng.io/errors"
	"cloudeng.io/sync/errgroup"
)

type T struct {
	Datasources          []string
	CrawlCommands        map[string][]string
	ProcessCommands      map[string][]string
	IndexCommands        map[string][]string
	IndexStatsCommands   map[string][]string
	TestCacheCommands    map[string][]string
	AuthFiles            map[string]string
	GlobalExecOpts       []cmdexec.Option
	DatasourceConfigFile string
}

const cmdSpec = `name: runner
commands:
  - name: crawl
    summary: crawl a datasource, optionally processing the results
    arguments:
      - name - datasource
  - name: crawl-all
    summary: crawl all configured datasources concurrently
  - name: index
    summary: index a datasource
    arguments:
      - name - datasource
  - name: index-all
    summary: index all configured datasources concurrently
  - name: crawl-index
    summary: crawl and index a datasource
    arguments:
      - name - datasource
  - name: crawl-index-all
    summary: crawl and index all datasources
  - name: show-all
    summary: show all commands
  - name: test-cache
    summary: test the cache configuration for a datasource and display the first n items therein
    arguments:
      - name - datasource
  - name: test-cache-all
    summary: test the cache configuration for all datasources
  - name: indexing-stats
    summary: indexing stats
    arguments:
      - datasource... - datasource to display stats for
`

type CrawlFlags struct {
	ProcessDownloads bool `subcmd:"process-downloads,true,'process downloaded files'"`
}

type IndexFlags struct {
	IndexingDryRun bool `subcmd:"indexing-dry-run,false,'dry run only'"`
}

type CrawlIndexFlags struct {
	CrawlFlags
	IndexFlags
}

// CommandSpec represents a simple top-level command that can
// be added to a subcmd.CommandSetYAML.
type CommandSpec struct {
	Name       string
	FlagValues any
	Runner     subcmd.Runner
}

func (c *T) Spec() (string, []CommandSpec) {
	return cmdSpec, []CommandSpec{
		{
			Name:       "crawl",
			FlagValues: &CrawlFlags{},
			Runner:     c.Crawl,
		},
		{
			Name:       "crawl-all",
			FlagValues: &CrawlFlags{},
			Runner:     c.CrawlAll,
		},
		{
			Name:       "index",
			FlagValues: &IndexFlags{},
			Runner:     c.Index,
		},
		{
			Name:       "index-all",
			FlagValues: &IndexFlags{},
			Runner:     c.IndexAll,
		},
		{
			Name:       "crawl-index",
			FlagValues: &CrawlIndexFlags{},
			Runner:     c.CrawlIndex,
		},
		{
			Name:       "crawl-index-all",
			FlagValues: &CrawlIndexFlags{},
			Runner:     c.CrawlIndexAll,
		},
		{
			Name:       "show-all",
			FlagValues: &struct{}{},
			Runner:     c.ShowAll,
		},
		{
			Name:       "test-cache",
			FlagValues: &struct{}{},
			Runner:     c.TestCache,
		},
		{
			Name:       "test-cache-all",
			FlagValues: &struct{}{},
			Runner:     c.TestCacheAll,
		},
		{
			Name:       "indexing-stats",
			FlagValues: &struct{}{},
			Runner:     c.IndexingStats,
		},
	}
}

func (c *T) NewRunnerOpts(datasource string, flags any) []cmdexec.Option {
	opts := append([]cmdexec.Option{}, c.GlobalExecOpts...)
	vars := TemplateVars{
		DatasourceName:       datasource,
		DatasourceConfigFile: c.DatasourceConfigFile,
		Flags:                flags,
	}
	opts = append(opts,
		cmdexec.WithTemplateFuncs(TemplateFuncs(c.AuthFiles)),
		cmdexec.WithTemplateVars(vars))
	return opts
}

func (c *T) NewRunner(datasource string, flags any) *cmdexec.Runner {
	return cmdexec.New(datasource, c.NewRunnerOpts(datasource, flags)...)
}

// RunCommands runs the supplied commands for the specified datasource.
// Flags represents the flags made available as template variables.
func (c *T) RunCommands(ctx context.Context, datasource string, flags any, cmdsets ...map[string][]string) error {
	for _, cmdmap := range cmdsets {
		cmds, ok := cmdmap[datasource]
		if !ok {
			return fmt.Errorf("no commands for datasource: %v", datasource)
		}
		if err := c.NewRunner(datasource, flags).Run(ctx, cmds...); err != nil {
			return err
		}
	}
	return nil
}

func (c *T) Crawl(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*CrawlFlags)
	return c.crawlAndProcess(ctx, fv, args[0])
}

func (c *T) crawlAndProcess(ctx context.Context, fv *CrawlFlags, datasource string) error {
	if err := c.RunCommands(ctx, datasource, fv, c.CrawlCommands); err != nil {
		return err
	}
	if !fv.ProcessDownloads || len(c.ProcessCommands[datasource]) == 0 {
		return nil
	}
	return c.RunCommands(ctx, datasource, fv, c.ProcessCommands)
}

func (c *T) CrawlAll(ctx context.Context, values interface{}, _ []string) error {
	fv := values.(*CrawlFlags)
	var g errgroup.T
	for _, ds := range c.Datasources {
		ds := ds
		g.Go(func() error {
			return c.crawlAndProcess(ctx, fv, ds)
		})
	}
	return g.Wait()
}

func (c *T) Index(ctx context.Context, values interface{}, args []string) error {
	return c.RunCommands(ctx, args[0], values, c.IndexCommands)
}

func (c *T) IndexAll(ctx context.Context, values interface{}, _ []string) error {
	var g errgroup.T
	for _, ds := range c.Datasources {
		ds := ds
		g.Go(func() error {
			return c.RunCommands(ctx, ds, values, c.IndexCommands)
		})
	}
	return g.Wait()
}

func (c *T) CrawlIndex(ctx context.Context, values interface{}, args []string) error {
	return c.RunCommands(ctx, args[0], values, c.CrawlCommands, c.IndexCommands)
}

func (c *T) CrawlIndexAll(ctx context.Context, values interface{}, _ []string) error {
	fv := values.(*CrawlIndexFlags)
	var g errgroup.T
	for _, ds := range c.Datasources {
		ds := ds
		g.Go(func() error {
			if err := c.crawlAndProcess(ctx, &fv.CrawlFlags, ds); err != nil {
				return err
			}
			return c.RunCommands(ctx, ds, &fv.IndexFlags, c.IndexCommands)
		})
	}
	return g.Wait()
}

func (c *T) TestCache(ctx context.Context, values interface{}, args []string) error {
	return c.RunCommands(ctx, args[0], values, c.TestCacheCommands)
}

func (c *T) TestCacheAll(ctx context.Context, values interface{}, _ []string) error {
	for _, ds := range c.Datasources {
		if err := c.RunCommands(ctx, ds, values, c.TestCacheCommands); err != nil {
			return err
		}
	}
	return nil
}

func (c *T) formatCommands(out io.Writer, cmds map[string][]string) error {
	for arg, cmd := range cmds {
		if len(cmd) == 0 {
			continue
		}
		cl, err := c.NewRunner("", struct{}{}).ExpandCommandLine(cmd...)
		if err != nil {
			return err
		}
		fmt.Fprintf(out, " %v: %v\n", arg, cl)
	}
	return nil
}

func (c *T) IndexingStats(ctx context.Context, values interface{}, args []string) error {
	datasources := args
	if len(datasources) == 0 {
		datasources = c.Datasources
	}
	for _, ds := range datasources {
		if err := c.RunCommands(ctx, ds, values, c.IndexStatsCommands); err != nil {
			return err
		}
	}
	return nil
}

func (c *T) ShowAll(_ context.Context, _ interface{}, _ []string) error {
	fmt.Printf("Crawl commands\n")
	var errs errors.M
	err := c.formatCommands(os.Stdout, c.CrawlCommands)
	errs.Append(err)
	fmt.Printf("Crawl Process commands (second step in Crawl command)\n")
	err = c.formatCommands(os.Stdout, c.ProcessCommands)
	errs.Append(err)
	fmt.Printf("Crawl Index commands\n")
	err = c.formatCommands(os.Stdout, c.IndexCommands)
	errs.Append(err)
	return errs.Err()
}

func NewCmdSet(spec string, cmds []CommandSpec, globalFlagValues any) *subcmd.CommandSetYAML {
	cmdSet := subcmd.MustFromYAMLTemplate(spec)
	for _, c := range cmds {
		cmdSet.Set(c.Name).MustRunner(c.Runner, c.FlagValues)
	}
	if globalFlagValues != nil {
		globals := subcmd.NewFlagSet()
		globals.MustRegisterFlagStruct(globalFlagValues, nil, nil)
		cmdSet.WithGlobalFlags(globals)
	}
	return cmdSet
}

// TemplateFuncs returns a template.FuncMap that contains an AuthFile
// function that provides access to the supplied authFiles map.
func TemplateFuncs(authFiles map[string]string) template.FuncMap {
	return template.FuncMap{
		"AuthFile": func(a string) string {
			return authFiles[a]
		},
	}
}

// TemplateVars represents the variables that can be accessed
// from templates.
type TemplateVars struct {
	DatasourceName       string
	DatasourceConfigFile string
	Flags                any
}

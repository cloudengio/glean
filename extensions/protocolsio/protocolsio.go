// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package protocolsio

import (
	"context"
	"encoding/json"
	"fmt"
	"path"

	"cloudeng.io/cmdutil/subcmd"
	"cloudeng.io/file/content"
	gleancfg "cloudeng.io/glean/config"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/converters"
	"cloudeng.io/glean/gleansdk"
	"cloudeng.io/webapi/operations"
	"cloudeng.io/webapi/operations/apicrawlcmd"
	"cloudeng.io/webapi/apis/protocolsio"
	"cloudeng.io/webapi/apis/protocolsio/protocolsiocmd"
	"cloudeng.io/webapi/apis/protocolsio/protocolsiosdk"
)

type CommonFlags struct {
	config.FileFlags
	AuthFile string `subcmd:"protocolsio-auth,$HOME/.protocolsio.yaml,'protocols.io auth config file'"`
}

type ScanFlags struct {
	protocolsiocmd.ScanFlags
	CommonFlags
}

type CrawlFlags struct {
	protocolsiocmd.CrawlFlags
	CommonFlags
}

type GetFlags struct {
	protocolsiocmd.GetFlags
	config.DatasourceName
	CommonFlags
}

const (
	cmdName = "protocols.io"
	cmdSpec = `
- name: protocols.io
  summary: protocols.io commands
  commands:
    - name: get
      summary: get a specific protocol
      arguments:
        - id
        - ...
    - name: scan-downloaded
      summary: scan downloaded protocols
      arguments:
        - datasource-name - the datasource whose downloaded protocols are to be scanned
    - name: crawl
      summary: crawl protocols, i.e. download all available protocols
      arguments:
        - datasource-name - the datasource to be crawled
`
)

var ExtensionSpec = gleancfg.ExtensionSpec{
	Name:       cmdName,
	CmdSpec:    cmdSpec,
	AuthCfg:    gleancfg.APIKey{},
	ServiceCfg: protocolsiocmd.Service{},
	AddFunc:    AddExtension,
}

func AddExtension(extension gleancfg.Extension, cmdSet *subcmd.CommandSetYAML, parents []string) error {
	c := &command{parent: extension}
	cmdSet.Set(append(parents, cmdName, "scan-downloaded")...).MustRunner(
		c.scanDownloadsCmd, &ScanFlags{})
	cmdSet.Set(append(parents, cmdName, "crawl")...).MustRunner(
		c.crawlCmd, &CrawlFlags{})
	cmdSet.Set(append(parents, cmdName, "get")...).MustRunner(
		c.getCmd, &GetFlags{})
	return nil
}

type command struct {
	parent gleancfg.Extension
}

func (cmd *command) new(ctx context.Context, fv CommonFlags, datasource string) (*protocolsiocmd.Command, error) {
	cfg, err := config.DatasourceForName(ctx, fv.ConfigFile, datasource)
	if err != nil {
		return nil, err
	}
	var key gleancfg.APIKey
	token, err := key.ParseAndRead(ctx, fv.AuthFile, cmd.parent.Options().TokenReaders)
	if err != nil {
		return nil, err
	}
	resources := apicrawlcmd.Resources{
		Token:           token,
		NewOperationsFS: cmd.parent.Options().NewOperationsFS,
		NewCheckpointOp: cmd.parent.Options().NewCheckpointOp,
	}
	return protocolsiocmd.NewCommand(ctx, cfg.APICrawls[datasource], resources)
}

func (cmd *command) crawlCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*CrawlFlags)
	c, err := cmd.new(ctx, fv.CommonFlags, args[0])
	if err != nil {
		return err
	}
	return c.Crawl(ctx, &fv.CrawlFlags)
}

func (cmd *command) getCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*GetFlags)
	c, err := cmd.new(ctx, fv.CommonFlags, fv.Datasource)
	if err != nil {
		return err
	}
	return c.Get(ctx, &fv.GetFlags, args)
}

func (cmd *command) scanDownloadsCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*ScanFlags)
	c, err := cmd.new(ctx, fv.CommonFlags, args[0])
	if err != nil {
		return err
	}
	return c.ScanDownloaded(ctx, &fv.ScanFlags)
}

type converter struct {
}

func (c *converter) Type() content.Type {
	return protocolsio.ContentType
}

func (c *converter) Convert(_ context.Context, _ string, cfg config.Conversion, ctype content.Type, data []byte) (gleansdk.DocumentDefinition, error) {
	var gd gleansdk.DocumentDefinition
	if ctype != cfg.Type {
		return gd, fmt.Errorf("htmlConverter: expected %v, not %v", cfg.Type, ctype)
	}

	var obj content.Object[protocolsiosdk.ProtocolPayload, *operations.Response]
	if err := obj.Decode(data); err != nil {
		return gd, fmt.Errorf("protocols.io converter: failed to decode object data: %v", err)
	}
	if obj.Value.StatusCode != 0 {
		return gd, nil
	}

	p := obj.Value.Protocol
	gd.SetDatasource(cfg.Datasource.Name)
	gd.SetObjectType(path.Base(string(protocolsio.ContentType)))
	gd.SetId(p.URI)
	gd.SetViewURL(p.URL)
	gd.SetTitle(p.Title)
	gd.Summary = &gleansdk.ContentDefinition{}
	gd.Summary.SetMimeType("text/plain")
	var tmpDesc struct {
		Blocks []struct {
			Text string
			Key  string
		}
	}

	if err := json.Unmarshal([]byte(p.Description), &tmpDesc); err == nil && len(tmpDesc.Blocks) > 0 {
		gd.Summary.SetTextContent(tmpDesc.Blocks[0].Text)
	} else {
		gd.Summary.SetTextContent(p.Description)
	}

	gd.Author = &gleansdk.UserReferenceDefinition{}
	gd.Author.SetName(p.Creator.Name)
	gd.Author.SetEmail(p.Creator.Username + "@protocols.io")
	gd.Author.SetDatasourceUserId(p.Creator.Username)
	gd.Permissions = &gleansdk.DocumentPermissionsDefinition{}
	gd.Permissions.SetAllowAnonymousAccess(true)
	gd.CreatedAt = new(int64)
	*gd.CreatedAt = int64(p.CreatedOn)
	return gd, nil
}

func NewDocumentConverter() converters.Document {
	return &converter{}
}

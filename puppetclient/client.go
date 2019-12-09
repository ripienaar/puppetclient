// generated code; DO NOT EDIT

package puppetclient

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"context"

	"github.com/choria-io/go-choria/choria"
	"github.com/choria-io/go-config"
	"github.com/choria-io/go-protocol/protocol"
	"github.com/choria-io/go-srvcache"
	rpcclient "github.com/choria-io/mcorpc-agent-provider/mcorpc/client"
	"github.com/choria-io/mcorpc-agent-provider/mcorpc/ddl/agent"
	"github.com/sirupsen/logrus"
)

// Stats are the statistics for a request
type Stats interface {
	Agent() string
	Action() string
	All() bool
	NoResponseFrom() []string
	UnexpectedResponseFrom() []string
	DiscoveredCount() int
	DiscoveredNodes() *[]string
	FailCount() int
	OKCount() int
	ResponsesCount() int
	PublishDuration() (time.Duration, error)
	RequestDuration() (time.Duration, error)
	DiscoveryDuration() (time.Duration, error)
}

// NodeSource discovers nodes
type NodeSource interface {
	Reset()
	Discover(ctx context.Context, fw ChoriaFramework, filters []FilterFunc) ([]string, error)
}

// ChoriaFramework is the choria framework
type ChoriaFramework interface {
	Logger(string) *logrus.Entry
	Configuration() *config.Config
	NewMessage(payload string, agent string, collective string, msgType string, request *choria.Message) (msg *choria.Message, err error)
	NewReplyFromTransportJSON(payload []byte, skipvalidate bool) (msg protocol.Reply, err error)
	NewTransportFromJSON(data string) (message protocol.TransportMessage, err error)
	MiddlewareServers() (servers srvcache.Servers, err error)
	NewConnector(ctx context.Context, servers func() (srvcache.Servers, error), name string, logger *logrus.Entry) (conn choria.Connector, err error)
	NewRequestID() (string, error)
	Certname() string
}

// FilterFunc can generate a choria filter
type FilterFunc func(f *protocol.Filter) error

// PuppetClient to the puppet agent
type PuppetClient struct {
	fw            ChoriaFramework
	cfg           *config.Config
	ddl           *agent.DDL
	ns            NodeSource
	clientOpts    *initOptions
	clientRPCOpts []rpcclient.RequestOption
	filters       []FilterFunc
}

// Metadata is the agent metadata
type Metadata struct {
	License     string `json:"license"`
	Author      string `json:"author"`
	Timeout     int    `json:"timeout"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

// Must create a new client and panics on error
func Must(opts ...InitializationOption) (client *PuppetClient) {
	c, err := New(opts...)
	if err != nil {
		panic(err)
	}

	return c
}

// New creates a new client to the puppet agent
func New(opts ...InitializationOption) (client *PuppetClient, err error) {
	c := &PuppetClient{
		ddl:           &agent.DDL{},
		clientRPCOpts: []rpcclient.RequestOption{},
		filters:       []FilterFunc{},
		clientOpts: &initOptions{
			cfgFile: choria.UserConfig(),
		},
	}

	for _, opt := range opts {
		opt(c.clientOpts)
	}

	if c.clientOpts.ns == nil {
		c.clientOpts.ns = &BroadcastNS{}
	}
	c.ns = c.clientOpts.ns

	c.fw, err = choria.New(c.clientOpts.cfgFile)
	if err != nil {
		return nil, fmt.Errorf("could not initialize choria: %s", err)
	}

	c.cfg = c.fw.Configuration()

	if c.clientOpts.logger == nil {
		c.clientOpts.logger = c.fw.Logger("puppet")
	}

	ddlj, err := base64.StdEncoding.DecodeString(rawDDL)
	if err != nil {
		return nil, fmt.Errorf("could not parse embedded DDL: %s", err)
	}

	err = json.Unmarshal(ddlj, c.ddl)
	if err != nil {
		return nil, fmt.Errorf("could not parse embedded DDL: %s", err)
	}

	return c, nil
}

// AgentMetadata is the agent metadata this client supports
func (p *PuppetClient) AgentMetadata() *Metadata {
	return &Metadata{
		License:     p.ddl.Metadata.License,
		Author:      p.ddl.Metadata.Author,
		Timeout:     p.ddl.Metadata.Timeout,
		Name:        p.ddl.Metadata.Name,
		Version:     p.ddl.Metadata.Version,
		URL:         p.ddl.Metadata.URL,
		Description: p.ddl.Metadata.Description,
	}
}

// Disable performs the disable action
//
// Description: Disable the Puppet agent
//
// Optional Inputs:
//    - message (string) - Supply a reason for disabling the Puppet agent
func (p *PuppetClient) Disable() *DisableRequestor {
	d := &DisableRequestor{
		outc: nil,
		r: &requestor{
			args:   map[string]interface{}{},
			action: "disable",
			client: p,
		},
	}

	return d
}

// Enable performs the enable action
//
// Description: Enable the Puppet agent
func (p *PuppetClient) Enable() *EnableRequestor {
	d := &EnableRequestor{
		outc: nil,
		r: &requestor{
			args:   map[string]interface{}{},
			action: "enable",
			client: p,
		},
	}

	return d
}

// LastRunSummary performs the last_run_summary action
//
// Description: Get the summary of the last Puppet run
//
// Optional Inputs:
//    - logs (bool) - Whether or not to parse the logs from last_run_report.yaml
func (p *PuppetClient) LastRunSummary() *LastRunSummaryRequestor {
	d := &LastRunSummaryRequestor{
		outc: nil,
		r: &requestor{
			args:   map[string]interface{}{},
			action: "last_run_summary",
			client: p,
		},
	}

	return d
}

// Resource performs the resource action
//
// Description: Evaluate Puppet RAL resources
//
// Required Inputs:
//    - name (string) - Resource Name
//    - type (string) - Resource Type
//
// Optional Inputs:
//    - environment (string) - Which Puppet environment to use
func (p *PuppetClient) Resource(nameInput string, typeInput string) *ResourceRequestor {
	d := &ResourceRequestor{
		outc: nil,
		r: &requestor{
			args: map[string]interface{}{
				"name": nameInput,
				"type": typeInput,
			},
			action: "resource",
			client: p,
		},
	}

	return d
}

// Runonce performs the runonce action
//
// Description: Invoke a single Puppet run
//
// Optional Inputs:
//    - environment (string) - Which Puppet environment to run
//    - force (bool) - Will force a run immediately else subject to default splay time
//    - noop (bool) - Do a Puppet dry run
//    - server (string) - Address and port of the Puppet Master in server:port format
//    - splay (bool) - Sleep for a period before initiating the run
//    - splaylimit (float64) - Maximum amount of time to sleep before run
//    - tags (string) - Restrict the Puppet run to a comma list of tags
//    - use_cached_catalog (bool) - Determine if to use the cached catalog or not
func (p *PuppetClient) Runonce() *RunonceRequestor {
	d := &RunonceRequestor{
		outc: nil,
		r: &requestor{
			args:   map[string]interface{}{},
			action: "runonce",
			client: p,
		},
	}

	return d
}

// Status performs the status action
//
// Description: Get the current status of the Puppet agent
func (p *PuppetClient) Status() *StatusRequestor {
	d := &StatusRequestor{
		outc: nil,
		r: &requestor{
			args:   map[string]interface{}{},
			action: "status",
			client: p,
		},
	}

	return d
}

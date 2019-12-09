// generated code; DO NOT EDIT

package puppetclient

import (
	"time"

	coreclient "github.com/choria-io/go-client/client"
	rpcclient "github.com/choria-io/mcorpc-agent-provider/mcorpc/client"
)

// OptionReset resets the client options to use across requests to an empty list
func (p *PuppetClient) OptionReset() *PuppetClient {
	p.clientRPCOpts = []rpcclient.RequestOption{}
	p.ns = p.clientOpts.ns
	p.filters = []FilterFunc{}

	return p
}

// OptionFactFilter adds a fact filter
func (p *PuppetClient) OptionFactFilter(f ...string) *PuppetClient {
	for _, i := range f {
		p.filters = append(p.filters, FilterFunc(coreclient.FactFilter(i)))
	}

	p.ns.Reset()

	return p
}

// OptionCollective sets the collective to target
func (p *PuppetClient) OptionCollective(c string) *PuppetClient {
	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.Collective(c))
	return p
}

// OptionInBatches performs requests in batches
func (p *PuppetClient) OptionInBatches(size int, sleep int) *PuppetClient {
	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.InBatches(size, sleep))
	return p
}

// OptionDiscoveryTimeout configures the request discovery timeout, defaults to configured discovery timeout
func (p *PuppetClient) OptionDiscoveryTimeout(t time.Duration) *PuppetClient {
	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.DiscoveryTimeout(t))
	return p
}

// OptionLimitMethod configures the method to use when limiting targets - "random" or "first"
func (p *PuppetClient) OptionLimitMethod(m string) *PuppetClient {
	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.LimitMethod(m))
	return p
}

// OptionLimitSize sets limits on the targets, either a number of a percentage like "10%"
func (p *PuppetClient) OptionLimitSize(s string) *PuppetClient {
	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.LimitSize(s))
	return p
}

// OptionLimitSeed sets the random seed used to select targets when limiting and limit method is "random"
func (p *PuppetClient) OptionLimitSeed(s int64) *PuppetClient {
	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.LimitSeed(s))
	return p
}

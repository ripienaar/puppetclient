// generated code; DO NOT EDIT

package puppetclient

import (
	"context"
	"encoding/json"

	"github.com/choria-io/go-protocol/protocol"
	rpcclient "github.com/choria-io/mcorpc-agent-provider/mcorpc/client"
)

// RunonceRequestor performs a RPC request to puppet#runonce
type RunonceRequestor struct {
	r    *requestor
	outc chan *RunonceOutput
}

// RunonceOutput is the output from the runonce action
type RunonceOutput struct {
	details *ResultDetails
	reply   map[string]interface{}
}

// RunonceResult is the result from a runonce action
type RunonceResult struct {
	stats   *rpcclient.Stats
	outputs []*RunonceOutput
}

// Stats is the rpc request stats
func (d *RunonceResult) Stats() Stats {
	return d.stats
}

// ResultDetails is the details about the request
func (d *RunonceOutput) ResultDetails() *ResultDetails {
	return d.details
}

// HashMap is the raw output data
func (d *RunonceOutput) HashMap() map[string]interface{} {
	return d.reply
}

// JSON is the JSON representation of the output data
func (d *RunonceOutput) JSON() ([]byte, error) {
	return json.Marshal(d.reply)
}

// Do performs the request
func (d *RunonceRequestor) Do(ctx context.Context) (*RunonceResult, error) {
	dres := &RunonceResult{}

	handler := func(pr protocol.Reply, r *rpcclient.RPCReply) {
		output := &RunonceOutput{
			reply: make(map[string]interface{}),
			details: &ResultDetails{
				sender:  pr.SenderID(),
				code:    int(r.Statuscode),
				message: r.Statusmsg,
				ts:      pr.Time(),
			},
		}

		err := json.Unmarshal(r.Data, &output.reply)
		if err != nil {
			d.r.client.errorf("Could not decode reply from %s: %s", pr.SenderID(), err)
		}

		if d.outc != nil {
			d.outc <- output
			return
		}

		dres.outputs = append(dres.outputs, output)
	}

	res, err := d.r.do(ctx, handler)
	if err != nil {
		return nil, err
	}

	dres.stats = res

	return dres, nil
}

// EachOutput iterates over all results received
func (d *RunonceResult) EachOutput(h func(r *RunonceOutput)) {
	for _, resp := range d.outputs {
		h(resp)
	}
}

// Environment is an optional input to the runonce action
//
// Description: Which Puppet environment to run
func (d *RunonceRequestor) Environment(v string) *RunonceRequestor {
	d.r.args["environment"] = v

	return d
}

// Force is an optional input to the runonce action
//
// Description: Will force a run immediately else subject to default splay time
func (d *RunonceRequestor) Force(v bool) *RunonceRequestor {
	d.r.args["force"] = v

	return d
}

// Noop is an optional input to the runonce action
//
// Description: Do a Puppet dry run
func (d *RunonceRequestor) Noop(v bool) *RunonceRequestor {
	d.r.args["noop"] = v

	return d
}

// Server is an optional input to the runonce action
//
// Description: Address and port of the Puppet Master in server:port format
func (d *RunonceRequestor) Server(v string) *RunonceRequestor {
	d.r.args["server"] = v

	return d
}

// Splay is an optional input to the runonce action
//
// Description: Sleep for a period before initiating the run
func (d *RunonceRequestor) Splay(v bool) *RunonceRequestor {
	d.r.args["splay"] = v

	return d
}

// Splaylimit is an optional input to the runonce action
//
// Description: Maximum amount of time to sleep before run
func (d *RunonceRequestor) Splaylimit(v float64) *RunonceRequestor {
	d.r.args["splaylimit"] = v

	return d
}

// Tags is an optional input to the runonce action
//
// Description: Restrict the Puppet run to a comma list of tags
func (d *RunonceRequestor) Tags(v string) *RunonceRequestor {
	d.r.args["tags"] = v

	return d
}

// UseCachedCatalog is an optional input to the runonce action
//
// Description: Determine if to use the cached catalog or not
func (d *RunonceRequestor) UseCachedCatalog(v bool) *RunonceRequestor {
	d.r.args["use_cached_catalog"] = v

	return d
}

// InitiatedAt is the value of the initiated_at output
//
// Description: Timestamp of when the runonce command was issues
func (d *RunonceOutput) InitiatedAt() int64 {
	val := d.reply["initiated_at"]
	return val.(int64)
}

// Summary is the value of the summary output
//
// Description: Summary of command run
func (d *RunonceOutput) Summary() string {
	val := d.reply["summary"]
	return val.(string)
}

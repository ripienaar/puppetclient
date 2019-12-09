// generated code; DO NOT EDIT

package puppetclient

import (
	"context"
	"encoding/json"

	"github.com/choria-io/go-protocol/protocol"
	rpcclient "github.com/choria-io/mcorpc-agent-provider/mcorpc/client"
)

// DisableRequestor performs a RPC request to puppet#disable
type DisableRequestor struct {
	r    *requestor
	outc chan *DisableOutput
}

// DisableOutput is the output from the disable action
type DisableOutput struct {
	details *ResultDetails
	reply   map[string]interface{}
}

// DisableResult is the result from a disable action
type DisableResult struct {
	stats   *rpcclient.Stats
	outputs []*DisableOutput
}

// Stats is the rpc request stats
func (d *DisableResult) Stats() Stats {
	return d.stats
}

// ResultDetails is the details about the request
func (d *DisableOutput) ResultDetails() *ResultDetails {
	return d.details
}

// HashMap is the raw output data
func (d *DisableOutput) HashMap() map[string]interface{} {
	return d.reply
}

// JSON is the JSON representation of the output data
func (d *DisableOutput) JSON() ([]byte, error) {
	return json.Marshal(d.reply)
}

// Do performs the request
func (d *DisableRequestor) Do(ctx context.Context) (*DisableResult, error) {
	dres := &DisableResult{}

	handler := func(pr protocol.Reply, r *rpcclient.RPCReply) {
		output := &DisableOutput{
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
func (d *DisableResult) EachOutput(h func(r *DisableOutput)) {
	for _, resp := range d.outputs {
		h(resp)
	}
}

// Message is an optional input to the disable action
//
// Description: Supply a reason for disabling the Puppet agent
func (d *DisableRequestor) Message(v string) *DisableRequestor {
	d.r.args["message"] = v

	return d
}

// Enabled is the value of the enabled output
//
// Description: Is the agent currently locked
func (d *DisableOutput) Enabled() bool {
	val := d.reply["enabled"]
	return val.(bool)
}

// Status is the value of the status output
//
// Description: Status
func (d *DisableOutput) Status() string {
	val := d.reply["status"]
	return val.(string)
}

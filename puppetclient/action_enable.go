// generated code; DO NOT EDIT

package puppetclient

import (
	"context"
	"encoding/json"

	"github.com/choria-io/go-protocol/protocol"
	rpcclient "github.com/choria-io/mcorpc-agent-provider/mcorpc/client"
)

// EnableRequestor performs a RPC request to puppet#enable
type EnableRequestor struct {
	r    *requestor
	outc chan *EnableOutput
}

// EnableOutput is the output from the enable action
type EnableOutput struct {
	details *ResultDetails
	reply   map[string]interface{}
}

// EnableResult is the result from a enable action
type EnableResult struct {
	stats   *rpcclient.Stats
	outputs []*EnableOutput
}

// Stats is the rpc request stats
func (d *EnableResult) Stats() Stats {
	return d.stats
}

// ResultDetails is the details about the request
func (d *EnableOutput) ResultDetails() *ResultDetails {
	return d.details
}

// HashMap is the raw output data
func (d *EnableOutput) HashMap() map[string]interface{} {
	return d.reply
}

// JSON is the JSON representation of the output data
func (d *EnableOutput) JSON() ([]byte, error) {
	return json.Marshal(d.reply)
}

// Do performs the request
func (d *EnableRequestor) Do(ctx context.Context) (*EnableResult, error) {
	dres := &EnableResult{}

	handler := func(pr protocol.Reply, r *rpcclient.RPCReply) {
		output := &EnableOutput{
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
func (d *EnableResult) EachOutput(h func(r *EnableOutput)) {
	for _, resp := range d.outputs {
		h(resp)
	}
}

// Enabled is the value of the enabled output
//
// Description: Is the agent currently locked
func (d *EnableOutput) Enabled() bool {
	val := d.reply["enabled"]
	return val.(bool)
}

// Status is the value of the status output
//
// Description: Status
func (d *EnableOutput) Status() string {
	val := d.reply["status"]
	return val.(string)
}

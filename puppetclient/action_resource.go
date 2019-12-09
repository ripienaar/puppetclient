// generated code; DO NOT EDIT

package puppetclient

import (
	"context"
	"encoding/json"

	"github.com/choria-io/go-protocol/protocol"
	rpcclient "github.com/choria-io/mcorpc-agent-provider/mcorpc/client"
)

// ResourceRequestor performs a RPC request to puppet#resource
type ResourceRequestor struct {
	r    *requestor
	outc chan *ResourceOutput
}

// ResourceOutput is the output from the resource action
type ResourceOutput struct {
	details *ResultDetails
	reply   map[string]interface{}
}

// ResourceResult is the result from a resource action
type ResourceResult struct {
	stats   *rpcclient.Stats
	outputs []*ResourceOutput
}

// Stats is the rpc request stats
func (d *ResourceResult) Stats() Stats {
	return d.stats
}

// ResultDetails is the details about the request
func (d *ResourceOutput) ResultDetails() *ResultDetails {
	return d.details
}

// HashMap is the raw output data
func (d *ResourceOutput) HashMap() map[string]interface{} {
	return d.reply
}

// JSON is the JSON representation of the output data
func (d *ResourceOutput) JSON() ([]byte, error) {
	return json.Marshal(d.reply)
}

// Do performs the request
func (d *ResourceRequestor) Do(ctx context.Context) (*ResourceResult, error) {
	dres := &ResourceResult{}

	handler := func(pr protocol.Reply, r *rpcclient.RPCReply) {
		output := &ResourceOutput{
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
func (d *ResourceResult) EachOutput(h func(r *ResourceOutput)) {
	for _, resp := range d.outputs {
		h(resp)
	}
}

// Environment is an optional input to the resource action
//
// Description: Which Puppet environment to use
func (d *ResourceRequestor) Environment(v string) *ResourceRequestor {
	d.r.args["environment"] = v

	return d
}

// Changed is the value of the changed output
//
// Description: Was a change applied based on the resource
func (d *ResourceOutput) Changed() bool {
	val := d.reply["changed"]
	return val.(bool)
}

// Result is the value of the result output
//
// Description: The result from the Puppet resource
func (d *ResourceOutput) Result() string {
	val := d.reply["result"]
	return val.(string)
}

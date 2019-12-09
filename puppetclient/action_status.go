// generated code; DO NOT EDIT

package puppetclient

import (
	"context"
	"encoding/json"

	"github.com/choria-io/go-protocol/protocol"
	rpcclient "github.com/choria-io/mcorpc-agent-provider/mcorpc/client"
)

// StatusRequestor performs a RPC request to puppet#status
type StatusRequestor struct {
	r    *requestor
	outc chan *StatusOutput
}

// StatusOutput is the output from the status action
type StatusOutput struct {
	details *ResultDetails
	reply   map[string]interface{}
}

// StatusResult is the result from a status action
type StatusResult struct {
	stats   *rpcclient.Stats
	outputs []*StatusOutput
}

// Stats is the rpc request stats
func (d *StatusResult) Stats() Stats {
	return d.stats
}

// ResultDetails is the details about the request
func (d *StatusOutput) ResultDetails() *ResultDetails {
	return d.details
}

// HashMap is the raw output data
func (d *StatusOutput) HashMap() map[string]interface{} {
	return d.reply
}

// JSON is the JSON representation of the output data
func (d *StatusOutput) JSON() ([]byte, error) {
	return json.Marshal(d.reply)
}

// Do performs the request
func (d *StatusRequestor) Do(ctx context.Context) (*StatusResult, error) {
	dres := &StatusResult{}

	handler := func(pr protocol.Reply, r *rpcclient.RPCReply) {
		output := &StatusOutput{
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
func (d *StatusResult) EachOutput(h func(r *StatusOutput)) {
	for _, resp := range d.outputs {
		h(resp)
	}
}

// Applying is the value of the applying output
//
// Description: Is a catalog being applied
func (d *StatusOutput) Applying() string {
	val := d.reply["applying"]
	return val.(string)
}

// DaemonPresent is the value of the daemon_present output
//
// Description: Is the Puppet agent daemon running on this system
func (d *StatusOutput) DaemonPresent() bool {
	val := d.reply["daemon_present"]
	return val.(bool)
}

// DisableMessage is the value of the disable_message output
//
// Description: Message supplied when agent was disabled
func (d *StatusOutput) DisableMessage() string {
	val := d.reply["disable_message"]
	return val.(string)
}

// Enabled is the value of the enabled output
//
// Description: Is the agent currently locked
func (d *StatusOutput) Enabled() bool {
	val := d.reply["enabled"]
	return val.(bool)
}

// Idling is the value of the idling output
//
// Description: Is the Puppet agent daemon running but not doing any work
func (d *StatusOutput) Idling() bool {
	val := d.reply["idling"]
	return val.(bool)
}

// Lastrun is the value of the lastrun output
//
// Description: When the Agent last applied a catalog in local time
func (d *StatusOutput) Lastrun() int64 {
	val := d.reply["lastrun"]
	return val.(int64)
}

// Message is the value of the message output
//
// Description: Descriptive message defining the overall agent status
func (d *StatusOutput) Message() string {
	val := d.reply["message"]
	return val.(string)
}

// SinceLastrun is the value of the since_lastrun output
//
// Description: How long ago did the Agent last apply a catalog in local time
func (d *StatusOutput) SinceLastrun() float64 {
	val := d.reply["since_lastrun"]
	return val.(float64)
}

// Status is the value of the status output
//
// Description: Current status of the Puppet agent
func (d *StatusOutput) Status() string {
	val := d.reply["status"]
	return val.(string)
}

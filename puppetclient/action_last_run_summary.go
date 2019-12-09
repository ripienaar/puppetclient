// generated code; DO NOT EDIT

package puppetclient

import (
	"context"
	"encoding/json"

	"github.com/choria-io/go-protocol/protocol"
	rpcclient "github.com/choria-io/mcorpc-agent-provider/mcorpc/client"
)

// LastRunSummaryRequestor performs a RPC request to puppet#last_run_summary
type LastRunSummaryRequestor struct {
	r    *requestor
	outc chan *LastRunSummaryOutput
}

// LastRunSummaryOutput is the output from the last_run_summary action
type LastRunSummaryOutput struct {
	details *ResultDetails
	reply   map[string]interface{}
}

// LastRunSummaryResult is the result from a last_run_summary action
type LastRunSummaryResult struct {
	stats   *rpcclient.Stats
	outputs []*LastRunSummaryOutput
}

// Stats is the rpc request stats
func (d *LastRunSummaryResult) Stats() Stats {
	return d.stats
}

// ResultDetails is the details about the request
func (d *LastRunSummaryOutput) ResultDetails() *ResultDetails {
	return d.details
}

// HashMap is the raw output data
func (d *LastRunSummaryOutput) HashMap() map[string]interface{} {
	return d.reply
}

// JSON is the JSON representation of the output data
func (d *LastRunSummaryOutput) JSON() ([]byte, error) {
	return json.Marshal(d.reply)
}

// Do performs the request
func (d *LastRunSummaryRequestor) Do(ctx context.Context) (*LastRunSummaryResult, error) {
	dres := &LastRunSummaryResult{}

	handler := func(pr protocol.Reply, r *rpcclient.RPCReply) {
		output := &LastRunSummaryOutput{
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
func (d *LastRunSummaryResult) EachOutput(h func(r *LastRunSummaryOutput)) {
	for _, resp := range d.outputs {
		h(resp)
	}
}

// Logs is an optional input to the last_run_summary action
//
// Description: Whether or not to parse the logs from last_run_report.yaml
func (d *LastRunSummaryRequestor) Logs(v bool) *LastRunSummaryRequestor {
	d.r.args["logs"] = v

	return d
}

// ChangedResources is the value of the changed_resources output
//
// Description: Resources that were changed
func (d *LastRunSummaryOutput) ChangedResources() int64 {
	val := d.reply["changed_resources"]
	return val.(int64)
}

// ConfigRetrievalTime is the value of the config_retrieval_time output
//
// Description: Time taken to retrieve the catalog from the master
func (d *LastRunSummaryOutput) ConfigRetrievalTime() float64 {
	val := d.reply["config_retrieval_time"]
	return val.(float64)
}

// ConfigVersion is the value of the config_version output
//
// Description: Puppet config version for the previously applied catalog
func (d *LastRunSummaryOutput) ConfigVersion() string {
	val := d.reply["config_version"]
	return val.(string)
}

// CorrectedResources is the value of the corrected_resources output
//
// Description: Resources that were correctively changed
func (d *LastRunSummaryOutput) CorrectedResources() int64 {
	val := d.reply["corrected_resources"]
	return val.(int64)
}

// FailedResources is the value of the failed_resources output
//
// Description: Resources that failed to apply
func (d *LastRunSummaryOutput) FailedResources() int64 {
	val := d.reply["failed_resources"]
	return val.(int64)
}

// Lastrun is the value of the lastrun output
//
// Description: When the Agent last applied a catalog in local time
func (d *LastRunSummaryOutput) Lastrun() int64 {
	val := d.reply["lastrun"]
	return val.(int64)
}

// Logs is the value of the logs output
//
// Description: Log lines from the last Puppet run
func (d *LastRunSummaryOutput) Logs() []interface{} {
	val := d.reply["logs"]
	return val.([]interface{})
}

// OutOfSyncResources is the value of the out_of_sync_resources output
//
// Description: Resources that were not in desired state
func (d *LastRunSummaryOutput) OutOfSyncResources() int64 {
	val := d.reply["out_of_sync_resources"]
	return val.(int64)
}

// SinceLastrun is the value of the since_lastrun output
//
// Description: How long ago did the Agent last apply a catalog in local time
func (d *LastRunSummaryOutput) SinceLastrun() int64 {
	val := d.reply["since_lastrun"]
	return val.(int64)
}

// Summary is the value of the summary output
//
// Description: Summary data as provided by Puppet
func (d *LastRunSummaryOutput) Summary() map[string]interface{} {
	val := d.reply["summary"]
	return val.(map[string]interface{})
}

// TotalResources is the value of the total_resources output
//
// Description: Total resources managed on a node
func (d *LastRunSummaryOutput) TotalResources() int64 {
	val := d.reply["total_resources"]
	return val.(int64)
}

// TotalTime is the value of the total_time output
//
// Description: Total time taken to retrieve and process the catalog
func (d *LastRunSummaryOutput) TotalTime() float64 {
	val := d.reply["total_time"]
	return val.(float64)
}

// TypeDistribution is the value of the type_distribution output
//
// Description: Resource counts per type managed by Puppet
func (d *LastRunSummaryOutput) TypeDistribution() map[string]interface{} {
	val := d.reply["type_distribution"]
	return val.(map[string]interface{})
}

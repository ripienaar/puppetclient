package main

import (
	"context"
	"fmt"

	"github.com/ripienaar/puppet/puppetclient"
)

func summarize(r puppetclient.Stats) {
	nr := r.NoResponseFrom()
	if len(nr) > 0 {
		fmt.Printf("No responses received from %d hosts", len(nr))
	}
}

func enable(ctx context.Context) error {
	// instance of puppet client, panics if fails
	// uses default puppet config path
	pc := puppetclient.Must()

	// call to the network, filtering on all nodes in `mt`,
	// Enable() has no required or optional inputs, see disable
	res, err := pc.OptionFactFilter("country=mt").Enable().Do(ctx)
	if err != nil {
		return err
	}

	// prints all the responses from the network
	res.EachOutput(func(r *puppetclient.EnableOutput) {
		if r.ResultDetails().OK() {
			fmt.Printf("OK: %-40s: enabled: %v\n", r.ResultDetails().Sender(), r.Enabled())
		} else {
			fmt.Printf("!!: %-40s: message: %s\n", r.ResultDetails().Sender(), r.ResultDetails().StatusMessage())
		}
	})

	summarize(res.Stats())

	return nil
}

func main() {
	err := enable(context.Background())
	if err != nil {
		panic(err)
	}
}

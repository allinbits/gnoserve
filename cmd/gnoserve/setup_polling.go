package main

import (
	"context"
	"fmt"
	"github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	"log/slog"
	"strings"
)

const renderPath = "vm/qrender"
const gnoFrameRealm = "gno.land/r/gnoframe"

func tick(cli *client.RPCClient, logger *slog.Logger) {
	data := []byte(gnoFrameRealm + ":frame") // FIXME: this should correspond to schema declarations

	// add a defer to recover from any panics during the tick
	defer func() {
		if r := recover(); r != nil {
			logger.Error("panic during tick", "error", r)
		}
	}()

	logger.Info("Ticking...")

	// read the '/r/gnoframe realm display func
	logger.Info("Polling for events...")

	res, err := cli.ABCIQuery(renderPath, data)
	if err != nil {
		logger.Error("error querying events", "error", err)
		return
	}
	if len(res.Response.Data) == 0 {
		logger.Info("No events found in response")
		return
	}
	logger.Info(fmt.Sprintf("Events found: %s", res.Response.Data))
}

func setupPolling(ctx context.Context, logger *slog.Logger, remoteAddr string) {
	httpClient, err := client.NewHTTPClient(remoteAddr)
	_ = httpClient // TODO: use this client to poll events
	if err != nil {
		logger.Error("unable to create HTTP client", "error", err)
		return
	}
	// Disabled for now while doing gno development
	// ticker := time.NewTicker(5 * time.Second) // match block time of gno chain

	// go func() {
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			tick(httpClient, logger)
	// 		case <-ctx.Done():
	// 			ticker.Stop()
	// 			return
	// 		}
	// 	}
	// }()
}

// TODO: actually remove
func exampleTest() {
	importedCid := cid.NewCid(importedModel).String()
	modelCid := cid.NewCid(exampleModel).String()
	if importedCid != modelCid {
		panic("Expected CIDs to match, got " + importedCid + " != " + modelCid)
	}
}

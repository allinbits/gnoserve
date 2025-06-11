package service

import (
	"fmt"
	"github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	"log/slog"
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


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

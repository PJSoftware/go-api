package api

import "log/slog"

var apiLogger *slog.Logger

func init() {
	apiLogger = slog.With(
		slog.String("api_name", "go-api"),
		slog.String("api_version", Version),
	)
}

func errLog(err error) error {
	apiLogger.Error("error", "err", err.Error())
	return err
}

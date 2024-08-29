package api

import "log/slog"

var apiLogger *slog.Logger

func init() {
	apiLogger = slog.Default().With(
		slog.String("pkg_name", pkgName),
		slog.String("pkg_version", pkgVersion),
	)
	apiLogger.Debug("initialising go-api logging")
}

func errLog(err error) error {
	if err != nil {
		apiLogger.Error("error", "err", err)
	}
	
	return err
}

package api

import "log/slog"

var apiLogger *slog.Logger

func init() {
	apiLogger = slog.With(
		slog.String("pkg_name", pkgName),
		slog.String("pkg_version", pkgVersion),
	)
}

func errLog(err error) error {
	apiLogger.Error("error", "err", err.Error())
	return err
}

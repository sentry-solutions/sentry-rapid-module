package interfaces

import "github.com/sentry-solutions/sentry-rapid-module/common"

type IRapidModule interface {
	Log(logType common.LoggerLogType, message string)
	Ok(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)
}

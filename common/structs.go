package common

type RapidModuleSettings struct {
	SyncMode       LoggerVeritasSyncMode
	LogLevel       LoggerLogLevel
	ContextFilter  LoggerContextFilter
	Origin         string
	RepositoryUrl  string
	ProjectVersion string
}

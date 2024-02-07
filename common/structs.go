package common

type RapidSettings struct {
	SyncMode       LoggerVeritasSyncMode
	LogLevel       LoggerLogLevel
	ContextFilter  LoggerContextFilter
	Origin         string
	RepositoryUrl  string
	ProjectVersion string
}

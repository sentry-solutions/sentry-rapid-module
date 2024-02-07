package common

// Veritas Sync Mode
const (
	VeritasDevelopment = LoggerVeritasSyncMode("VERITAS_SYNC_MODE_DEVELOPMENT")
	VeritasProduction  = LoggerVeritasSyncMode("VERITAS_SYNC_MODE_PRODUCTION")
)

// Log Type
const (
	LogTypeOk    = LoggerLogType("LOG_TYPE_OK")
	LogTypeInfo  = LoggerLogType("LOG_TYPE_INFO")
	LogTypeWarn  = LoggerLogType("LOG_TYPE_WARN")
	LogTypeError = LoggerLogType("LOG_TYPE_ERROR")
	LogTypeFatal = LoggerLogType("LOG_TYPE_FATAL")
)

// Log Level
const (
	LogLevelSilent LoggerLogLevel = iota
	LogLevelFull
	LogLevelOk
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

// Log Context
const (
	LoggerContextFilterEvent   = LoggerContextFilter("LOG_CONTEXT_FILTER_EVENT")
	LoggerContextFilterLib     = LoggerContextFilter("LOG_CONTEXT_FILTER_LIB")
	LoggerContextFilterGeneric = LoggerContextFilter("LOG_CONTEXT_FILTER_GENERIC")
	LoggerContextFilterNone    = LoggerContextFilter("LOG_CONTEXT_FILTER_NONE")
)

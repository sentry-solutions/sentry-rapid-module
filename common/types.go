package common

type LoggerVeritasSyncMode string
type LoggerContextFilter string
type LoggerLogType string

type LoggerLogLevel int

func (level LoggerLogType) String() string {
	switch level {
	case LogTypeOk:
		return "OK"
	case LogTypeInfo:
		return "INFO"
	case LogTypeWarn:
		return "WARN"
	case LogTypeError:
		return "ERROR"
	case LogTypeFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

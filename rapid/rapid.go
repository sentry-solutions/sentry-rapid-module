package rapid

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/sentry-solutions/sentry-rapid-module/common"
	"github.com/sentry-solutions/sentry-rapid-module/interfaces"
)

type RapidModule struct {
	Settings    common.RapidModuleSettings
	LogColors   map[common.LoggerLogType]color.Color
	ExtraColors map[string]color.Color
}

// Ensure the service satisfies the interface
var _ interfaces.IRapidModule = (*RapidModule)(nil)

func NewModule(settings common.RapidModuleSettings) *RapidModule {

	logColors := map[common.LoggerLogType]color.Color{
		common.LogTypeOk:    *color.New(color.FgHiGreen),
		common.LogTypeInfo:  *color.New(color.FgHiBlue),
		common.LogTypeWarn:  *color.New(color.FgHiYellow),
		common.LogTypeError: *color.New(color.FgRed),
		common.LogTypeFatal: *color.New(color.FgHiRed),
	}

	extraColors := map[string]color.Color{
		"timestamp": *color.New(color.FgYellow),
	}

	return &RapidModule{
		Settings:    settings,
		LogColors:   logColors,
		ExtraColors: extraColors,
	}
}

func (module *RapidModule) Log(logType common.LoggerLogType, message string) {

	switch module.Settings.LogLevel {
	case common.LogLevelSilent:
		return
	case common.LogLevelOk:
		if logType != common.LogTypeOk {
			return
		}
	case common.LogLevelInfo:
		if logType != common.LogTypeInfo {
			return
		}
	case common.LogLevelWarn:
		if logType != common.LogTypeWarn {
			return
		}
	case common.LogLevelError:
		if logType != common.LogTypeError {
			return
		}
	case common.LogLevelFatal:
		if logType != common.LogTypeFatal {
			return
		}
	}

	t := time.Now().Format("2006-01-02 15:04:05")
	logTypeColor := module.LogColors[logType]
	timestampColor := module.ExtraColors["timestamp"]
	fmt.Printf("%s %s %s\n", timestampColor.Sprint(t), logTypeColor.Sprint(logType), message)
}

func (module *RapidModule) Ok(message string) {
	module.Log(common.LogTypeOk, message)
}

func (module *RapidModule) Info(message string) {
	module.Log(common.LogTypeInfo, message)
}

func (module *RapidModule) Warn(message string) {
	module.Log(common.LogTypeWarn, message)
}

func (module *RapidModule) Error(message string) {
	module.Log(common.LogTypeError, message)
}

func (module *RapidModule) Fatal(message string) {
	module.Log(common.LogTypeFatal, message)
}

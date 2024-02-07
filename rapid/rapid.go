package rapid

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/sentry-solutions/sentry-rapid-module/common"
)

type Rapid struct {
	Settings    common.RapidSettings
	LogColors   map[common.LoggerLogType]color.Color
	ExtraColors map[string]color.Color
}

func NewModule(settings common.RapidSettings) *Rapid {

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

	return &Rapid{
		Settings:    settings,
		LogColors:   logColors,
		ExtraColors: extraColors,
	}
}

func (rapid *Rapid) Log(logType common.LoggerLogType, message string) {

	switch rapid.Settings.LogLevel {
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
	logTypeColor := rapid.LogColors[logType]
	timestampColor := rapid.ExtraColors["timestamp"]
	fmt.Printf("%s %s %s\n", timestampColor.Sprint(t), logTypeColor.Sprint(logType), message)
}

func (rapid *Rapid) Ok(message string) {
	rapid.Log(common.LogTypeOk, message)
}

func (rapid *Rapid) Info(message string) {
	rapid.Log(common.LogTypeInfo, message)
}

func (rapid *Rapid) Warn(message string) {
	rapid.Log(common.LogTypeWarn, message)
}

func (rapid *Rapid) Error(message string) {
	rapid.Log(common.LogTypeError, message)
}

func (rapid *Rapid) Fatal(message string) {
	rapid.Log(common.LogTypeFatal, message)
}

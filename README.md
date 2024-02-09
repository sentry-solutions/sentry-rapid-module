# RAPID (Sentry Module)
Real-time Activity and Performance Information Dispatch Module (R.A.P.I.D. Module)

## TODO
- [ ] implement veritas sync module
- [ ] implement events
- [ ] advanced logging logics
- [ ] separate sub-modules for veritas, logger etc

## Usage example

```
import (
	"github.com/sentry-solutions/sentry-rapid-module/common"
	"github.com/sentry-solutions/sentry-rapid-module/rapid"
)

func main() {
	logger := rapid.NewModule(common.RapidModuleSettings{
		SyncMode:       common.VeritasDevelopment,
		LogLevel:       common.LogLevelFull,
		ContextFilter:  common.LoggerContextFilterNone,
		Origin:         "RAPID",
		RepositoryUrl:  "github.com/sentry-solutions/sentry-rapid-module",
		ProjectVersion: "1.0.0",
	})

	logger.Ok("ok test")
	logger.Info("info test")
	logger.Warn("warn test")
	logger.Error("error test")
	logger.Fatal("fatal test")
}
```
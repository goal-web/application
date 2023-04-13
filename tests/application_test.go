package tests

import (
	"github.com/goal-web/application"
	"github.com/goal-web/config"
	"github.com/goal-web/contracts"
	exceptions2 "github.com/goal-web/supports/exceptions"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ShouldReportException struct {
}

func (s ShouldReportException) Error() string {
	return "testing"
}

func (s ShouldReportException) GetPrevious() contracts.Exception {
	return nil
}

func TestMakeApplication(t *testing.T) {
	app := application.New()

	app.RegisterServices(
		config.NewService(config.NewDotEnv(config.File("testing.env")), map[string]contracts.ConfigProvider{
			"app": func(env contracts.Env) any {
				return application.Config{
					Name:     env.GetString("app.name"),
					Debug:    env.GetBool("app.debug"),
					Timezone: env.GetString("app.timezone"),
					Env:      env.GetString("app.env"),
					Locale:   env.GetString("app.locale"),
					Key:      env.GetString("app.key"),
				}
			},
		}),
		exceptions2.NewService([]contracts.Exception{
			exceptions2.New(""),
		}),
	)

	app.Start()

	app.Call(func(handler contracts.ExceptionHandler, config contracts.Config) {
		assert.True(t, handler.ShouldReport(ShouldReportException{}))

		assert.True(t, config.GetString("app.name") == "testing111")
	})
}

package tests

import (
	"github.com/goal-web/application"
	"github.com/goal-web/application/exceptions"
	"github.com/goal-web/config"
	"github.com/goal-web/contracts"
	exceptions2 "github.com/goal-web/supports/exceptions"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type ShouldReportException struct {
}

func (s ShouldReportException) Error() string {
	return "testing"
}

func (s ShouldReportException) GetPrevious() *contracts.Exception {
	return nil
}

func TestMakeApplication(t *testing.T) {
	app := application.New()
	hostname, _ := os.Hostname()
	userHome, _ := os.UserHomeDir()

	app.RegisterServices(
		config.NewService("testing", ".", map[string]contracts.ConfigProvider{
			"app": application.ConfigProvider(hostname, userHome),
		}),
		exceptions.NewService([]contracts.Exception{
			exceptions2.New(""),
		}),
	)

	app.Start()

	assert.Equal(t, app.GetConfig().Get("app.name"), "testing111")

	assert.True(t, app.GetExceptionHandler().ShouldReport(ShouldReportException{}))
}

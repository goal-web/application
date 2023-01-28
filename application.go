package application

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/qbhy/parallel"
	"reflect"
)

const EnvProduction = "production"

type application struct {
	contracts.Container
	serviceProviders []contracts.ServiceProvider
	exceptionHandler contracts.ExceptionHandler
	config           contracts.Config
}

func (app *application) GetExceptionHandler() contracts.ExceptionHandler {
	if app.exceptionHandler == nil {
		app.exceptionHandler = app.Get("exception.handler").(contracts.ExceptionHandler)
	}
	return app.exceptionHandler
}

func (app *application) GetConfig() contracts.Config {
	if app.config == nil {
		app.config = app.Get("config").(contracts.Config)
	}
	return app.config
}

func (app *application) Environment() string {
	return app.GetConfig().Get("app").(Config).Env
}

func (app *application) IsProduction() bool {
	return app.Environment() == EnvProduction
}

func (app *application) Debug() bool {
	return app.GetConfig().Get("app").(Config).Debug
}

func (app *application) Start() map[string]error {
	errors := make(map[string]error)
	queue := parallel.NewParallel(len(app.serviceProviders))

	for _, service := range app.serviceProviders {
		(func(service contracts.ServiceProvider) {
			_ = queue.Add(func() interface{} {
				return service.Start()
			})
		})(service)
	}

	results := queue.Wait()
	for serviceIndex, result := range results {
		if err, isErr := result.(error); isErr {
			errors[utils.GetTypeKey(reflect.TypeOf(app.serviceProviders[serviceIndex]))] = err
		}
	}

	return errors
}

// Stop 倒序执行各服务的关闭
func (app *application) Stop() {
	for serviceIndex := len(app.serviceProviders) - 1; serviceIndex > -1; serviceIndex-- {
		app.serviceProviders[serviceIndex].Stop()
	}
}

// RegisterServices 顺序启动各个服务
func (app *application) RegisterServices(services ...contracts.ServiceProvider) {
	app.serviceProviders = append(app.serviceProviders, services...)

	for _, service := range services {
		service.Register(app)
	}
}

package application

import (
	"github.com/goal-web/contracts"
	"sync"
)

const EnvProduction = "production"

type application struct {
	contracts.Container
	serviceProviders []contracts.ServiceProvider
	debug            bool
}

func (app *application) Debug() bool {
	return app.debug
}

func (app *application) Start() map[contracts.ServiceProvider]error {
	var errors = make(map[contracts.ServiceProvider]error)
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, service := range app.serviceProviders {
		wg.Add(1)
		go func(service contracts.ServiceProvider) {
			defer wg.Done()
			if err := service.Start(); err != nil {
				mutex.Lock()
				errors[service] = err
				mutex.Unlock()
			}
		}(service)
	}
	wg.Wait()

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

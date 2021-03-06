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
	services []contracts.ServiceProvider
}

func (this *application) Environment() string {
	return this.Get("config").(contracts.Config).Get("app").(Config).Env
}

func (this *application) IsProduction() bool {
	return this.Environment() == EnvProduction
}

func (this *application) Debug() bool {
	return this.Get("config").(contracts.Config).Get("app").(Config).Debug
}

func (this *application) Start() map[string]error {
	errors := make(map[string]error)
	queue := parallel.NewParallel(len(this.services))

	for _, service := range this.services {
		(func(service contracts.ServiceProvider) {
			_ = queue.Add(func() interface{} {
				return service.Start()
			})
		})(service)
	}

	results := queue.Wait()
	for serviceIndex, result := range results {
		if err, isErr := result.(error); isErr {
			errors[utils.GetTypeKey(reflect.TypeOf(this.services[serviceIndex]))] = err
		}
	}

	return errors
}

func (this *application) Stop() {
	// 倒序执行各服务的关闭
	for serviceIndex := len(this.services) - 1; serviceIndex > -1; serviceIndex-- {
		this.services[serviceIndex].Stop()
	}
}

func (this *application) RegisterServices(services ...contracts.ServiceProvider) {
	this.services = append(this.services, services...)

	for _, service := range services {
		service.Register(this)
	}
}

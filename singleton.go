package application

import (
	"github.com/goal-web/container"
	"github.com/goal-web/contracts"
	"sync"
)

var (
	instance contracts.Application
	once     sync.Once
)

func Singleton() contracts.Application {
	if instance != nil {
		return instance
	}
	once.Do(func() {
		instance = &application{
			Container:        container.New(),
			serviceProviders: make([]contracts.ServiceProvider, 0),
		}
	})

	return instance
}

func SetSingleton(app contracts.Application) {
	instance = app
}

func Get(key string, args ...any) any {
	return Singleton().Get(key, args...)
}

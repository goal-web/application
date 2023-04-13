package application

import (
	"github.com/goal-web/contracts"
	"sync"
)

var (
	instance contracts.Application
	once     sync.Once
)

func Singleton(debug ...bool) contracts.Application {
	if instance != nil {
		return instance
	}
	once.Do(func() {
		instance = New(debug...)
	})

	return instance
}

func SetSingleton(app contracts.Application) {
	instance = app
}

func Get(key string, args ...any) any {
	return Singleton().Get(key, args...)
}

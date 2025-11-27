package application

import (
	"sync"

	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
)

var singleton contracts.Application
var once sync.Once

func Default(debug ...bool) contracts.Application {
	once.Do(func() {
		singleton = New(utils.DefaultBool(debug, true))
	})

	return singleton
}

func Get(key string, args ...any) any {
	return Default().Get(key, args...)
}

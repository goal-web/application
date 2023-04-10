package application

import (
	"github.com/goal-web/container"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
)

func New(debug ...bool) contracts.Application {
	return &application{
		debug:            utils.DefaultBool(debug),
		Container:        container.New(),
		serviceProviders: make([]contracts.ServiceProvider, 0),
	}
}

package application

import (
	"github.com/goal-web/container"
	"github.com/goal-web/contracts"
)

func New() contracts.Application {
	return &application{
		Container:        container.New(),
		serviceProviders: make([]contracts.ServiceProvider, 0),
	}
}

func NewWithExceptionHandler(handler contracts.ExceptionHandler) contracts.Application {
	return &application{
		Container:        container.New(),
		serviceProviders: make([]contracts.ServiceProvider, 0),
		exceptionHandler: handler,
	}
}

func NewWithConfig(config contracts.Config) contracts.Application {
	return &application{
		Container:        container.New(),
		serviceProviders: make([]contracts.ServiceProvider, 0),
		config:           config,
	}
}

func NewFully(config contracts.Config, handler contracts.ExceptionHandler) contracts.Application {
	return &application{
		Container:        container.New(),
		serviceProviders: make([]contracts.ServiceProvider, 0),
		config:           config,
		exceptionHandler: handler,
	}
}

package application

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
)

type Config struct {
	ServerId string
	Name     string
	Debug    bool
	Timezone string
	Env      string
	Locale   string
	Key      string
}

func ConfigProvider(hostname, userHome string) contracts.ConfigProvider {
	return func(env contracts.Env) interface{} {
		return Config{
			ServerId: fmt.Sprintf("%s:%s.%s", hostname, userHome, utils.RandStr(6)),
			Name:     env.GetString("app.name"),
			Debug:    env.GetBool("app.debug"),
			Timezone: env.GetString("app.timezone"),
			Env:      env.GetString("app.env"),
			Locale:   env.GetString("app.locale"),
			Key:      env.GetString("app.key"),
		}
	}
}

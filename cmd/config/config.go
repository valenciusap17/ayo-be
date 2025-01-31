package config

import (
	"context"
	"fmt"
	"os"
)

type Application struct {
	Common Common `mapstructure:"common" validate:"required"`
	Routes Routes `mapstructure:"routes" validate:"required"`
}

type Common struct {
	Postgres Postgres `mapstructure:"postgres" validate:"required"`
}

type Routes struct {
	User UserRoutes `mapstructure:"user" validate:"required"`
}

type UserRoutes struct {
	SignUp string `mapstructure:"sign-up" validate:"required"`
	SignIn string `mapstructure:"sign-in" validate:"required"`
}

type Postgres struct {
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	Port     string `mapstructure:"port" validate:"required"`
	Host     string `mapstructure:"host" validate:"required"`
	Name     string `mapstructure:"name" validate:"required"`
}

func Load() Application {
	ctx := context.Background()
    cfgManager := NewConfigManager()

    conf := Application{}

    if err := cfgManager.Start(ctx, &conf); err != nil {
        fmt.Printf("failed to load config: %v\n", err)
        os.Exit(1)
    }

    fmt.Printf("Config loaded: %+v\n", conf)
    return conf
}
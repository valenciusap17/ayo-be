package config

import (
	"context"
	"fmt"
	"os"
)

type Application struct {
	Common Common `mapstructure:"common" validate:"required"`
}

type Common struct {
	Postgres Postgres `mapstructure:"postgres" validate:"required"`
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
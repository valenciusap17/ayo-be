package config

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"github.com/tidwall/jsonc"
)

const configFileName = "config.jsonc"

type configManager struct {
    goValidator *validator.Validate
    configDecoder mapstructure.DecoderConfig
}

func (c *configManager) Start(ctx context.Context, dest any) error {
    err := c.readFile(ctx, dest)
    if err != nil {
        if !os.IsNotExist(err) {
            return fmt.Errorf("config: local: %w", err)
        }
        fmt.Printf("Warning: local config file doesn't exist: %s\n", configFileName)
    }
    return nil
}

func (c *configManager) readFile(ctx context.Context, dest any) error {
    rawLocalConfig, err := os.ReadFile(configFileName)
    if err != nil {
        return err
    }

    var localConfigMap map[string]any
    rawLocalConfig = jsonc.ToJSON(rawLocalConfig)
    
    if err = json.Unmarshal(rawLocalConfig, &localConfigMap); err != nil {
        return fmt.Errorf("failed to parse JSONC: %w", err)
    }

    //decode
    if err = c.decodeAndValidate(ctx, localConfigMap, dest); err != nil {
        return err
    }
    
    return nil
}

func (c *configManager) decodeAndValidate(ctx context.Context, input any, dest any) error {
    if err := c.decode(input, dest); err != nil {
        return fmt.Errorf("failed to decode values into %T: %w", dest, err)
    }

    if err := c.goValidator.StructCtx(ctx, dest); err != nil {
        js, _ := json.Marshal(dest)
        return fmt.Errorf("failed to validate '%s' in %T: %w", string(js), dest, err)
    }

    return nil
}

func (c *configManager) decode(input any, dest any) (err error) {
    defer func() {
        if panicked := recover(); panicked != nil {
            err = fmt.Errorf("panicked when decoding config: %v", panicked)
        }
    }()
    cd := c.configDecoder
    cd.Result = dest
    mdc, err := mapstructure.NewDecoder(&cd)
    if err != nil {
        return err
    }

    return mdc.Decode(input); 
}

func NewConfigManager() *configManager {
    return &configManager{
        goValidator: validator.New(),
        configDecoder: mapstructure.DecoderConfig{
            DecodeHook: mapstructure.DecodeHookFunc(
                mapstructure.StringToTimeDurationHookFunc(),
            ),
        },
    }
}
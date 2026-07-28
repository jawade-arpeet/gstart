package config

import "gstart/internal/constants"

type ServerConfig struct {
	RunEnv constants.Env `mapstructure:"run_env" validate:"required,oneof=dev  stg prod"`
	Port   uint          `mapstructure:"port" validate:"required"`
}

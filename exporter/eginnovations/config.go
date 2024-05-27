package eginnovations

import (
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configopaque"
)

type Config struct {
	UserID                  string              `mapstructure:"userId"`
	Token                   configopaque.String `mapstructure:"token"`
	configgrpc.ClientConfig `mapstructure:",squash"`
}

var _ component.Config = (*Config)(nil)

func (c *Config) Validate() error {
	if c.Endpoint == "" {
		return fmt.Errorf("endpoint not specified, please fix the configuration")
	}
	if c.UserID == "" || c.Token == "" {
		return fmt.Errorf("userId or token not specified,please fix the configuration")
	}
	return nil

}

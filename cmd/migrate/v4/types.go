package v4

import (
	databaseconfig "github.com/villagelabsco/juno/v4/database/config"
	loggingconfig "github.com/villagelabsco/juno/v4/logging/config"
	"github.com/villagelabsco/juno/v4/modules/pruning"
	"github.com/villagelabsco/juno/v4/modules/telemetry"
	nodeconfig "github.com/villagelabsco/juno/v4/node/config"
	parserconfig "github.com/villagelabsco/juno/v4/parser/config"
	pricefeedconfig "github.com/villagelabsco/juno/v4/pricefeed"
	"github.com/villagelabsco/juno/v4/types/config"
)

// Config defines all necessary juno configuration parameters.
type Config struct {
	Chain    config.ChainConfig    `yaml:"chain"`
	Node     nodeconfig.Config     `yaml:"node"`
	Parser   parserconfig.Config   `yaml:"parsing"`
	Database databaseconfig.Config `yaml:"database"`
	Logging  loggingconfig.Config  `yaml:"logging"`

	// The following are there to support modules which config are present if they are enabled

	Telemetry *telemetry.Config       `yaml:"telemetry,omitempty"`
	Pruning   *pruning.Config         `yaml:"pruning,omitempty"`
	PriceFeed *pricefeedconfig.Config `yaml:"pricefeed,omitempty"`
}

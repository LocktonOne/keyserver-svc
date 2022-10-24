package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type WalletsConfig struct {
	DisableConfirm bool `fig:"disable_confirm"`
}

func (c *config) WalletsConfig() WalletsConfig {
	if c.walletsConfig == nil {
		var result WalletsConfig

		err := figure.
			Out(&result).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "wallets")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out wallets"))
		}

		c.walletsConfig = &result
	}

	return *c.walletsConfig
}

package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/keyserver-svc/internal/notificator"
)

func (c *config) Notificator() *notificator.Connector {
	log := c.Log().WithField("service", "notificator")

	if c.notificator == nil {
		var result notificator.Config

		var disabled struct {
			Disabled bool `fig:"disabled"`
		}

		if err := figure.Out(&disabled).From(kv.MustGetStringMap(c.getter, "notificator")).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out notificator disabled"))
		}

		if disabled.Disabled {
			c.notificator = notificator.NewDisabledConnector(log)
			return c.notificator
		}

		err := figure.
			Out(&result).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "notificator")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out notificator"))
		}

		c.notificator = notificator.NewConnector(result, log)
	}

	return c.notificator
}

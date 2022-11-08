package config

import (
	"crypto/rand"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"math/big"
)

type TFAConfig struct {
	Dictionary string `fig:"dictionary"`
	Digits     int    `fig:"digits"`
}

func (params TFAConfig) Token() string {
	n := params.Digits
	source := params.Dictionary

	bg := big.NewInt(int64(len(source)))

	b := make([]byte, n)
	for i := range b {
		n, _ := rand.Int(rand.Reader, bg)
		b[i] = source[n.Int64()]
	}
	return string(b)
}

func (c *config) TFAConfig() TFAConfig {
	if c.tfaConfig == nil {
		result := TFAConfig{
			Dictionary: "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz",
			Digits:     8,
		}

		err := figure.
			Out(&result).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "tfa_params")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out wallets"))
		}

		c.tfaConfig = &result
	}

	return *c.tfaConfig
}

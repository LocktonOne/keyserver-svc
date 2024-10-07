package config

import (
	"fmt"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"math"
	"math/rand"
)

type TFAConfig struct {
	Digits int `fig:"digits"`
}

func (params TFAConfig) Token() string {
	digits := params.Digits

	minNum := cast.ToInt(math.Pow10(int(digits - 1)))
	maxNum := cast.ToInt(math.Pow10(int(digits)))
	fmt.Println(minNum, maxNum)
	return cast.ToString(rand.Intn(maxNum-minNum) + minNum)
}

func (c *config) TFAConfig() TFAConfig {
	if c.tfaConfig == nil {
		result := TFAConfig{
			Digits: 6,
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

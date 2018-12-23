package vipertools

import (
	"errors"

	"github.com/spf13/viper"
)

func ReadFile(cfgFile string) (err error) {
	if cfgFile == "" {
		err = errors.New("file " + cfgFile + " not found")
		return
	}

	viper.SetConfigFile(cfgFile)
	err = viper.ReadInConfig()

	return
}

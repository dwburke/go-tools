package vipertools

import (
	"errors"

	"github.com/spf13/viper"
)

// Read in specific config file
// err := ReadFile("configs/foo.yml")
func ReadFile(cfgFile string) (err error) {
	if cfgFile == "" {
		err = errors.New("file " + cfgFile + " not found")
		return
	}

	viper.SetConfigFile(cfgFile)
	err = viper.ReadInConfig()

	return
}

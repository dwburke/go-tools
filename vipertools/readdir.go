package vipertools

import (
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
)

func ReadDir(directory string) error {

	config_files, err := ioutil.ReadDir(directory)

	if err != nil {
		return err
	}

	for _, info := range config_files {
		if _, err := os.Stat(info.Name()); err == nil {
			viper.SetConfigFile(info.Name())
			if err := viper.MergeInConfig(); err != nil {
				return err
			}
		}
	}

	return nil
}

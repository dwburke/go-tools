package vipertools

import (
	"os"

	"github.com/spf13/viper"
)

// Merge config of file names passed in
// err := MergeConfigs([]string{ "configs/foo1.yml", "configs/foo2.yml" })
func MergeConfigs(config_files []string) error {
	for _, filename := range config_files {
		if _, err := os.Stat(filename); err == nil {
			viper.SetConfigFile(filename)
			if err := viper.MergeInConfig(); err != nil {
				return err
			}
		}
	}
	return nil
}

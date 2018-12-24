package vipertools

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func ReadPattern(pattern string) error {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	for _, filename := range matches {
		if _, err := os.Stat(filename); err == nil {
			viper.SetConfigFile(filename)
			if err := viper.MergeInConfig(); err != nil {
				return err
			}
		}
	}

	return nil
}

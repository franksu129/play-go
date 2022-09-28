package chatInit

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port int    `yaml:"port"`
	Chat string `yaml:"chat"`
	Wait string `yaml:"wait"`
}

func GetViperConfig(fileName string, fileType string, path string) *viper.Viper {
	// name of config file (without extension)
	viper.SetConfigName(fileName)
	//  REQUIRED if the config file does not have the extension in the name
	viper.SetConfigType(fileType)
	// path to look for the config file in
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return viper.GetViper()
}

func GetConfig(fileName string, fileType string, path string) *Config {
	var viper = GetViperConfig(fileName, fileType, path)
	resConfig := &Config{}

	err := viper.Unmarshal(resConfig)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return resConfig
}

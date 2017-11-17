package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var configPath, configType, configName string

type Configuration struct {
	Hashtags string
	Outfile  string
	Loglevel int
	LastId   int
}

func init() {
	configPath = "./"
	configType = "yaml"
	configName = "config"
}

func GetConfig() (Configuration, error) {
	return parseConfigFile()
}

func parseConfigFile() (Configuration, error) {
	var configuration Configuration

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("No configuration file loaded - using defaults")
	}

	configuration.Hashtags = viper.GetString("hashtags")
	configuration.Outfile = viper.GetString("outfile")
	configuration.Loglevel = viper.GetInt("loglevel")
	configuration.LastId = viper.GetInt("last_id")

	return configuration, nil
}

func (c *Configuration) SetHashtags(hashtags string) {
	c.Hashtags = hashtags
}

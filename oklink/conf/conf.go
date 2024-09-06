package conf

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	ApiKey  string
	BaseUrl string
	Verbose bool
	Timeout time.Duration
}

func (c *Config) New() {
	// init viper
	viper.SetConfigName("oklink-config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../../")

	// set default config value
	viper.SetDefault("Verbose", false)
	viper.SetDefault("timeout", 30)

	// try read config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Read config file error: %s", err)
	}

	// check if the required fields are provided
	if viper.GetString("apiKey") == "" {
		log.Fatalf("apiKey is required in config file")
	}
	if viper.GetString("baseUrl") == "" {
		log.Fatalf("baseUrl is required in config file")
	}

	// bound config value to struct
	c.ApiKey = viper.GetString("apiKey")
	c.BaseUrl = viper.GetString("baseUrl")
	c.Verbose = viper.GetBool("verbose")
	c.Timeout = time.Duration(viper.GetInt("timeout")) * time.Second
}

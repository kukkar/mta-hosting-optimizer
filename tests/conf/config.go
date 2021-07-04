package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	MTAHostingOptimizer struct {
		IPPort  string
		Version string
	}
	Mysql struct {
		User               string `json:"User"`
		Password           string `json:"Password"`
		DBName             string `json:"DbName"`
		MaxOpenConnections int    `json:"MaxOpenConnections"`
		MaxIdleConnections int    `json:"MaxIdleConnections"`
		DefaultTimeZone    string `json:"DefaultTimeZone"`
		Host               string `json:"Host"`
		Port               string `json:"Port"`
	} `json:"Mysql"`
}

var testConfig *Configuration

func GetConfig() (*Configuration, error) {

	if testConfig != nil {
		return testConfig, nil
	}

	config := Configuration{}
	viper.SetConfigName("test_conf")
	viper.AddConfigPath("../tests/conf")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Could not read config file, Error: %s", err.Error())
	}
	err = viper.Unmarshal(&config)

	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal config file, Error: %s", err.Error())
	}
	testConfig = &config

	return &config, nil
}

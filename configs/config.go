package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Conf struct {
	API_PORT           string `mapstructure:"API_HTTP_PORT"`
	ORCHESTRATOR_HOST  string `mapstructure:"ORCHESTRATOR_HOST"`
	ORCHESTRATOR_PORT  string `mapstructure:"ORCHESTRATOR_PORT"`
	WEATHERMAP_API_KEY string `mapstructure:"WEATHERMAP_API_KEY"`
	APP_NAME           string `mapstructure:"APP_NAME"`
	APP2_NAME          string `mapstructure:"APP2_NAME"`
	COLLECTOR_HOST     string `mapstructure:"COLLECTOR_HOST"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf
	var err error

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	viper.BindEnv("API_HTTP_PORT")
	viper.BindEnv("ORCHESTRATOR_HOST")
	viper.BindEnv("ORCHESTRATOR_PORT")
	viper.BindEnv("WEATHERMAP_API_KEY")
	viper.BindEnv("APP_NAME")
	viper.BindEnv("APP2_NAME")
	viper.BindEnv("COLLECTOR_HOST")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("WARNING: %v\n", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, err
}

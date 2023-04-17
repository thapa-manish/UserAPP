package config

import "github.com/spf13/viper"

type Config struct {
	DbUser     string
	DbPassword string
	DbName     string
	DbHost     string
	DbPort     string
	HttpPort   string
}

func NewConfig() (*Config, error) {
	// Create a new instance of Viper
	v := viper.New()

	v.SetConfigFile("./env/.env")

	// Read in the configuration file
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	config := &Config{
		DbUser:     v.Get("DB_USER").(string),
		DbPassword: v.Get("DB_PASSWORD").(string),
		DbName:     v.Get("DB_NAME").(string),
		DbHost:     v.Get("DB_HOST").(string),
		DbPort:     v.Get("DB_PORT").(string),
		HttpPort:   v.Get("HTTP_PORT").(string),
	}

	return config, nil
}

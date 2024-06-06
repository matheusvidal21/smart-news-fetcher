package configs

import (
	"github.com/spf13/viper"
)

var cfg *Conf

type Conf struct {
	DBDriver             string `mapstructure:"DB_DRIVER"`
	DBHost               string `mapstructure:"DB_HOST"`
	DBPort               string `mapstructure:"DB_PORT"`
	DBUser               string `mapstructure:"DB_USER"`
	DBPassword           string `mapstructure:"DB_PASSWORD"`
	DBSource             string `mapstructure:"DB_SOURCE"`
	DBName               string `mapstructure:"DB_NAME"`
	WebServerPort        string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecretKey         string `mapstructure:"JWT_SECRET_KEY"`
	JWTExpirationMinutes string `mapstructure:"JWT_EXPIRATION_MINUTES"`
	SMTP_HOST            string `mapstructure:"SMTP_HOST"`
	SMTP_PORT            string `mapstructure:"SMTP_PORT"`
	SMTP_USER            string `mapstructure:"SMTP_USER"`
	SMTP_PASSWORD        string `mapstructure:"SMTP_PASSWORD"`
	SMTP_FROM_EMAIL      string `mapstructure:"SMTP_FROM_EMAIL"`
}

func LoadConfigs(path string) *Conf {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

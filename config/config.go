package config

import "github.com/ilyakaznacheev/cleanenv"

type (
	WeatherAPI struct {
		ApiKey string `env:"WEATHER_API_KEY" required:"true"`
	}

	MySQL struct {
		Host     string `env:"MYSQL_HOST" required:"true"`
		Port     string `env:"MYSQL_PORT" required:"true"`
		Username string `env:"MYSQL_USER" required:"true"`
		Password string `env:"MYSQL_PASSWORD" required:"true"`
		DbName   string `env:"MYSQL_DB_NAME" required:"true"`
	}

	Config struct {
		Db         MySQL
		WeatherAPI WeatherAPI
	}
)

func LoadConfig() Config {
	var cfg Config

	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		panic(err)
	}

	return cfg
}

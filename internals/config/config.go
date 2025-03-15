package config

import "github.com/ilyakaznacheev/cleanenv"


var AppConfig *Config


type ServerConfig struct {
	AppName string `env:"APP_NAME"`
	Port string  `env:"APP_PORT"`
}


type DatabaseConfig struct {
	DatabaseName string `env:"DB_NAME"`
}

type Config struct {
	Server ServerConfig
	Database DatabaseConfig

}

func LoadConfigFromfile(filePath string) (*Config, error) {
	if AppConfig != nil {
		return AppConfig, nil
	}

	AppConfig = &Config{}
	if err := cleanenv.ReadConfig(filePath, AppConfig); err != nil {
		return nil, err
	}

	return AppConfig, nil
}

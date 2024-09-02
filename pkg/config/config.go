package config

import "github.com/spf13/viper"

// Loader load config from reader into Viper
type Loader interface {
	Load(viper.Viper) (*viper.Viper, error)
}

type Config struct {
	// db
	Postgres DBConnection

	// server
	ApiServer ApiServer

	// service
	APIKey       string
	Debug        bool
	Env          string
	JWTSecretKey string
}

type DBConnection struct {
	Host string
	Port string
	User string
	Name string
	Pass string

	SSLMode string
}

type ApiServer struct {
	Port           string
	AllowedOrigins string
}

func DefaultConfigLoaders() []Loader {
	var loaders []Loader
	fileLoader := NewFileLoader(".env", ".")
	loaders = append(loaders, fileLoader)
	loaders = append(loaders, NewENVLoader())

	return loaders
}

// LoadConfig load config from loader list
//
// Example:
// cfg := config.LoadConfig(config.DefaultConfigLoaders())
//
// log := logger.NewLogrusLogger()
//
// log.Infof("Server starting", cfg)
func LoadConfig(loaders []Loader) *Config {
	return &Config{
		ApiServer: ApiServer{
			Port:           "8080",
			AllowedOrigins: "*",
		},
	}
}

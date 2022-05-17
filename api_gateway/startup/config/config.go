package config

type Config struct {
	Port               string
	AuthenticationHost string
	AuthenticationPort string
	PostHost           string
	PostPort           string
}

func NewConfig() *Config {
	return &Config{
		Port:               "8080",
		AuthenticationHost: "localhost",
		AuthenticationPort: "8081",
		PostHost:           "localhost",
		PostPort:           "8082",
	}
}

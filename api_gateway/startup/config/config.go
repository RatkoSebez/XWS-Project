package config

type Config struct {
	Port               string
	AuthenticationHost string
	AuthenticationPort string
	PostHost           string
	PostPort           string
	RegistrationHost   string
	RegistrationPort   string
	ProfileHost        string
	ProfilePort        string
	FollowHost         string
	FollowPort         string
	JobOfferHost       string
	JobOfferPort       string
}

func NewConfig() *Config {
	return &Config{
		Port:               "8080",
		AuthenticationHost: "localhost",
		AuthenticationPort: "8081",
		PostHost:           "localhost",
		PostPort:           "8082",
		RegistrationHost:   "localhost",
		RegistrationPort:   "8083",
		ProfileHost:        "localhost",
		ProfilePort:        "8084",
		FollowHost:         "localhost",
		FollowPort:         "8085",
		JobOfferHost:       "localhost",
		JobOfferPort:       "8086",
	}
}

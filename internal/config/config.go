package config

type Config struct {
	Port           string
	PortPostgreSQL string
	PortRedis      string
}

func NewConfig() *Config {
	return &Config{
		Port:           ":8080",
		PortPostgreSQL: "host=127.0.0.1 port=5432 user=postgres sslmode=disable password=1234",
		PortRedis:      "127.0.0.1:6379",
	}
}

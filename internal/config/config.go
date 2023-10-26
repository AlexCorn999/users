package config

type Config struct {
	Port   string
	PortDB string
}

func NewConfig() *Config {
	return &Config{
		Port:   ":8080",
		PortDB: "host=127.0.0.1 port=5432 user=postgres sslmode=disable password=1234",
	}
}

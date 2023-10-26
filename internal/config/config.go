package config

type Config struct {
	Port   string
	PortDB string
}

func NewConfig() *Config {
	return &Config{
		Port:   ":8080",
		PortDB: ":8080",
	}
}

package config

import (
	"flag"
	"fmt"
)

type Config struct {
	Port           string
	PortPostgreSQL string
	PortRedis      string
}

func NewConfig() *Config {
	return &Config{
		Port:           ":8080",
		PortPostgreSQL: "host=127.0.0.1 port=5432 user=postgres sslmode=disable password=1234",
	}
}

// ParseFlags обрабатывает аргументы командной строки и сохраняет их значения в PortRedis.
func (c *Config) ParseFlags() {
	// 127.0.0.1:6379
	host := flag.String("host", "", "host for redis")
	port := flag.String("port", "", "port for redis")
	flag.Parse()

	addr := fmt.Sprintf("%s:%s", *host, *port)
	c.PortRedis = addr
}

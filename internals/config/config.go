package config

import (
	"flag"
	"time"
)

type Config struct {
	Port      string
	Limit     int
	Window    time.Duration
	RedisAddr string
}

func SetFlags() *Config {
	port := flag.String("port", "8080", "Server port")
	limit := flag.Int("limit", 10, "Number if allowed request")
	window := flag.Int("window", 60, "Rate limit window in seconds")
	redisAddr := flag.String("redis", "localhost:6379", "Redis address")

	flag.Parse()
	cfg := &Config{
		Port:      ":" + *port,
		Limit:     *limit,
		Window:    time.Duration(*window) * time.Second,
		RedisAddr: *redisAddr,
	}
	validate(cfg)
	return cfg
}

func validate(cfg *Config) {
	if cfg.Limit <= 0 {
		panic("limit must be greater than 0")
	}

	if cfg.Window <= 0 {
		panic("limit must be greater than 0")
	}

	if cfg.RedisAddr == "" {
		panic("redis address cannot be empty")
	}
}

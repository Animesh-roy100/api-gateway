package config

type Config struct {
	Server struct {
		Port int
		Host string
	}
	Services map[string]ServiceConfig
	Redis    struct {
		Host string
		Port int
	}
	RateLimit struct {
		RequestsPerSecond int
		Burst             int
	}
}

type ServiceConfig struct {
	URL         string
	Timeout     int
	RateLimit   int
	RequireAuth bool
}

func LoadConfig() (*Config, error) {
	// TODO: Implement configuration loading from file or environment
	return &Config{}, nil
}

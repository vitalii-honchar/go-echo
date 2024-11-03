package config

type Config struct {
	TemplatesDir  string
	WebServerPort int
}

func NewConfig() *Config {
	return &Config{
		TemplatesDir:  "../../web/templates",
		WebServerPort: 1323,
	}
}

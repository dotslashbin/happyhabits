package configs

type Config struct {
	Version string
}

func (config *Config) Init() {
	config.Version = "1.0.0"
}

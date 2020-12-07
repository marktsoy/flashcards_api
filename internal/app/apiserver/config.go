package apiserver

// Config TODO
type Config struct {
	BindAddr string `toml:"bind_addr"`

	LoggerType string `toml:"logger"`
	LogPath    string `toml:"log_path"`

	DatabaseURL string `toml:"store"`
}

//NewConfig TODO
func NewConfig() *Config {
	return &Config{
		BindAddr:    "127.0.0.1:8080",
		LoggerType:  "dev",
		LogPath:     "apiserver.log",
		DatabaseURL: "",
	}
}

package types

// ServerConfig ...
type ServerConfig struct {
	Server struct {
		ListenAddr    string
		MaxRetryCount int
	}
	DB struct {
		IP       string
		Port     int
		User     string
		Password string
		DBName   string
	}
}

package types

// ServerConfig ...
type ServerConfig struct {
	Server struct {
		ListenAddr    string
		MaxRetryCount int
	}
	DB struct {
		DBAddr   string
		IP       string
		Port     int
		User     string
		Password string
		DBName   string
	}
}

package server

// ServerConfig конфигурация сервера
type ServerConfig struct {
}

// NewServerConfig создает новый экземпляр ServerConfig
func NewServerConfig() *ServerConfig {
	return &ServerConfig{}
}

// LoadServerConfig загружает конфигурацию сервера из файла
func LoadServerConfig() (*ServerConfig, error) {
	return NewServerConfig(), nil
}

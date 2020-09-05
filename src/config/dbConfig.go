package config

// Config for Database configuration struct
type Config struct {
	DB *DBConfig
}

// DBConfig Databse configuration properties
type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	Charset  string
}

// GetDBConfigurations database properties setup
func GetDBConfigurations() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "root",
			Password: "admin123",
			DBName:   "app",
			Charset:  "utf8",
		},
	}
}

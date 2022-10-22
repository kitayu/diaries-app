package config

import "os"

var Config *config

type config struct {
	DBUsername string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string

	ServerPort string
}

func init() {
	Config = &config{
		DBUsername: os.Getenv("MYSQL_USER"),
		DBPassword: os.Getenv("MYSQL_PASSWORD"),
		DBName:     os.Getenv("MYSQL_DATABASE"),
		DBHost:     os.Getenv("MYSQL_HOST"),
		DBPort:     os.Getenv("MYSQL_PORT"),
		ServerPort: os.Getenv("SERVER_PORTS"),
	}
}

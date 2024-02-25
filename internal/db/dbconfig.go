package db

import "fmt"

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func NewDBConfig(host string, port int, user string, password string, dbName string) DBConfig {
	return DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
	}
}

func (c *DBConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Sao_Paulo",
		c.Host, c.User, c.Password, c.DBName, c.Port,
	)
}

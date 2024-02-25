package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB

type DBConn struct {
	config DBConfig
}

func NewDBConn(config DBConfig) DBConn {
	return DBConn{config: config}
}

func (d *DBConn) Open() error {
	var err error
	Instance, err = gorm.Open(postgres.Open(d.config.GetDSN()), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func (d *DBConn) GetDBConn() *gorm.DB {
	return Instance
}

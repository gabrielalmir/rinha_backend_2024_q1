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

func (d *DBConn) Connect() error {
	if Instance != nil {
		println("Database connection already established. Skipping ...")
		return nil
	}

	var err error
	Instance, err = d.Open()
	if err != nil {
		return err
	}

	return nil
}

func (d *DBConn) Open() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(d.config.GetDSN()), &gorm.Config{})
}

func (d *DBConn) GetDBConn() *gorm.DB {
	return Instance
}

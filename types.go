package gorm_db_init

import "gorm.io/gorm"

type MySql struct {
	Dsn string
}

type GormInitConf map[string]MySql

type inst map[string]*gorm.DB

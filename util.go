package gorm_db_init

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Get(name string) *gorm.DB {
	return instance[name]
}

func Init(c GormInitConf, gormConfigs ...*gorm.Config) {
	conf = c
	var db *gorm.DB
	var err error
	for name, mc := range conf {
		if len(gormConfigs) > 0 {
			db, err = gorm.Open(mysql.Open(mc.Dsn), gormConfigs[0])
		} else {
			db, err = gorm.Open(mysql.Open(mc.Dsn))
		}

		if err != nil {
			panic(err)
		}

		instance[name] = db
	}
}

package db

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormDB *gorm.DB

func Connt(dsn string, driver string) error {
	var err error
	var dialector gorm.Dialector
	if driver == "sqlite" {
		dialector = sqlite.Open(dsn)
	} else {
		dialector = postgres.Open(dsn)
	}
	gormDB, err = gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	return err
}

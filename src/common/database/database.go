package database

import (
	"fiber/src/common"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func init() {
	var datetimePrecision = 2
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", common.Config.MariadbUsername, common.Config.MariadbPassword, common.Config.MariadbHost, common.Config.MariadbPort, common.Config.MariadbDatabase)

	Connection, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DefaultStringSize:        256,
		DefaultDatetimePrecision: &datetimePrecision,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

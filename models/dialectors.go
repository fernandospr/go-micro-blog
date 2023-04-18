package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SqliteDialector(filename string) gorm.Dialector {
	return sqlite.Open(filename)
}

func MySqlDialector(user string, password string, host string, port string, dbname string) gorm.Dialector {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	return mysql.Open(dsn)
}

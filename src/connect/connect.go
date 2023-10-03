package connect

import (
	"tategoto/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// open
func GetConnection() *gorm.DB {
	dsn := config.Config.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

// close
func CloseConnection(db *gorm.DB) {
	if sqlDb, err := db.DB(); err != nil {
		panic(err)
	} else {
		if err := sqlDb.Close(); err != nil {
			panic(err)
		}
	}
}

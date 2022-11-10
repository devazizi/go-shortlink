package adapter

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Conn *gorm.DB
}

func NewDB(dsn string) DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("can not connect to db")
	}

	entities := []any{
		Link{},
	}

	migrationErr := db.AutoMigrate(entities...)
	if migrationErr != nil {
		panic("auto migration fail")
	}

	return DB{Conn: db}
}

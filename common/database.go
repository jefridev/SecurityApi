package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Initilice sqlite packages.
)

// Database struct reference.
type Database struct {
	*gorm.DB
}

// DB Holder.
var DB *gorm.DB

// Init opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../securityApi.db")
	if err != nil {
		fmt.Println("DB Error", err)
	}
	db.DB().SetMaxIdleConns(10)
	DB = db
	return DB
}

// GetDB returns a pointer to current open db connection.
func GetDB() *gorm.DB {
	return DB
}

// package controllers

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// func connect() *sql.DB {
// 	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latihan_pbp?parseTime=true&loc=Asia%2FJakarta")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return db
// }

// func connectGorm() *gorm.DB {
// 	sqlDB, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latihan_pbp")
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
// 		return nil
// 	}

// 	gormDB, err := gorm.Open(mysql.New(mysql.Config{
// 		Conn: sqlDB,
// 	}), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to initialize GORM: %v", err)
// 		return nil
// 	}

// 	return gormDB
// }

package controllers

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latihan_pbp?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func connectGorm() (*gorm.DB, error) {
	dsn := "root:@tcp(localhost:3306)/db_latihan_pbp?parseTime=true&loc=Asia%2FJakarta"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
	"github.com/joho/godotenv"
	"github.com/takumi2786/takumiweb/models"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	// 環境変数を読み込み
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	uri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, dbname)
	db, err = gorm.Open("postgres", uri)
	if err != nil {
		panic(err)
	}
	autoMigration()

	// 初期データを投入
	default_user := models.User{
		ID:   1,
		Name: "admin",
	}
	db.Create(&default_user)
}

func autoMigration() {
	db.AutoMigrate(&models.User{})
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

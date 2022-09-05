package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mohammadDV/backend_gin/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() *gorm.DB {
	env := godotenv.Load()
	if env != nil {
		panic("We don't have any evn")
	}

	dbUser := os.Getenv("DBUSER")
	dbHost := os.Getenv("DBHOST")
	dbPss := os.Getenv("DBPASS")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		dbHost, dbUser, dbPss, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Can not connect to pgdb")
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic("Unable to migrate")
	}

	fmt.Println("Connect to db")

	return db

}

func ClosePostgres(db *gorm.DB) {

	dbPostgre, err := db.DB()
	if err != nil {
		panic("Can not close connection")
	}

	err = dbPostgre.Close()

	if err != nil {
		panic("Can not close connection")
	}

}

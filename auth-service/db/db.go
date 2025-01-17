package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// type User struct{
// 	gorm.Model
// 	Email string 	`gorm:"unique"`
// 	Password string `gorm:"not null"`
// }

func InitDB(dsn string){
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected!")
}
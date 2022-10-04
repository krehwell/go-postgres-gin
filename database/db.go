package database

import (
	"fmt"
	"rest-api-practice/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() (Database, error) {
	host := "localhost"
	port := 5432
	user := "postgres"
	// password
	dbName := "golangTestDB"

	dbConnectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbName)

	db, dbConnectErr := gorm.Open(postgres.Open(dbConnectionString), &gorm.Config{})
	if dbConnectErr != nil {
		return Database{}, dbConnectErr
	}

	if dbMigrateErr := db.AutoMigrate(models.Person{}); dbMigrateErr != nil {
		return Database{}, dbMigrateErr
	}

	return Database{db: db}, nil
}

type Database struct {
	db *gorm.DB
}

func (d Database) GetUsers() []models.Person {
	persons := []models.Person{}
	d.db.Find(&persons)

	return persons
}

func (d Database) CreateUser(newUser models.Person) {
	err := d.db.Create(&newUser).Error
	if err != nil {
		fmt.Println("Error in creating new user", err)
	}

	fmt.Println("New user created:", newUser)
}

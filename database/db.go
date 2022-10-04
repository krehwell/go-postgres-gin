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

	if dbMigrateErr := db.AutoMigrate(models.User{}); dbMigrateErr != nil {
		return Database{}, dbMigrateErr
	}

	return Database{db: db}, nil
}

type Database struct {
	db *gorm.DB
}

func (d Database) GetUsers() []models.User {
	users := []models.User{}
	d.db.Find(&users)

	return users
}

func (d Database) GetUserById(id int) (models.User, error) {
	user := models.User{ID: id}
	if err := d.db.First(&user).Error; err != nil {
		fmt.Println("Error in getting the user with ID:", id, err)
		return models.User{}, err
	}

	return user, nil
}

func (d Database) CreateUser(newUser models.User) {
	err := d.db.Create(&newUser).Error
	if err != nil {
		fmt.Println("Error in creating new user", err)
		return
	}

	fmt.Println("New user created:", newUser)
}

func (d Database) UpdateUserWithNewUserData(user models.User, newUserData models.User) {
	err := d.db.Model(&user).Updates(newUserData).Error
	if err != nil {
		fmt.Println("Error in updating user", err)
		return
	}

	fmt.Println("One user updated:")
	fmt.Println("prev data", user)
	fmt.Println("updated data", newUserData)
}

func (d Database) DeleteUserById(id int) {
	userToBeDeleted := models.User{ID: id}
	err := d.db.Delete(&userToBeDeleted).Error
	if err != nil {
		fmt.Println("Error in deleting user", err)
		return
	}
	fmt.Println("User with id:", id, "successfully deleted")
}

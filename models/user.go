package models

type User struct {
	ID        int    `gorm:"primaryKey;autoIncrement:true"`
	FirstName string `gorm:"not null;type:varchar(255)"`
	LastName  string `gorm:"not null;type:varchar(255)"`
}


package models

type Person struct {
	ID        int    `gorm:"primaryKey"`
	FirstName string `gorm:"not null;type:varchar(255)"`
	LastName  string `gorm:"not null;type:varchar(255)"`
}


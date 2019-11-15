package models

import "time"

type DBConnInfo struct {
	USERNAME     string
	USERPASSWORD string
	DBHOST       string
	DBNAME       string
}

type Account struct {
	ID        int64     `gorm:"type:int AUTO_INCREMENT;primary_key"`
	LineID    string    `gorm:"type:varchar(10);not null"`
	Password  string    `gorm:"type:varchar(10);not null"`
	Token     string    `gorm:"type:varchar(10);"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

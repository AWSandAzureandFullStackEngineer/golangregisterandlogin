package entity

import (
	"github.com/jinzhu/gorm"
)


type User struct {
	gorm.Model
	ID        uint      `gorm:"AUTO_INCREMENT" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `gorm:"unique" json:"username"`
	Email     string    `gorm:"unique" json:"email"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	Password  string    `json:"-"`
}

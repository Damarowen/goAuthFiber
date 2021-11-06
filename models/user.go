
package models


type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null" validate:"required"`
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CheckUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
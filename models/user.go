
package models


type User struct {
	Id       uint   `json:"id" `
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" gorm:"unique" valid:"email"`
	Password []byte `json:"-"`
}
package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	// Id	 		int  `sql:"type:int;primary key"`
	Username  string
	Password  string
	FirstName string
	LastName  string
	Phone     string
	// CreatedAt	string
	// ModifiedAt	string
	Status bool
	Addrs  []UserAddress
}

type UserAddress struct {
	gorm.Model
	// Id	 			int    `sql:"type:int"`
	UserId       int `gorm:"column:user_id;not_null"`
	AddressLine1 string
	AddressLine2 string
	City         string
	PostalCode   string
	Country      string
	Phone        string
	Telephone    string
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &UserAddress{})
	return db
}

func (u *User) Disable() {
	u.Status = false
}

func (u *User) Enable() {
	u.Status = true
}

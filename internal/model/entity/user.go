package entity

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "github.com/jmoiron/sqlx"
	// "GO_APP/internal/queries"

)

type User struct {
	ID        int	    `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at"`
	Username  string	`db:"username" json:"username"`
	Password  string	`db:"password" json:"password"`
	FirstName string	`db:"first_name" json:"first_name"`
	LastName  string	`db:"last_name" json:"last_name"`
	Phone     string	`db:"phone" json:"phone"`
	Status bool			`db:"status" json:"status"`
	Addrs  []UserAddress
}

type UserAddress struct {
	ID        int	    `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at"`
	UserId       int 	`db:"user_id" json:"user_id"`
	AddressLine1 string `db:"address_line1" json:"address_line1"`
	AddressLine2 string	`db:"address_line2" json:"address_line2"`
	City         string	`db:"city" json:"city"`	
	PostalCode   string	`db:"postal_code" json:"postal_code"`
	Country      string	`db:"country" json:"country"`
	Phone        string	`db:"phone" json:"phone"`
	Telephone    string	`db:"telephone" json:"telephone"`
}

// // DBMigrate will create and migrate the tables, and then make the some relationships if necessary
// func DBMigrate(db *sqlx.DB) *sqlx.DB {	
// 	db.MustExec(queries.CreateDB)
// 	return db
// }

func (u *User) Disable() {
	u.Status = false
}

func (u *User) Enable() {
	u.Status = true
}

package user

import "time"

type User struct {
	ID            string    `db:"id" json:"id"`
	Email         string    `db:"email" json:"email"`
	Password      string    `db:"password" json:"password"`
	Username      string    `db:"username" json:"username"`
	Fullname      string    `db:"fullname" json:"fullname"`
	PhoneNumber   string    `db:"phone_number" json:"phone_number"`
	CreatedDate   time.Time `db:"created_date" json:"created_date"`
	ModifiedDate time.Time `db:"modified_date" json:"modified_date"`
}
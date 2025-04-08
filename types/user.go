package types

import "time"

type User struct {
	Id          int       `sql:"id"`
	SecondaryId int       `sql:"secondary_id"`
	Email       string    `sql:"email"`
	CreatedAt   time.Time `sql:"created_at"`
	UpdatedAt   time.Time `sql:"updated_at"`
}

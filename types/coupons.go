package types

import "time"

type Coupon struct {
	ID                 int       `sql:"id"`
	Code               string    `sql:"code"`
	Discount           float64   `sql:"discount"`
	ExpiresAt          time.Time `sql:"exp"`
	Status             string    `sql:"status"`
	Remaining          int       `sql:"remaining"`
	CreatedAt          time.Time `sql:"created_at"`
	UpdatedAt          time.Time `sql:"updated_at"`
	FormattedCreatedAt string
	FormattedUpdatedAt string
	FormattedExpiresAt string
}

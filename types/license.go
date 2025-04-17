package types

import "time"

type License struct {
	Id                 int       `sql:"id"`
	Uuid               string    `sql:"uuid"`
	UserId             int       `sql:"user_id"`
	ValidTill          time.Time `sql:"valid_till"`
	GracePeriod        int       `sql:"grace_period"`
	CreatedAt          time.Time `sql:"created_at"`
	UpdatedAt          time.Time `sql:"updated_at"`
	FormattedCreatedAt string
	FormattedUpdatedAt string
	FormattedValidTill string
}

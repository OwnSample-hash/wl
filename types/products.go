package types

import "time"

type Product struct {
	ID                     int       `sql:"id"`
	Name                   string    `sql:"name"`
	Description            string    `sql:"description"`
	PricePerMonth          float64   `sql:"price_per_month"`
	Price                  float64   `sql:"price"`
	OneMonth               bool      `sql:"one_month"`
	LifeTime               bool      `sql:"life_time"`
	Discount               float64   `sql:"discount"`
	CreatedAt              time.Time `sql:"created_at"`
	UpdatedAt              time.Time `sql:"updated_at"`
	FormattedCreatedAt     string
	FormattedUpdatedAt     string
	Path                   string
	DiscountPrice          float64
	DiscountedMonthlyPrice float64
}

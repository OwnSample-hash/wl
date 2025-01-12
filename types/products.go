package types

import "time"

type Product struct {
	ID          int       `sql:"id"`
	Name        string    `sql:"name"`
	Description string    `sql:"description"`
	Price       float64   `sql:"price"`
	OneMonth    bool      `sql:"one_month"`
	LifeTime    bool      `sql:"life_time"`
	CreatedAt   time.Time `sql:"created_at"`
	UpdatedAt   time.Time `sql:"updated_at"`
}

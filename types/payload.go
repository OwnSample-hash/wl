package types

import (
	"html/template"
)

type Payload struct {
	Title     string
	SteamID   string
	IsAdmin   bool
	Path      string
	User      SteamUser
	Prods     []Product
	Coupons   []Coupon
	License   []License
	Users     []User
	CsrfField template.HTML
	CsrfToken string
}

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
	CsrfField template.HTML
	CsrfToken string
}

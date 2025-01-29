package types

type IPayment interface {
	ProcessPayment() bool
	PaymentHandler() any
	GetPath() string
	GetName() string
	GetHoverText() string
	GetImage() string
	GetFinal() float64
}

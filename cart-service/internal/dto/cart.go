package dto

type CartItem struct {
	ProductId   int32
	ProductName string
	Price       int32
	Description string
	Amounts     int32
	Image       string
	Status      string
}

type Cart struct {
	UserID int32
	Items  []*CartItem
}

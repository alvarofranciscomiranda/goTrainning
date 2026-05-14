package files

import "encoding/json"

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

func getCartFromJson(jsonString string) []CartItem {
	var cart []CartItem
	err := json.Unmarshal([]byte(jsonString), &cart)
	CheckError(err)
	return cart
}

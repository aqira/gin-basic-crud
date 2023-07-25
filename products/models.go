package products

type Product struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}

var currentId uint = 1

var productsMap = make(map[uint]Product)

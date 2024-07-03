package types

type ProductRequest struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func (p *ProductRequest) RequestToProduct() Product {
	return Product{
		Name:  p.Name,
		Price: p.Price,
	}
}

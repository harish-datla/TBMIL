package product

type Product struct{
	Name string // In go variable starting with capital letters are public
	price float64 // any variable that starts with small letter is not public
}

func NewProduct(name string, price float64) *Product{
	p:= &Product{Name: name}
	p.SetPrice(price)
	return p
}

func (p *Product) GetPrice() float64 {
	return p.price
}

func (p *Product) SetPrice(price float64) {
	p.price = price
}

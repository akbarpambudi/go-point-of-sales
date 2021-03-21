package query

type Product struct {
	ID       string
	Name     string
	Category string
	Variants []Variant
}

type Variant struct {
	ID    string
	Code  string
	Name  string
	Price float64
}

type Category struct {
	ID   string
	Name string
}

type Products []Product

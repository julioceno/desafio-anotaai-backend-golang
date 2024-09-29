package catalog_domain

type Product struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type Category struct {
	CategoryTitle       string    `json:"category_title"`
	CategoryDescription string    `json:"category_description"`
	Items               []Product `json:"items"`
}

type Catalog struct {
	OwnerId string     `json:"owner_id"`
	Catalog []Category `json:"catalog"`
}

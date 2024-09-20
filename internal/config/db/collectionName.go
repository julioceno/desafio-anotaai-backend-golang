package db

type CollectionsNames struct {
	CATEGORY string
	PRODUCT  string
}

var CollectionsName = CollectionsNames{
	CATEGORY: "categories",
	PRODUCT:  "products",
}

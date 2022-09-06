package product_domain

import "time"

type Category struct {
	ID           int
	Name         string
	CategorySlug string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Product struct {
	UUID         string
	Name         string
	ProductSlug  string
	CategorySlug string
	Image        string
	Price        uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Detail struct {
	ID          int
	ProductSlug string
	Name        string
	DetailSlug  string
	Area        string
	Qty         uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Business interface {
	// Create
	CreateCategory(domain Category) error
	CreateProduct(catSlug string, domain Product, file interface{}) error
	CreateDetail(proSlug string, domain Detail) error
	// Read
	GetCategory(catSlug string) (Category, error)
	GetProduct(proSlug string) (Product, error)
	GetDetail(detSlug string) (Detail, error)
}

type Repository interface {
	StoreCategory(domain Category) error
	StoreProduct(domain Product) error
	StoreDetail(domain Detail) error
	GetCategorySlug(catSlug string) (Category, error)
	GetProductSlug(proSlug string) (Product, error)
	GetDetailSlug(detSlug string) (Detail, error)
}

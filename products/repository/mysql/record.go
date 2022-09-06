package product_mysql

import (
	product_domain "gozzafadillah/products/domain"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID           int
	Name         string
	CategorySlug string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func ToDomainCategory(rec Category) product_domain.Category {
	return product_domain.Category{
		ID:           rec.ID,
		Name:         rec.Name,
		CategorySlug: rec.CategorySlug,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
}

type Product struct {
	gorm.Model
	UUID         string
	Name         string
	ProductSlug  string
	CategorySlug string
	Image        string
	Price        uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func ToDomainProduct(rec Product) product_domain.Product {
	return product_domain.Product{
		UUID:         rec.UUID,
		Name:         rec.Name,
		ProductSlug:  rec.ProductSlug,
		CategorySlug: rec.CategorySlug,
		Image:        rec.Image,
		Price:        rec.Price,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
}

type Detail struct {
	gorm.Model
	ID          int
	ProductSlug string
	Name        string
	DetailSlug  string
	Area        string
	Qty         uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ToDomainDetail(rec Detail) product_domain.Detail {
	return product_domain.Detail{
		ID:          rec.ID,
		ProductSlug: rec.ProductSlug,
		Name:        rec.Name,
		DetailSlug:  rec.DetailSlug,
		Area:        rec.Area,
		Qty:         rec.Qty,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

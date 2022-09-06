package product_request

import product_domain "gozzafadillah/products/domain"

type Category struct {
	Name         string `json:"name" form:"name" validate:"required"`
	CategorySlug string
}

func ToDomainCategory(req Category) product_domain.Category {
	return product_domain.Category{
		Name:         req.Name,
		CategorySlug: req.CategorySlug,
	}
}

type Product struct {
	UUID         string
	Name         string `json:"name" form:"name" validate:"required"`
	ProductSlug  string
	CategorySlug string
	Image        string
	Price        uint        `json:"price" form:"price" validate:"required"`
	File         interface{} `json:"file,omitempty"`
}

func ToDomainProduct(req Product) product_domain.Product {
	return product_domain.Product{
		UUID:         req.UUID,
		Name:         req.Name,
		ProductSlug:  req.ProductSlug,
		CategorySlug: req.CategorySlug,
		Image:        req.Image,
		Price:        req.Price,
	}
}

type Detail struct {
	Name       string `json:"name" form:"name" validate:"required"`
	DetailSlug string
	Area       string `json:"area" form:"area" validate:"required"`
	Qty        uint   `json:"qty" form:"qty" validate:"required"`
}

func ToDomainDetail(req Detail) product_domain.Detail {
	return product_domain.Detail{
		ProductSlug: req.DetailSlug,
		Name:        req.Name,
		DetailSlug:  req.DetailSlug,
		Area:        req.Area,
		Qty:         req.Qty,
	}
}

package product_business

import (
	product_domain "gozzafadillah/products/domain"
	"gozzafadillah/products/helper/claudinary"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type ProductBusiness struct {
	ProductRepo product_domain.Repository
}

// CreateCategory implements product_domain.Business
func (pb ProductBusiness) CreateCategory(domain product_domain.Category) error {
	// make slug
	domain.CategorySlug = slug.Make(domain.Name)

	// store data
	err := pb.ProductRepo.StoreCategory(domain)
	if err != nil {
		return err
	}
	return err
}

// CreateDetail implements product_domain.Business
func (pb ProductBusiness) CreateDetail(proSlug string, domain product_domain.Detail) error {
	// make slug
	domain.ProductSlug = proSlug
	domain.DetailSlug = slug.Make(domain.Name)
	// store data
	err := pb.ProductRepo.StoreDetail(domain)
	if err != nil {
		return err
	}
	return nil
}

// CreateProduct implements product_domain.Business
func (pb ProductBusiness) CreateProduct(catSlug string, domain product_domain.Product, file interface{}) error {
	// make slug
	domain.CategorySlug = catSlug
	domain.ProductSlug = slug.Make(domain.Name)
	// make uuid
	domain.UUID = uuid.NewString()
	// upload image
	img, _ := claudinary.ImageUploadHelper(file, "products")

	domain.Image = img
	if domain.Image == "" {
		domain.Image = "https://res.cloudinary.com/dt91kxctr/image/upload/v1655825545/go-bayeue/users/download_o1yrxx.png"
	}
	// store data
	err := pb.ProductRepo.StoreProduct(domain)
	if err != nil {
		return err
	}
	return nil
}

// GetCategory implements product_domain.Business
func (pb ProductBusiness) GetCategory(catSlug string) (product_domain.Category, error) {
	dataCat, err := pb.ProductRepo.GetCategorySlug(catSlug)
	if err != nil {
		return product_domain.Category{}, err
	}
	return dataCat, nil
}

// GetDetail implements product_domain.Business
func (pb ProductBusiness) GetDetail(detSlug string) (product_domain.Detail, error) {
	dataDetail, err := pb.ProductRepo.GetDetailSlug(detSlug)
	if err != nil {
		return product_domain.Detail{}, err
	}
	return dataDetail, nil
}

// GetProduct implements product_domain.Business
func (pb ProductBusiness) GetProduct(proSlug string) (product_domain.Product, error) {
	dataProduct, err := pb.ProductRepo.GetProductSlug(proSlug)
	if err != nil {
		return product_domain.Product{}, err
	}
	return dataProduct, nil
}

func NewProductBusiness(productRepo product_domain.Repository) product_domain.Business {
	return ProductBusiness{
		ProductRepo: productRepo,
	}
}

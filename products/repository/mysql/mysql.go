package product_mysql

import (
	product_domain "gozzafadillah/products/domain"

	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

// GetCategorySlug implements product_domain.Repository
func (pr ProductRepo) GetCategorySlug(catSlug string) (product_domain.Category, error) {
	record := Category{}
	err := pr.DB.Where("category_slug = ?", catSlug).First(&record).Error
	return ToDomainCategory(record), err
}

// GetDetailSlug implements product_domain.Repository
func (pr ProductRepo) GetDetailSlug(detSlug string) (product_domain.Detail, error) {
	record := Detail{}
	err := pr.DB.Where("detail_slug = ?", detSlug).First(&record).Error
	return ToDomainDetail(record), err
}

// GetProductSlug implements product_domain.Repository
func (pr ProductRepo) GetProductSlug(proSlug string) (product_domain.Product, error) {
	record := Product{}
	err := pr.DB.Where("product_slug = ?", proSlug).First(&record).Error
	return ToDomainProduct(record), err
}

// StoreCategory implements product_domain.Repository
func (pr ProductRepo) StoreCategory(domain product_domain.Category) error {
	err := pr.DB.Create(&domain).Error
	return err
}

// StoreDetail implements product_domain.Repository
func (pr ProductRepo) StoreDetail(domain product_domain.Detail) error {
	err := pr.DB.Create(&domain).Error
	return err
}

// StoreProduct implements product_domain.Repository
func (pr ProductRepo) StoreProduct(domain product_domain.Product) error {
	err := pr.DB.Create(&domain).Error
	return err
}

func NewProductRepo(db *gorm.DB) product_domain.Repository {
	return ProductRepo{
		DB: db,
	}
}

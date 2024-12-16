package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/rohan03122001/inventory-management-system/internal/models"
)

type ProductRepository struct {
	db *gorm.DB
}


//Initialise the repository
func NewProductRepository(db *gorm.DB) *ProductRepository{
	return &ProductRepository{
		db: db,
	}
}

//Creating the product
func (r *ProductRepository)Create(product *models.Product) error{
	return r.db.Create(product).Error
}

//Update
func (r *ProductRepository)Update(product *models.Product)(error){
	return r.db.Save(product).Error
}

//delete
func(r *ProductRepository)Delete(id uint) error{
	return r.db.Delete(&models.Product{}, id).Error
}

//---------------------------------------------------------------------------------

//Reading the product
func (r *ProductRepository)GetById(id uint) (*models.Product,error){
	var product models.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository)GetAll(page , pageSize int)([]models.Product, int, error){
	var products []models.Product
	var total int64

	//count total records
	if err:= r.db.Model(&models.Product{}).Count(&total).Error; err!=nil{
		return nil,0,err
	}


	//Pagination
	offset:= (page-1) *pageSize
	err := r.db.Offset(offset).Limit(pageSize).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return products, int(total), nil
}


//Get By SKU --> Custom Queries

func(r *ProductRepository)GetBySKU(sku string)(*models.Product, error){
	var product models.Product
	err:= r.db.Where("sku = ?",sku).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
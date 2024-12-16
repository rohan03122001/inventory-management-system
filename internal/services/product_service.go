package services

import (
	"fmt"

	"github.com/rohan03122001/inventory-management-system/internal/models"
	"github.com/rohan03122001/inventory-management-system/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

//Starting a new service 
func NewProductService(repo *repository.ProductRepository) *ProductService{
	return &ProductService{
		repo: repo,
	}
}


//Validating Product Creation for Product_Repository
func(s *ProductService)CreateProduct(product *models.Product)error{
	existing, err := s.repo.GetBySKU(product.SKU)
	if err == nil && existing != nil{
		return fmt.Errorf("Product Already Exists")
	}

	if product.Price <= 0{
		return fmt.Errorf("Product Price should be Greater Than Zero")
	}

	if product.Quantity < 0{
		return fmt.Errorf("Product Quantity Can't be Negative") 
	} 

	return s.repo.Create(product)
}


//Validating Get Product By ID for Product_Repository

func(s *ProductService)GetProduct(id uint)(*models.Product, error){
	product, err := s.repo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("Error Getting product %w",err)
	}

	return product,nil
}

//Validating Get ALL Product for Product_Repository

func(s *ProductService)ListProducts(page, pageSize int)([]models.Product, int, error){
	if page < 1{
		page = 1
	}
	if pageSize <1 || pageSize > 100{
		pageSize=10
	}

	return s.repo.GetAll(page, pageSize)
}

func(s *ProductService)UpdateProduct(id uint, product *models.Product) error{

	existing, err := s.repo.GetById(id)
	if err != nil{
		return fmt.Errorf("product not found: %w", err)
	}

	//For Our Custome GETBYSKU function
	if product.SKU != existing.SKU{
		skuProduct, err := s.repo.GetBySKU(product.SKU)
		if err == nil && skuProduct != nil {
            return fmt.Errorf("product with SKU %s already exists", product.SKU)
        }
	}

	existing.Name = product.Name
    existing.Description = product.Description
    existing.Price = product.Price
    existing.Quantity = product.Quantity
    existing.SKU = product.SKU

	return s.repo.Update(existing)
}



func(s *ProductService)DeleteProduct(id uint) error{
	_, err := s.repo.GetById(id)
	if err != nil {
		return fmt.Errorf("cant find the product %w", err)
	} 

	return s.repo.Delete(id)
}

func(s *ProductService)UpdateStock(id uint, quantity int) error{
	existing, err := s.repo.GetById(id)
	if err != nil {
		return fmt.Errorf("cant Find the Product %w",err)
	}

	if existing.Quantity+ quantity < 0{
		return fmt.Errorf("Insufficient Stock %w",err)
	}

	existing.Quantity+= quantity

	return s.repo.Update(existing)
}
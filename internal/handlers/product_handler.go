package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rohan03122001/inventory-management-system/internal/models"
	"github.com/rohan03122001/inventory-management-system/internal/services"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler{
	return &ProductHandler{
		service: service,
	}
}



func(h *ProductHandler)CreateProduct(c *gin.Context){
	var product models.Product

	//Bind json to above struct

	if err:= c.ShouldBindJSON(&product); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Payload" + err.Error(),
		})
		return
	}

	if err := h.service.CreateProduct(&product); err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed To create product" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated,product)
}



func(h *ProductHandler)GetProduct(c *gin.Context){
	id, err := strconv.ParseUint(c.Param("id"),10,32)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Invalid Product Id" + err.Error(),
		})
		return
	}
	
	product, err := h.service.GetProduct(uint(id))
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Product not found" +err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}


func(h *ProductHandler)ListProduct(c *gin.Context){

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "1"))

	product, total, err := h.service.ListProducts(page,pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch list" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"data":product,
		"total": total,
		"page": page,
		"page_size": pageSize,
	})
}

func(h *ProductHandler)UpdateProduct(c *gin.Context){

	//getting the param from user json
	id, err := strconv.ParseUint(c.Param("id"),10,32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Id"+err.Error(),
		})
		return
	}


	//binding json to struct
	var product models.Product
	if err:= c.ShouldBindJSON(&product); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Invalid Payload"+err.Error(),
		})
		return
	}

	//using the struct to update the database
	if err:= h.service.UpdateProduct(uint(id), &product); err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update the product"+err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,product)

}


func(h *ProductHandler)DeleteProduct(c *gin.Context){
	id, err := strconv.ParseUint(c.Param("id"),10 ,32)
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Product Id"+err.Error(),
		})
		return
	}

	if err:= h.service.DeleteProduct(uint(id)); err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete product"+err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}


func(h *ProductHandler)UpdateStock(c *gin.Context){
	//Get the Id from param
	id, err:= strconv.ParseUint(c.Param("id"),10,32)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"invalid product id"+err.Error(),
		})
		return
	}

	// new struct to pass into the upadate stock function
	var stockUpdate struct{
		Quantity int `json:"quantity" binding:"required"`
	}

	//bind the json to the stock update struct json BODY

	if err:= c.ShouldBindJSON(&stockUpdate); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Payload"+err.Error(), 
		})
		return
	}

	if err:= h.service.UpdateStock(uint(id), stockUpdate.Quantity); err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update Stock"+err.Error(), 
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
        "message": "Stock updated successfully",
    })
}
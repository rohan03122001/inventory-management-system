package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rohan03122001/inventory-management-system/internal/config"
	"github.com/rohan03122001/inventory-management-system/internal/handlers"
	"github.com/rohan03122001/inventory-management-system/internal/middleware"
	"github.com/rohan03122001/inventory-management-system/internal/repository"
	"github.com/rohan03122001/inventory-management-system/internal/services"
	"github.com/rohan03122001/inventory-management-system/pkg/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err!=nil{
		log.Fatal("Failed to Load Config",err)
	}

	//load Db
	db, err := database.NewDatabase(cfg)
	if err != nil {
        log.Fatal("Failed to initialize database:", err)
    }

	//setup middlewares
	//authMiddleware := middleware.NewAuthMiddleware(cfg.JWT_SECRET)
	logMiddleware := middleware.NewLogMiddleware()
	errorMiddleware :=middleware.NewErrorMiddleware()

	// setup Layers
	productRepo := repository.NewProductRepository(db.DB)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)


	// initialize Router

	router := gin.Default()


	//Apply Global Middlware

	router.Use(errorMiddleware.ErrorHandler())
	router.Use(logMiddleware.Logger())


	//API ROUTES

	api := router.Group("/api/v1")
	{
		api.GET("/health",func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"status": "OK",
			})
		})
		
		//protected Routes
		
		products := router.Group("/products")
		//products.Use(authMiddleware.Authenticate())
		{
			products.POST("",productHandler.CreateProduct)
			products.GET("",productHandler.ListProduct)
			products.GET("/:id", productHandler.GetProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
			products.PATCH("/:id/stock", productHandler.UpdateStock)
		}
	}

	if err := router.Run(":8000"); err!=nil{
		log.Fatal("Failed to start Server",err)
	}
}
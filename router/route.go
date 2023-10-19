package router

import (
	"github.com/AxelanO7/toko-yudha-backend-web-go/handler"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")

	// product
	product := api.Group("/product")
	// routes
	product.Get("/", handler.GetAllProducts)
	product.Get("/:id", handler.GetSingleProduct)
	product.Post("/", handler.CreateProduct)
	product.Put("/:id", handler.UpdateProduct)
	product.Delete("/:id", handler.DeleteProduct)

	// customer
	customer := api.Group("/customer")
	// routes
	customer.Get("/", handler.GetAllCustomers)
	customer.Get("/:id", handler.GetSingleCustomer)
	customer.Post("/", handler.CreateCustomer)
	customer.Put("/:id", handler.UpdateCustomer)
	customer.Delete("/:id", handler.DeleteCustomer)

	// user
	user := api.Group("/user")
	// routes
	user.Get("/", handler.GetAllUsers)
	user.Get("/:id", handler.GetSingleUser)
	user.Post("/", handler.CreateUser)
	user.Put("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUser)

	// sale
	sale := api.Group("/sale")
	// routes
	sale.Get("/", handler.GetAllSales)
	sale.Get("/:id", handler.GetSingleSale)
	sale.Post("/", handler.CreateSale)
	sale.Put("/:id", handler.UpdateSale)
	sale.Delete("/:id", handler.DeleteSale)

	// type product
	typeProduct := api.Group("/type-product")
	// routes
	typeProduct.Get("/", handler.GetAllTypeProducts)
	typeProduct.Get("/:id", handler.GetSingleTypeProduct)
	typeProduct.Post("/", handler.CreateTypeProduct)
	typeProduct.Put("/:id", handler.UpdateTypeProduct)
	typeProduct.Delete("/:id", handler.DeleteTypeProduct)

	// cart
	cart := api.Group("/cart")
	// routes
	cart.Get("/", handler.GetAllCarts)
	cart.Get("/:id", handler.GetSingleCart)
	cart.Post("/", handler.CreateCart)
	cart.Put("/:id", handler.UpdateCart)
	cart.Delete("/:id", handler.DeleteCart)
}

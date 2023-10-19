package handler

import (
	"fmt"

	"github.com/AxelanO7/toko-yudha-backend-web-go/database"
	"github.com/AxelanO7/toko-yudha-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find product by id
func FindProductByID(id string, product *model.Product) error {
	db := database.DB.Db
	// find single product in the database by id
	db.Find(&product, "id = ?", id)
	// if no product found, return an error
	if product.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a product
func CreateProduct(c *fiber.Ctx) error {
	db := database.DB.Db
	product := new(model.Product)
	typeProduct := new(model.TypeProduct)
	// store the body in the product and return error if encountereds
	if err := c.BodyParser(product); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find typeProduct in the database by id
	if err := FindTypeProductByID(fmt.Sprint(product.TypeProductID), typeProduct); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "TypeProduct not found"})
	}
	// assign typeProduct to product
	product.TypeProduct = *typeProduct
	// create product
	if err := db.Create(product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create product", "data": err})
	}
	// return the created product
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Product has created", "data": product})
}

// get all products from db
func GetAllProducts(c *fiber.Ctx) error {
	db := database.DB.Db
	products := []model.Product{}
	// find all products in the database
	if err := db.Find(&products).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Products not found", "data": nil})
	}
	// if no product found, return an error
	if len(products) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Products not found", "data": nil})
	}
	responseProducts := []model.Product{}
	for _, product := range products {
		typeProduct := new(model.TypeProduct)
		// find typeProduct in the database by id
		if err := FindTypeProductByID(fmt.Sprint(product.TypeProductID), typeProduct); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "TypeProduct not found"})
		}
		// assign typeProduct to product
		product.TypeProduct = *typeProduct
		responseProducts = append(responseProducts, product)
	}
	// return products
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Products Found", "data": responseProducts})
}

// get single product from db
func GetSingleProduct(c *fiber.Ctx) error {
	product := new(model.Product)
	typeProduct := new(model.TypeProduct)
	// get id params
	id := c.Params("id")
	// find single product in the database by id
	if err := FindProductByID(id, product); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Product not found"})
	}
	// find typeProduct in the database by id
	if err := FindTypeProductByID(fmt.Sprint(product.TypeProduct), typeProduct); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "TypeProduct not found"})
	}
	// assign typeProduct to product
	product.TypeProduct = *typeProduct
	// return product
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Product Found", "data": product})
}

// update a product in db
func UpdateProduct(c *fiber.Ctx) error {
	db := database.DB.Db
	product := new(model.Product)
	typeProduct := new(model.TypeProduct)
	// get id params
	id := c.Params("id")
	// find single product in the database by id
	if err := FindProductByID(id, product); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Product not found"})
	}
	// store the body in the product and return error if encountereds
	if err := c.BodyParser(product); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find typeProduct in the database by id
	if err := FindTypeProductByID(fmt.Sprint(product.TypeProduct), typeProduct); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "TypeProduct not found"})
	}
	// assign typeProduct to product
	product.TypeProduct = *typeProduct
	// update product
	if err := db.Save(product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update product", "data": err})
	}
	// return the updated product
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Product has updated", "data": product})
}

// delete a product in db
func DeleteProduct(c *fiber.Ctx) error {
	db := database.DB.Db
	product := new(model.Product)
	// get id params
	id := c.Params("id")
	// find single product in the database by id
	if err := FindProductByID(id, product); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Product not found"})
	}
	// delete product
	if err := db.Delete(product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete product", "data": err})
	}
	// return deleted product
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Product deleted"})
}

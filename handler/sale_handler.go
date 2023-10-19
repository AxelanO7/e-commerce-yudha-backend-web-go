package handler

import (
	"fmt"

	"github.com/AxelanO7/toko-yudha-backend-web-go/database"
	"github.com/AxelanO7/toko-yudha-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find sale by id
func FindSaleByID(id string, Sale *model.Sale) error {
	db := database.DB.Db
	// find single sale in the database by id
	db.Find(&Sale, "id = ?", id)
	// if no sale found, return an error
	if Sale.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a sale
func CreateSale(c *fiber.Ctx) error {
	db := database.DB.Db
	sale := new(model.Sale)
	user := new(model.User)
	product := new(model.Product)
	customer := new(model.Customer)
	// store the body in the sale and return error if encountereds
	if err := c.BodyParser(sale); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your sale", "data": err})
	}
	// find user in the database by id
	if err := FindUserByID(fmt.Sprint(sale.UserID), user); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}
	// find product in the database by id
	if err := FindProductByID(fmt.Sprint(sale.ProductID), product); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Product not found"})
	}
	// find customer in the database by id
	if err := FindCustomerByID(fmt.Sprint(sale.CustomerID), customer); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Customer not found"})
	}
	// assign user to sale
	sale.User = *user
	// assign product to sale
	sale.Product = *product
	// assign customer to sale
	sale.Customer = *customer
	// create sale
	if err := db.Create(sale).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create sale", "data": err})
	}
	// return the created sale
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Sale has created", "data": sale})
}

// get all sales from db
func GetAllSales(c *fiber.Ctx) error {
	db := database.DB.Db
	sales := []model.Sale{}
	// find all sales in the database
	if err := db.Find(&sales).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Sales not found", "data": nil})
	}
	// if no sale found, return an error
	if len(sales) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Sales not found", "data": nil})
	}
	responseSales := []model.Sale{}
	for _, sale := range sales {
		user := new(model.User)
		product := new(model.Product)
		customer := new(model.Customer)
		// find user in the database by id
		if err := FindUserByID(fmt.Sprint(sale.UserID), user); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found"})
		}
		// find product in the database by id
		if err := FindProductByID(fmt.Sprint(sale.ProductID), product); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Product not found"})
		}
		// find customer in the database by id
		if err := FindCustomerByID(fmt.Sprint(sale.CustomerID), customer); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Customer not found"})
		}
		// assign user to sale
		sale.User = *user
		// assign product to sale
		sale.Product = *product
		// assign customer to sale
		sale.Customer = *customer
		responseSales = append(responseSales, sale)
	}
	// return sales
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Sales Found", "data": responseSales})
}

// get single sale from db
func GetSingleSale(c *fiber.Ctx) error {
	sale := new(model.Sale)
	user := new(model.User)
	product := new(model.Product)
	customer := new(model.Customer)
	// get id params
	id := c.Params("id")
	// find sale in the database by id
	if err := FindSaleByID(id, sale); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Sale not found"})
	}
	// find user in the database by id
	if err := FindUserByID(fmt.Sprint(sale.UserID), user); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}
	// find product in the database by id
	if err := FindProductByID(fmt.Sprint(sale.ProductID), product); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Product not found"})
	}
	// find customer in the database by id
	if err := FindCustomerByID(fmt.Sprint(sale.CustomerID), customer); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Customer not found"})
	}
	// assign user to sale
	sale.User = *user
	// assign product to sale
	sale.Product = *product
	// assign customer to sale
	sale.Customer = *customer
	// return sale
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Sale found", "data": sale})
}

// update a sale in db
func UpdateSale(c *fiber.Ctx) error {
	db := database.DB.Db
	sale := new(model.Sale)
	user := new(model.User)
	product := new(model.Product)
	customer := new(model.Customer)
	// get id params
	id := c.Params("id")
	// find sale in the database by id
	if err := FindSaleByID(id, sale); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Sale not found"})
	}
	// store the body in the sale and return error if encountereds
	if err := c.BodyParser(sale); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your sale", "data": err})
	}
	// find user in the database by id
	if err := FindUserByID(fmt.Sprint(sale.UserID), user); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}
	// find product in the database by id
	if err := FindProductByID(fmt.Sprint(sale.ProductID), product); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Product not found"})
	}
	// find customer in the database by id
	if err := FindCustomerByID(fmt.Sprint(sale.CustomerID), customer); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Customer not found"})
	}
	// assign user to sale
	sale.User = *user
	// assign product to sale
	sale.Product = *product
	// assign customer to sale
	sale.Customer = *customer
	// update sale
	if err := db.Save(sale).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update sale", "data": err})
	}
	// return the updated sale
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Sale has updated", "data": sale})
}

// delete a sale in db
func DeleteSale(c *fiber.Ctx) error {
	db := database.DB.Db
	sale := new(model.Sale)
	// get id params
	id := c.Params("id")
	// find sale in the database by id
	if err := FindSaleByID(id, sale); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Sale not found"})
	}
	// delete sale
	if err := db.Delete(&sale).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete sale", "data": err})
	}
	// return the deleted sale
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Sale has deleted", "data": sale})
}

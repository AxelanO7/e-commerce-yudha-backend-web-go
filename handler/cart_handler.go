package handler

import (
	"fmt"

	"github.com/AxelanO7/toko-yudha-backend-web-go/database"
	"github.com/AxelanO7/toko-yudha-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find cart by id
func FindCartByID(id string, Cart *model.Cart) error {
	db := database.DB.Db
	// find single cart in the database by id
	db.Find(&Cart, "id = ?", id)
	// if no cart found, return an error
	if Cart.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a cart
func CreateCart(c *fiber.Ctx) error {
	db := database.DB.Db
	cart := new(model.Cart)
	sale := new(model.Sale)
	product := new(model.Product)
	// store the body in the cart and return error if encountereds
	if err := c.BodyParser(cart); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your cart", "data": err})
	}
	// find sale in the database by id
	if err := FindSaleByID(fmt.Sprint(cart.SaleId), sale); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Sale not found"})
	}
	// find product in the database by id
	if err := FindProductByID(fmt.Sprint(cart.ProductId), product); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Product not found"})
	}
	// assign sale to cart
	cart.Sale = *sale
	// assign product to cart
	cart.Product = *product
	// create cart
	if err := db.Create(cart).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create cart", "data": err})
	}
	// return the created cart
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Cart has created", "data": cart})
}

// get all carts from db
func GetAllCarts(c *fiber.Ctx) error {
	db := database.DB.Db
	carts := []model.Cart{}
	// find all carts in the database
	if err := db.Find(&carts).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Carts not found", "data": nil})
	}
	// if no cart found, return an error
	if len(carts) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Carts not found", "data": nil})
	}
	responseCarts := []model.Cart{}
	for _, cart := range carts {
		sale := new(model.Sale)
		product := new(model.Product)
		// convert id to string
		idSale := fmt.Sprint(cart.SaleId)
		idProduct := fmt.Sprint(cart.ProductId)
		// find account in the database by id
		if err := FindSaleByID(idSale, sale); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
		}
		// find cart in the database by id
		if err := FindProductByID(idProduct, product); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Cart not found"})
		}
		// assign sale to cart
		cart.Sale = *sale
		// assign product to cart
		cart.Product = *product
		responseCarts = append(responseCarts, cart)
	}
	// return carts
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Carts Found", "data": responseCarts})
}

// get single cart from db
func GetSingleCart(c *fiber.Ctx) error {
	cart := new(model.Cart)
	sale := new(model.Sale)
	product := new(model.Product)
	// get id params
	id := c.Params("id")
	// find single cart in the database by id
	if err := FindCartByID(id, cart); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Cart not found"})
	}
	// find sale in the database by id
	if err := FindSaleByID(fmt.Sprint(cart.SaleId), sale); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Sale not found"})
	}
	// find product in the database by id
	if err := FindProductByID(fmt.Sprint(cart.ProductId), product); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Product not found"})
	}
	// assign sale to cart
	cart.Sale = *sale
	// assign product to cart
	cart.Product = *product
	// return cart
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Cart Found", "data": cart})
}

// update a cart in db
func UpdateCart(c *fiber.Ctx) error {
	db := database.DB.Db
	cart := new(model.Cart)
	sale := new(model.Sale)
	product := new(model.Product)
	// get id params
	id := c.Params("id")
	// find single cart in the database by id
	if err := FindCartByID(id, cart); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Cart not found"})
	}
	// store the body in the cart and return error if encountereds
	if err := c.BodyParser(cart); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your cart", "data": err})
	}
	// find sale in the database by id
	if err := FindSaleByID(fmt.Sprint(cart.SaleId), sale); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Sale not found"})
	}
	// find product in the database by id
	if err := FindProductByID(fmt.Sprint(cart.ProductId), product); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Product not found"})
	}
	// assign sale to cart
	cart.Sale = *sale
	// assign product to cart
	cart.Product = *product
	// update cart
	if err := db.Save(cart).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update cart", "data": err})
	}
	// return the updated cart
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Cart has updated", "data": cart})
}

// delete a cart in db
func DeleteCart(c *fiber.Ctx) error {
	db := database.DB.Db
	Cart := new(model.Cart)
	// get id params
	id := c.Params("id")
	// find single cart in the database by id
	if err := FindCartByID(id, Cart); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Cart not found"})
	}
	// delete cart
	if err := db.Delete(Cart).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete cart", "data": err})
	}
	// return the deleted cart
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Cart has deleted", "data": Cart})
}

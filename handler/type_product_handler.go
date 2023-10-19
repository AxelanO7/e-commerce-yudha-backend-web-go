package handler

import (
	"github.com/AxelanO7/toko-yudha-backend-web-go/database"
	"github.com/AxelanO7/toko-yudha-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find typeProduct by id
func FindTypeProductByID(id string, typeProduct *model.TypeProduct) error {
	db := database.DB.Db
	// find single typeProduct in the database by id
	db.Find(&typeProduct, "id = ?", id)
	// if no typeProduct found, return an error
	if typeProduct.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a typeProduct
func CreateTypeProduct(c *fiber.Ctx) error {
	db := database.DB.Db
	typeProduct := new(model.TypeProduct)
	// store the body in the typeProduct and return error if encountered
	if err := c.BodyParser(typeProduct); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// create typeProduct
	if err := db.Create(typeProduct).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create typeProduct", "data": err})
	}
	// return the created typeProduct
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "TypeProduct has created", "data": typeProduct})
}

// get all TypeProducts from db
func GetAllTypeProducts(c *fiber.Ctx) error {
	db := database.DB.Db
	typeProducts := []model.TypeProduct{}
	if err := db.Find(&typeProducts).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get typeProducts", "data": err})
	}
	// if no typeProduct found, return an error
	if len(typeProducts) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "TypeProducts not found", "data": nil})
	}
	// return typeProducts
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "TypeProducts Found", "data": typeProducts})
}

// GetSingleTypeProduct from db
func GetSingleTypeProduct(c *fiber.Ctx) error {
	typeProduct := new(model.TypeProduct)
	// get id params
	id := c.Params("id")
	// find single typeProduct in the database by id
	if err := FindTypeProductByID(id, typeProduct); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "TypeProduct not found"})
	}
	// return typeProduct
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "TypeProduct Found", "data": typeProduct})
}

// update a typeProduct in db
func UpdateTypeProduct(c *fiber.Ctx) error {
	db := database.DB.Db
	typeProduct := new(model.TypeProduct)
	// get id params
	id := c.Params("id")
	// find single typeProduct in the database by id
	if err := FindTypeProductByID(id, typeProduct); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "TypeProduct not found"})
	}
	// store the body in the typeProduct and return error if encountered
	if err := c.BodyParser(typeProduct); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// update typeProduct
	if err := db.Save(typeProduct).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update typeProduct", "data": err})
	}
	// return the updated typeProduct
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "typeProducts Found", "data": typeProduct})
}

// delete a typeProduct in db
func DeleteTypeProduct(c *fiber.Ctx) error {
	db := database.DB.Db
	typeProduct := new(model.TypeProduct)
	// get id params
	id := c.Params("id")
	// find single typeProduct in the database by id
	if err := FindTypeProductByID(id, typeProduct); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "TypeProduct not found"})
	}
	// delete typeProduct
	if err := db.Delete(typeProduct, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete typeProduct", "data": err})
	}
	// return success message
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "TypeProduct deleted"})
}

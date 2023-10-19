package handler

import (
	"fmt"

	"github.com/AxelanO7/toko-yudha-backend-web-go/database"
	"github.com/AxelanO7/toko-yudha-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find customer by id
func FindCustomerByID(id string, customer *model.Customer) error {
	db := database.DB.Db
	// find single customer in the database by id
	db.Find(&customer, "id = ?", id)
	// if no customer found, return an error
	if customer.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a customer
func CreateCustomer(c *fiber.Ctx) error {
	db := database.DB.Db
	customer := new(model.Customer)
	user := new(model.User)
	// store the body in the customer and return error if encountereds
	if err := c.BodyParser(customer); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find user in the database by id
	if err := FindUserByID(fmt.Sprint(customer.UserID), user); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}
	// assign user to customer
	customer.User = *user
	// create customer
	if err := db.Create(customer).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create customer", "data": err})
	}
	// return the created customer
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Customer has created", "data": customer})
}

// get all customers from db
func GetAllCustomers(c *fiber.Ctx) error {
	db := database.DB.Db
	customers := []model.Customer{}
	// find all customers in the database
	if err := db.Find(&customers).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Customers not found", "data": nil})
	}
	// if no customer found, return an error
	if len(customers) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Customers not found", "data": nil})
	}
	responseCustomers := []model.Customer{}
	for _, customer := range customers {
		user := new(model.User)
		// find user in the database by id
		if err := FindUserByID(fmt.Sprint(customer.UserID), user); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found"})
		}
		// assign user to customer
		customer.User = *user
		responseCustomers = append(responseCustomers, customer)
	}
	// return customers
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Customers Found", "data": responseCustomers})
}

// get single customer from db
func GetSingleCustomer(c *fiber.Ctx) error {
	customer := new(model.Customer)
	user := new(model.User)
	// get id params
	id := c.Params("id")
	// find single customer in the database by id
	if err := FindCustomerByID(id, customer); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Customer not found"})
	}
	// find user in the database by id
	if err := FindUserByID(fmt.Sprint(customer.User), user); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}
	// assign user to customer
	customer.User = *user
	// return customer
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Customer Found", "data": customer})
}

// update a customer in db
func UpdateCustomer(c *fiber.Ctx) error {
	db := database.DB.Db
	customer := new(model.Customer)
	user := new(model.User)
	// get id params
	id := c.Params("id")
	// find single customer in the database by id
	if err := FindCustomerByID(id, customer); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Customer not found"})
	}
	// store the body in the customer and return error if encountereds
	if err := c.BodyParser(customer); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find user in the database by id
	if err := FindUserByID(fmt.Sprint(customer.User), user); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}
	// assign user to customer
	customer.User = *user
	// update customer
	if err := db.Save(customer).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update customer", "data": err})
	}
	// return the updated customer
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Customer has updated", "data": customer})
}

// delete a customer in db
func DeleteCustomer(c *fiber.Ctx) error {
	db := database.DB.Db
	customer := new(model.Customer)
	// get id params
	id := c.Params("id")
	// find single customer in the database by id
	if err := FindCustomerByID(id, customer); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Customer not found"})
	}
	// delete customer
	if err := db.Delete(customer).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete customer", "data": err})
	}
	// return deleted customer
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Customer deleted"})
}

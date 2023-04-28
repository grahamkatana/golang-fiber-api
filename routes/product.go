package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/grahamkatana/fiber-api/database"
	"github.com/grahamkatana/fiber-api/models"
)

// serializer
type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

// end serializer

func CreateResponseProduct(productModel models.Product) Product {
	return Product{ID: productModel.ID, Name: productModel.Name, SerialNumber: productModel.SerialNumber}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())

	}
	database.Database.Db.Create(&product)
	response := CreateResponseProduct(product)
	return c.Status(201).JSON(response)

}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.Database.Db.Find(&products)
	response_all := []Product{}
	for _, product := range products {
		response := CreateResponseProduct(product)
		response_all = append(response_all, response)

	}
	return c.Status(200).JSON(response_all)

}
func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("no such record exists")
	}
	return nil

}
func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure that the id is an integer ")
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())

	}
	response := CreateResponseProduct(product)
	return c.Status(200).JSON(response)

}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure that the id is an integer ")
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())

	}

	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var updatedData UpdateProduct
	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(500).JSON(err.Error())

	}
	product.Name = updatedData.Name
	product.SerialNumber = updatedData.SerialNumber

	database.Database.Db.Save(&product)
	response := CreateResponseProduct(product)
	return c.Status(200).JSON(response)

}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure that the id is an integer ")
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())

	}
	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Product was deleted")

}

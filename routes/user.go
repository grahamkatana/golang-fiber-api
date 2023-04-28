package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/grahamkatana/fiber-api/database"
	"github.com/grahamkatana/fiber-api/models"
)

// serializer
type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// end serializer

func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())

	}
	database.Database.Db.Create(&user)
	response := CreateResponseUser(user)
	return c.Status(201).JSON(response)

}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Database.Db.Find(&users)
	response_all := []User{}
	for _, user := range users {
		response := CreateResponseUser(user)
		response_all = append(response_all, response)

	}
	return c.Status(200).JSON(response_all)

}
func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("No such record exists")
	}
	return nil

}
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that the id is an integer ")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())

	}
	response := CreateResponseUser(user)
	return c.Status(200).JSON(response)

}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that the id is an integer ")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())

	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updatedData UpdateUser
	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(500).JSON(err.Error())

	}
	user.FirstName = updatedData.FirstName
	user.LastName = updatedData.LastName

	database.Database.Db.Save(&user)
	response := CreateResponseUser(user)
	return c.Status(200).JSON(response)

}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that the id is an integer ")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())

	}
	if err :=database.Database.Db.Delete(&user).Error; err!=nil{
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("User was deleted")


}

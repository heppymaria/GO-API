package controllers

import (
	"go-api/config"
	"go-api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Get users (with optional name filter)
func GetUsers(c *fiber.Ctx) error {
	name := c.Query("name") // ambil query param ?name=

	var users []models.User
	query := config.DB

	if name != "" {
		// pakai LIKE untuk partial match
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	result := query.Find(&users)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.JSON(users)
}
func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var user models.User
	result := config.DB.First(&user, userID)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(user)
}

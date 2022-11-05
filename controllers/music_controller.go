package controllers

import (
	"meditation/database"
	"meditation/models"

	"github.com/gofiber/fiber/v2"
)

func GetMusic(c *fiber.Ctx) error {
	data := []models.Music{}

	err := database.DBQuery.Find(&data).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error",
		})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Successfully get data music",
		"data":    data,
	})
}

func GetMusicCategory(c *fiber.Ctx) error {
	category := c.Params("category")

	data := []models.Music{}

	err := database.DBQuery.Where("category = ?", category).Find(&data).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error",
		})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Successfully get data music",
		"data":    data,
	})
}

package handlers

import (
	"NormsServer/database"
	"NormsServer/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Получение информации о станках
// @Description Возвращает информацию о станках предприятия
// @Tags cutting_info
// @Accept json
// @Produce json
// @Success 200 {array} models.Machine
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /machines [get]
func GetMachines(c *fiber.Ctx) error {
	var machines []models.Machine
	err := database.DB.Select(&machines, "SELECT machine_id, machine_name FROM machines")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(machines)
}

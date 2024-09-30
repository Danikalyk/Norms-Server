package handlers

import (
	"NormsServer/database"
	"NormsServer/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Получение информации о материале
// @Description Возвращает информацию о материале по заданному станку
// @Tags cutting_info
// @Accept json
// @Produce json
// @Param machine_name query string true "Название станка"
// @Success 200 {array} models.Materials
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /materials [get]
func GetMaterials(c *fiber.Ctx) error {
	machine_name := c.Query("machine_name")

	var materials []models.Materials
	query := `SELECT DISTINCT materials.material_name
		FROM cutting_info
		INNER JOIN materials ON cutting_info.material = materials.material_id
		INNER JOIN machines ON cutting_info.machine = machines.machine_id
		WHERE machines.machine_name = $1 and materials.welding_check = 'false'`
	err := database.DB.Select(&materials, query, machine_name)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(materials)
}

// @Summary Получение информации о материале используемом для сварочных работ
// @Description Возвращает информацию о материале
// @Tags welding_info
// @Accept json
// @Produce json
// @Success 200 {array} models.WeldingMaterials
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /welding_materials [get]
func GetWeldingMaterials(c *fiber.Ctx) error {

	var materials []models.WeldingMaterials
	query := `SELECT DISTINCT material_name
		      FROM materials where welding_check = 'true'`
	err := database.DB.Select(&materials, query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(materials)
}

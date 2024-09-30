package handlers

import (
	"NormsServer/database"
	"NormsServer/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Получение информации о типе сварки
// @Description Возвращает информацию о типах сварки для заданного материала
// @Tags welding_info
// @Accept json
// @Produce json
// @Param material_name query string true "Название материала"
// @Success 200 {array} models.WeldingTypes
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /welding_types [get]
func GetWeldingType(c *fiber.Ctx) error {
	material_name := c.Query("material_name")

	var weldingTypes []models.WeldingTypes
	query := `select distinct welding_types.welding_type_name
			  from tnsh inner join welding_types on tnsh.welding_type_fk = welding_types.welding_type_id
			  inner join materials on tnsh.material_fk = materials.material_id
			  where materials.material_name = $1 and materials.welding_check = 'true'`

	err := database.DB.Select(&weldingTypes, query, material_name)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(weldingTypes)
}

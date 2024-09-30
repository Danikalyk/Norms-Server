package handlers

import (
	"NormsServer/database"
	"NormsServer/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Получение информации о катетах(толщине) металла
// @Description Возвращает информацию о типах сварки для заданного материала и типа сварки
// @Tags welding_info
// @Accept json
// @Produce json
// @Param material_name query string true "Название материала"
// @Param welding_type_name query string true "Название типа сварки"
// @Success 200 {array} models.Katets
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /katets [get]
func GetKatet(c *fiber.Ctx) error {
	material_name := c.Query("material_name")
	welding_type_name := c.Query("welding_type_name")

	var katets []models.Katets
	query := `select distinct katets.katet_value from tnsh
			  inner join katets on tnsh.katet_fk = katets.katet_id
			  inner join materials on tnsh.material_fk = materials.material_id
			  inner join welding_types on tnsh.welding_type_fk = welding_types.welding_type_id
			  where materials.material_name = $1 and welding_types.welding_type_name = $2 and materials.welding_check = 'true'`

	err := database.DB.Select(&katets, query, material_name, welding_type_name)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(katets)
}

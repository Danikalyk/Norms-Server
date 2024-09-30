package handlers

import (
	"NormsServer/database"
	"NormsServer/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Получение информации о типе шва
// @Description Возвращает информацию о типе шва по заданным материалу, типу сварки и катету(толщине)
// @Tags welding_info
// @Accept json
// @Produce json
// @Param material_name query string true "Название материала"
// @Param welding_type_name query string true "Название типа сварки"
// @Param katet_value query float64 true "Катет(толщина) материала"
// @Success 200 {array} models.SeamTypes
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /seam_type [get]
func GetSeamType(c *fiber.Ctx) error {
	material_name := c.Query("material_name")
	welding_type_name := c.Query("welding_type_name")
	katet_value := c.Query("katet_value")

	var seamTypes []models.SeamTypes

	query := `select distinct seam_types.seam_type_name from tnsh
			  inner join seam_types on tnsh.seam_type_fk = seam_types.seam_type_id
			  inner join materials on tnsh.material_fk = materials.material_id
			  inner join welding_types on tnsh.welding_type_fk = welding_types.welding_type_id
			  inner join katets on tnsh.katet_fk = katets.katet_id
			  where materials.material_name = $1 and welding_types.welding_type_name = $2 and katets.katet_value = $3 and materials.welding_check = 'true'`

	err := database.DB.Select(&seamTypes, query, material_name, welding_type_name, katet_value)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(seamTypes)
}

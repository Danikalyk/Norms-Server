package handlers

import (
	"NormsServer/database"
	"NormsServer/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Получение информации о площади сечения шва
// @Description Возвращает информацию о площади сечения шва на основе материала, типа сварки, катета(толщины), типа шва
// @Tags welding_info
// @Accept json
// @Produce json
// @Param material_name query string true "Название материала"
// @Param welding_type_name query string true "Название типа сварки"
// @Param katet_value query float64 true "Значение катета(толщины)"
// @Param seam_type_name query string true "Название типа шва"
// @Success 200 {array} models.CrossSecArea
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /cross_sec_area [get]
func GetCrossSecArea(c *fiber.Ctx) error {
	material_name := c.Query("material_name")
	welding_type_name := c.Query("welding_type_name")
	katet_value := c.Query("katet_value")
	seam_type_name := c.Query("seam_type_name")

	var CrossSecArea []models.CrossSecArea
	query := `select cross_sec_areas.area_value from tnsh
			  inner join cross_sec_areas on tnsh.cross_sec_area_fk = cross_sec_areas.area_id
			  inner join materials on tnsh.material_fk = materials.material_id
			  inner join welding_types on tnsh.welding_type_fk = welding_types.welding_type_id
			  inner join katets on tnsh.katet_fk = katets.katet_id
			  inner join seam_types on tnsh.seam_type_fk = seam_types.seam_type_id
			  where materials.material_name = $1 and welding_types.welding_type_name = $2 and katets.katet_value = $3 and seam_types.seam_type_name = $4 and materials.welding_check = 'true'`
	err := database.DB.Select(&CrossSecArea, query, material_name, welding_type_name, katet_value, seam_type_name)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(CrossSecArea)
}

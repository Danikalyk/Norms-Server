package handlers

import (
	"NormsServer/database"
	"NormsServer/models"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// @Summary Получение времени затрачиваемого на одну штуку
// @Description Возвращает информацию о времени затрачиваемом на обработку одной единицы изделия на основании материала, типа сварки, катета(толщины), типа шва, площади сечения и диаметра проволоки
// @Tags welding_info
// @Accept json
// @Produce json
// @Param material_name query string true "Название материала"
// @Param welding_type_name query string true "Название типа сварки"
// @Param katet_value query float64 true "Значение катета(толщины)"
// @Param seam_type_name query string true "Название типа шва"
// @Param area_value query float64 true "Площадь сечения"
// @Param diameter_value query float64 true "Диаметер проволоки"
// @Success 200 {array} models.Tnsh
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /tnsh [get]
func GetTNSH(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	material_name := c.Query("material_name")
	welding_type_name := c.Query("welding_type_name")
	katet_value := c.Query("katet_value")
	seam_type_name := c.Query("seam_type_name")
	cross_sec_area := c.Query("area_value")
	diameter_value := c.Query("diameter_value")

	if material_name == "" || welding_type_name == "" || katet_value == "" || seam_type_name == "" || cross_sec_area == "" || diameter_value == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Все параметры запроса обязательны",
		})
	}

	var tnsh []models.Tnsh

	query := `select tnsh.tnsh_value from tnsh
			  inner join wire_diameter on tnsh.wire_diameter_fk = wire_diameter.record_id
			  inner join cross_sec_areas on tnsh.cross_sec_area_fk = cross_sec_areas.area_id
			  inner join materials on tnsh.material_fk = materials.material_id
			  inner join welding_types on tnsh.welding_type_fk = welding_types.welding_type_id
			  inner join katets on tnsh.katet_fk = katets.katet_id
			  inner join seam_types on tnsh.seam_type_fk = seam_types.seam_type_id
			  where materials.material_name = $1 and welding_types.welding_type_name = $2 and katets.katet_value = $3 
			  and seam_types.seam_type_name = $4 and cross_sec_areas.area_value = $5 and wire_diameter.diameter_value = $6 and materials.welding_check = 'true'`
	err := database.DB.Select(&tnsh, query, material_name, welding_type_name, katet_value, seam_type_name, cross_sec_area, diameter_value)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		if ctx.Err() == context.DeadlineExceeded {
			return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
				"error": "Запрос к базе данных превысил время ожидания",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось получить значение tnsh",
		})
	}

	return c.JSON(tnsh)
}

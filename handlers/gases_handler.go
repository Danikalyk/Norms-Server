package handlers

import (
	"NormsServer/database"
	"NormsServer/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Получение информации о режущем газе
// @Description Возвращает информацию о режущем газе для заданной машины, материала
// @Tags cutting_info
// @Accept json
// @Produce json
// @Param machine_name query string true "Название машины"
// @Param material_name query string true "Название материала"
// @Param tickness query float64 true "Толщина"
// @Success 200 {array} models.Gases
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /gases [get]
func GetGases(c *fiber.Ctx) error {
	machine_name := c.Query("machine_name")
	material_name := c.Query("material_name")
	tickness := c.Query("tickness")

	var gases []models.Gases
	query := `select distinct gases.gas_id, gases.gas_name from cutting_info
			  inner join gases on cutting_info.gas = gases.gas_id
			  inner join machines on cutting_info.machine = machines.machine_id
			  inner join materials on cutting_info.material = materials.material_id
			  where machines.machine_name = $1 and materials.material_name = $2 and cutting_info.tickness = $3`
	err := database.DB.Select(&gases, query, machine_name, material_name, tickness)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(gases)
}

package handlers

import (
	"NormsServer/database"
	"NormsServer/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Получение информации о резке
// @Description Возвращает информацию о резке для заданной машины, материала, толщины и газа.
// @Tags cutting_info
// @Accept json
// @Produce json
// @Param machine_name query string true "Название машины"
// @Param material_name query string true "Название материала"
// @Param tickness query string true "Толщина материала"
// @Param gas query string true "Название газа"
// @Success 200 {array} models.CUTTING_INFO
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /cutting_info [get]
func GetCuttingInfo(c *fiber.Ctx) error {
	machine_name := c.Query("machine_name")
	material_name := c.Query("material_name")
	tickness := c.Query("tickness")
	gas := c.Query("gas")

	var cuttingInfo []models.CUTTING_INFO

	query := `select cutting_info.record_id, cutting_info.mincuttingspeed, cutting_info.avecuttingspeed, cutting_info.maxcuttingspeed, cutting_info.insertiontime, cutting_info.data_updated
			  from cutting_info inner join machines on cutting_info.machine = machines.machine_id
			  inner join materials on cutting_info.material = materials.material_id
			  inner join gases on cutting_info.gas = gases.gas_id
			  where machines.machine_name = $1 and materials.material_name = $2 and cutting_info.tickness = $3 and gases.gas_name = $4`
	err := database.DB.Select(&cuttingInfo, query, machine_name, material_name, tickness, gas)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(cuttingInfo)
}

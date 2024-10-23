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

type InsertSeamTypeRequest struct {
	SeamTypeName string `json:"seam_type_name" validate:"required"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SeamTypeResponse struct {
	SeamTypeName string `json:"seam_type_name"`
}

// @Summary Вставить новый тип шва
// @Description Создает новый тип шва с указанным именем.
// @Tags welding_info
// @Accept  json
// @Produce  json
// @Param   seam_type_name body InsertSeamTypeRequest true "Название типа шва"
// @Success 201 {object} SeamTypeResponse "Тип шва успешно создан"
// @Failure 400 {object} ErrorResponse "Некорректный формат данных или отсутствуют обязательные поля"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /create_seam_type [post]
func InsertSeamType(c *fiber.Ctx) error {
	var req InsertSeamTypeRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Некорректный формат данных",
		})
	}

	if req.SeamTypeName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Поле seam_type_name обязательно и должно быть заполнено",
		})
	}

	query := `INSERT INTO seam_types (seam_type_name) VALUES ($1) RETURNING seam_type_id`

	var newID int

	err := database.DB.QueryRow(query, req.SeamTypeName).Scan(&newID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Не удалось вставить данные",
		})
	}

	newSeamType := SeamTypeResponse{
		SeamTypeName: req.SeamTypeName,
	}

	return c.Status(fiber.StatusCreated).JSON(newSeamType)
}

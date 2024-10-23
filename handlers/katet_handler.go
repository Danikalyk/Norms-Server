package handlers

import (
	"NormsServer/database"
	"NormsServer/models"
	"fmt"

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

type InsertKatetRequest struct {
	KatetValue string `json:"katet_value" validate:"required"`
}

type KatetErrorResponse struct {
	Error string `json:"error"`
}

type KatetResponse struct {
	KatetValue string `json:"katet_value"`
}

// @Summary Вставить новый катет(толщину)
// @Description Добавляет в базу данных новый катет(толщину)
// @Tags welding_info
// @Accept  json
// @Produce  json
// @Param   katet_value body InsertKatetRequest true "Значение катета"
// @Success 201 {object} SeamTypeResponse "Катет успешно создан"
// @Failure 400 {object} ErrorResponse "Некорректный формат данных или отсутствуют обязательные поля"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /create_katet [post]
func InsertKatet(c *fiber.Ctx) error {
	type Request struct {
		KatetValue float64 `json:"katet_value"`
	}

	var req Request

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(KatetErrorResponse{
			Error: "Неккоректный формат данных",
		})
	}

	if req.KatetValue == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(KatetErrorResponse{
			Error: "Поле katet_value обязательно и должно быть больше нуля",
		})
	}

	query := `insert into katets (katet_value) values ($1) returning katet_id`

	var newID int

	err := database.DB.QueryRow(query, req.KatetValue).Scan(&newID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(KatetErrorResponse{
			Error: "Не удалось вставить данные",
		})
	}

	f := fmt.Sprint(req.KatetValue)

	newKatet := models.Katets{
		KATET_VALUE: f,
	}

	return c.Status(fiber.StatusCreated).JSON(newKatet)
}

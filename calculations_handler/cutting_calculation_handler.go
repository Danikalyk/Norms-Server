package calculationshandler

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// @Summary Рассчитать время резки
// @Description Рассчитывает время, необходимое для резки металла по заданному периметру. Может работать в HTTP и Websocket режимах
// @Tags calculations
// @Accept json
// @Produce json
// @Param perimeter query int true "Периметр, мм."
// @Param insertion_count query int true "Количество врезок, шт."
// @Param cutting_speed query float64 true "Скорость резки, мм/с."
// @Param insertion_time query float64 true "Время врезки, сек."
// @Success 200 {object} map[string]string "{"result": "calculated time"}"
// @Failure 400 {string} string "Неверно указаны параметры"
// @Router /cutting_calculate [get]
func CalculateCuttingTime(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		// Обрабатываем WebSocket-соединение
		return websocketHandler(c)
	}

	// Обрабатываем обычный HTTP-запрос
	perimeter, err := strconv.Atoi(c.Query("perimeter"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Неверно указан периметр")
	}
	insertionCount, err := strconv.Atoi(c.Query("insertion_count"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Неверно указано количество врезок")
	}
	cuttingSpeed, err := strconv.ParseFloat(c.Query("cutting_speed"), 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Неверно указана скорость резки")
	}
	insertionTime, err := strconv.ParseFloat(c.Query("insertion_time"), 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Неверно указано время врезки")
	}

	result := ((float64(perimeter) / (cuttingSpeed * 1000)) + (float64(insertionCount) * insertionTime)) * 1.05
	formattedResult := fmt.Sprintf("%.2f", result)

	return c.JSON(fiber.Map{"result": formattedResult})
}

func websocketHandler(c *fiber.Ctx) error {
	return websocket.New(func(conn *websocket.Conn) {
		defer conn.Close()
		for {
			var message map[string]interface{}
			err := conn.ReadJSON(&message)
			if err != nil {
				log.Println("Ошибка чтения:", err)
				break
			}

			perimeter, ok := message["perimeter"].(float64)
			if !ok {
				conn.WriteJSON(fiber.Map{"error": "Неверно указан периметр"})
				continue
			}
			insertionCount, ok := message["insertion_count"].(float64)
			if !ok {
				conn.WriteJSON(fiber.Map{"error": "Неверно указано количество врезок"})
				continue
			}
			cuttingSpeed, ok := message["cutting_speed"].(float64)
			if !ok {
				conn.WriteJSON(fiber.Map{"error": "Неверно указана скорость резки"})
				continue
			}
			insertionTime, ok := message["insertion_time"].(float64)
			if !ok {
				conn.WriteJSON(fiber.Map{"error": "Неверно указано время врезки"})
				continue
			}

			result := ((perimeter / (cuttingSpeed * 1000)) + (insertionCount * insertionTime)) * 1.05
			formattedResult := fmt.Sprintf("%.2f", result)

			conn.WriteJSON(fiber.Map{"result": formattedResult})
		}
	})(c)
}

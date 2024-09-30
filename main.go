package main

import (
	calculationshandler "NormsServer/calculations_handler"
	"NormsServer/database"
	"NormsServer/handlers"
	"log"
	"os"

	_ "NormsServer/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	app := fiber.New()

	database.InitDB("postgres://postgres:Neoclassic!23@localhost:5432/RCLT_DATA?sslmode=disable")

	logFilePath := "logs/access.log"

	err := os.MkdirAll("logs", 0755)
	if err != nil {
		log.Fatalf("Не удалось создать каталог для логов: %v", err)
	}

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Не удалось открыть файл логов: %v", err)
	}
	defer func() {
		if err := logFile.Close(); err != nil {
			log.Printf("Ошибка при закрытии файла логов: %v", err)
		}
	}()

	app.Use(logger.New(logger.Config{
		Format:     "${time} ${ip} ${method} ${path} ${status} ${latency}\n",
		TimeFormat: "02-Jan-2006",
		Output:     logFile,
	}))

	app.Get("/machines", handlers.GetMachines)
	app.Get("/materials", handlers.GetMaterials)
	app.Get("/gases", handlers.GetGases)
	app.Get("/cutting_info", handlers.GetCuttingInfo)
	app.Get("/cutting_calculate", calculationshandler.CalculateCuttingTime)
	app.Get("/welding_materials", handlers.GetWeldingMaterials)
	app.Get("/welding_types", handlers.GetWeldingType)
	app.Get("/katets", handlers.GetKatet)
	app.Get("/seam_type", handlers.GetSeamType)
	app.Get("/cross_sec_area", handlers.GetCrossSecArea)
	app.Get("/wire_diameter", handlers.GetWireDiameter)
	app.Get("/tnsh", handlers.GetTNSH)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	log.Fatal(app.Listen(":3000"))
}

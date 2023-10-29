package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"io"
	"lab-argo-app/pkg/entities"
	"lab-argo-app/pkg/schedule"
	"lab-argo-app/web/routes"
	"log"
	"os"
)

func main() {
	dbFilePath := "./mock/schedule.json"
	dbFile, err := os.Open(dbFilePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer dbFile.Close()

	data, err := io.ReadAll(dbFile)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}
	var db entities.Schedule
	err = json.Unmarshal(data, &db)
	if err != nil {
		log.Fatalf("failed to unmarshal data: %s", err)
	}

	scheduleRepo := schedule.NewRepository(db)
	scheduleService := schedule.NewService(scheduleRepo)

	app := fiber.New()
	routes.ScheduleRouter(app, scheduleService)
	log.Fatal(app.Listen(":3000"))
}

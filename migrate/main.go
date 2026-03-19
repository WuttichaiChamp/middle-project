package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/Teeranuns000/middle-project/migrate/runner"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using system environment")
	}

	runner.MigrateOMS()
	// เพิ่ม database อื่นตรงนี้ในอนาคต เช่น
	// runner.MigrateWMS()
}

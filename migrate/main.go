package main

import (
	"flag"
	"log"
	"strings"

	"github.com/Teeranuns000/middle-project/migrate/runner"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using system environment")
	}

	db := flag.String("db", "", "target database: oms, wms, ...")
	tablesFlag := flag.String("tables", "", "specific tables to migrate, comma-separated e.g. courier,sale_item (leave empty to migrate all)")
	flag.Parse()

	if *db == "" {
		log.Fatal("usage: go run migrate/main.go -db=oms [-tables=courier,sale_item]")
	}

	var tables []string
	if *tablesFlag != "" {
		tables = strings.Split(*tablesFlag, ",")
	}

	switch *db {
	case "oms":
		runner.MigrateOMS(tables)
	// case "wms":
	// 	runner.MigrateWMS(tables)
	default:
		log.Fatalf("unknown db: %s", *db)
	}
}

package runner

import (
	"fmt"
	"log"
	"os"

	"github.com/Teeranuns000/middle-project/database/oms/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MigrateOMS() {
	dsn := os.Getenv("database_gorm_url_prime_oms")
	if dsn == "" {
		log.Fatal("database_gorm_url_prime_oms is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
	})
	if err != nil {
		log.Fatalf("failed to connect OMS database: %v", err)
	}

	tables := []any{
		// &models.Courier{},
		// &models.Shipment{},
		// &models.ShipmentItem{},
		// &models.Order{},
		// &models.OrderItem{},
		// &models.OrderPackage{},
		// &models.OrderPackageItem{},
		&models.PlatformStatement{},
		// &models.Platform{},
		// &models.PlatformLog{},
		// &models.Promotion{},
		// &models.PromotionItem{},
		// &models.PromotionItemSource{},
		// &models.PromotionItemDestination{},
		// &models.PromotionItemUsage{},
		// &models.PromotionItemUsageDestination{},
		// &models.Approval{},
		// &models.Quotation{},
		// &models.QuotationItem{},
		// &models.SalesOrder{},
		// &models.SalesOrderItem{},
		// &models.Shop{},
		// &models.ShopAuth{},
		// &models.SystemConfig{},
	}

	for _, t := range tables {
		if !db.Migrator().HasTable(t) {
			if err := db.Migrator().CreateTable(t); err != nil {
				log.Printf("Warning: create table for %T: %v", t, err)
			} else {
				log.Printf("Success: created table %T", t)
			}
		} else {
			migrateColumns(db, t)
		}
	}

	fmt.Println("OMS migration completed successfully")
}

func migrateColumns(db *gorm.DB, model any) {
	stmt := &gorm.Statement{DB: db}
	if err := stmt.Parse(model); err != nil {
		log.Printf("Warning: failed to parse %T: %v", model, err)
		return
	}

	tableName := stmt.Schema.Table
	for _, field := range stmt.Schema.Fields {
		if field.DBName == "" {
			continue
		}
		if !db.Migrator().HasColumn(model, field.DBName) {
			if err := db.Migrator().AddColumn(model, field.DBName); err != nil {
				log.Printf("Warning: add column %s.%s: %v", tableName, field.DBName, err)
			} else {
				log.Printf("Success: added column %s.%s", tableName, field.DBName)
			}
		}
	}
}

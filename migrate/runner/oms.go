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
	})
	if err != nil {
		log.Fatalf("failed to connect OMS database: %v", err)
	}

	tables := []any{
		&models.Courier{},
		&models.Shipment{},
		&models.ShipmentItem{},
		// &models.Order{},
		// &models.OrderItem{},
		// &models.OrderPackage{},
		// &models.OrderPackageItem{},
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
		if err := db.AutoMigrate(t); err != nil {
			log.Printf("Warning: migration for %T: %v", t, err)
		}
	}

	fmt.Println("OMS migration completed successfully")
}

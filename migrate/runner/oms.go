package runner

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Teeranuns000/middle-project/database/oms/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// registry ของ tableName → model instance
var omsRegistry = map[string]interface{}{
	"courier":                          &models.Courier{},
	"order":                            &models.Order{},
	"order_item":                       &models.OrderItem{},
	"order_package":                    &models.OrderPackage{},
	"order_package_item":               &models.OrderPackageItem{},
	"platform":                         &models.Platform{},
	"platform_log":                     &models.PlatformLog{},
	"promotion":                        &models.Promotion{},
	"promotion_item":                   &models.PromotionItem{},
	"promotion_item_source":            &models.PromotionItemSource{},
	"promotion_item_destination":       &models.PromotionItemDestination{},
	"promotion_item_usage":             &models.PromotionItemUsage{},
	"promotion_item_usage_destination": &models.PromotionItemUsageDestination{},
	"approval":                         &models.Approval{},
	"quotation":                        &models.Quotation{},
	"quotation_item":                   &models.QuotationItem{},
	"sale":                             &models.SalesOrder{},
	"sale_item":                        &models.SalesOrderItem{},
	"shop":                             &models.Shop{},
	"shop_auth":                        &models.ShopAuth{},
	"system_config":                    &models.SystemConfig{},
}

func MigrateOMS(tables []string) {
	dsn := os.Getenv("database_gorm_url_prime_oms")
	if dsn == "" {
		log.Fatal("database_gorm_url_prime_oms is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[oms] failed to connect: %v", err)
	}

	targets, err := resolveTargets(omsRegistry, tables)
	if err != nil {
		log.Fatalf("[oms] %v", err)
	}

	if err := db.AutoMigrate(targets...); err != nil {
		log.Fatalf("[oms] migration failed: %v", err)
	}

	fmt.Printf("[oms] migration completed: %s\n", strings.Join(migratedNames(tables, omsRegistry), ", "))
}

// resolveTargets แปลง table names → model instances, return error ถ้าหาไม่เจอ
func resolveTargets(registry map[string]interface{}, tables []string) ([]interface{}, error) {
	if len(tables) == 0 {
		all := make([]interface{}, 0, len(registry))
		for _, m := range registry {
			all = append(all, m)
		}
		return all, nil
	}

	var targets []interface{}
	var notFound []string

	for _, t := range tables {
		t = strings.TrimSpace(t)
		m, ok := registry[t]
		if !ok {
			notFound = append(notFound, t)
			continue
		}
		targets = append(targets, m)
	}

	if len(notFound) > 0 {
		return nil, fmt.Errorf("table(s) not found in registry: %s", strings.Join(notFound, ", "))
	}

	return targets, nil
}

func migratedNames(tables []string, registry map[string]interface{}) []string {
	if len(tables) == 0 {
		names := make([]string, 0, len(registry))
		for k := range registry {
			names = append(names, k)
		}
		return names
	}
	return tables
}

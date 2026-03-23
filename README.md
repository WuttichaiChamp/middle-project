# middle-project

แหล่งรวม shared models และ database migration สำหรับ microservices ทุกตัวในระบบ

## สิ่งที่ project นี้ทำ

1. เป็น shared Go module ให้ service อื่น import model, request struct, response struct ไปใช้ได้โดยตรง
2. เป็น migration runner สำหรับทุก database ในระบบ รันครั้งเดียวแล้วปิดตัวเอง ไม่ต้องค้างโปรเซสไว้

---

## Project Structure

```
middle-project/
├── database/
│   ├── oms/
│   │   ├── models/       # GORM models สำหรับ OMS database
│   │   └── request/      # Request structs (เพิ่มได้)
│   └── wms/              # WMS database (เพิ่มในอนาคต)
├── shared/
│   └── model/            # Common types เช่น Pagination
├── migrate/
│   ├── main.go           # Migration entry point
│   └── runner/
│       └── oms.go        # Registry + AutoMigrate สำหรับ OMS
├── .env                  # Database connection strings (ห้าม commit)
├── go.mod
└── go.sum
```

---

## Tables ที่มีใน OMS Registry

| Table name | Model |
|---|---|
| courier | Courier |
| order | Order |
| order_item | OrderItem |
| order_package | OrderPackage |
| order_package_item | OrderPackageItem |
| order_log | OrderLog |
| platform | Platform |
| platform_log | PlatformLog |
| promotion | Promotion |
| promotion_item | PromotionItem |
| promotion_item_source | PromotionItemSource |
| promotion_item_destination | PromotionItemDestination |
| promotion_item_usage | PromotionItemUsage |
| promotion_item_usage_destination | PromotionItemUsageDestination |
| approval | Approval |
| quotation | Quotation |
| quotation_item | QuotationItem |
| sale | SalesOrder |
| sale_item | SalesOrderItem |
| shop | Shop |
| shop_auth | ShopAuth |
| system_config | SystemConfig |

---

## การใช้งาน Migration

### ตั้งค่า .env

```env
database_gorm_url_prime_oms=host=<host> user=<user> password=<password> dbname=<dbname> port=5432 sslmode=disable
```

### รัน migrate

```bash
# migrate ทุก table ใน OMS
go run migrate/main.go -db=oms

# migrate เฉพาะบาง table
go run migrate/main.go -db=oms -tables=courier,sale_item

# migrate WMS (เมื่อเพิ่มในอนาคต)
go run migrate/main.go -db=wms
```

ถ้าระบุ table ที่ไม่มีใน registry จะได้ error ทันที:
```
[oms] table(s) not found in registry: unknown_table
```

### AutoMigrate ทำอะไรได้บ้าง

| การเปลี่ยนแปลง | AutoMigrate จัดการได้ |
|---|---|
| เพิ่ม table ใหม่ | ✅ |
| เพิ่ม column ใหม่ | ✅ |
| เพิ่ม index | ✅ |
| เปลี่ยนชื่อ column | ❌ ต้องแก้มือด้วย SQL |
| เปลี่ยน type ของ column | ❌ ต้องแก้มือด้วย SQL |
| ลบ column | ❌ ต้องแก้มือด้วย SQL |

---

## การ Import Model ไปใช้ใน Service อื่น

### ติดตั้ง

```bash
go get github.com/Teeranuns000/middle-project@v1.0.0
```

> ควร pin version เสมอ ห้ามใช้ @latest ใน production

### ตัวอย่างการใช้งาน

```go
import (
    "github.com/Teeranuns000/middle-project/database/oms/models"
)

var order models.Order
db.First(&order, "id = ?", id)

newOrder := models.Order{
    OrderCode: "ORD-001",
    ShopCode:  "SHOP001",
    Status:    "pending",
}
db.Create(&newOrder)
```

ถ้าใช้หลาย database พร้อมกันให้ใช้ alias เพื่อป้องกัน package name ชนกัน:

```go
import (
    omsmodels "github.com/Teeranuns000/middle-project/database/oms/models"
    wmsmodels "github.com/Teeranuns000/middle-project/database/wms/models"
)
```

---

## Checklist เมื่อต้องการเพิ่ม Table ใหม่

1. สร้าง model file ใน `database/oms/models/<name>.go`
2. ประกาศ struct และ `TableName()` ให้ถูกต้อง
3. เพิ่ม entry ใน `omsRegistry` ใน `migrate/runner/oms.go`
4. commit และ tag version ใหม่

```bash
git add .
git commit -m "feat: add order_log model"
git tag v1.1.0
git push origin main
git push origin v1.1.0
```

5. service ปลายทาง update version

```bash
go get github.com/Teeranuns000/middle-project@v1.1.0
go mod tidy
```

6. devops รัน migrate

```bash
go run migrate/main.go -db=oms -tables=order_log
```

---

## Checklist เมื่อต้องการเพิ่ม Database ใหม่

1. สร้าง folder `database/<db>/models/`
2. สร้าง `migrate/runner/<db>.go` โดย copy pattern จาก `oms.go`
   - สร้าง registry map ของ table name → model
   - สร้าง function `Migrate<DB>(tables []string)`
3. เพิ่ม case ใน `migrate/main.go`
4. เพิ่ม connection string ใน `.env`

---

## ข้อควรระวัง

- AutoMigrate ไม่ลบ column ออกจาก DB แม้จะลบออกจาก struct แล้ว
- การเปลี่ยนชื่อ field ใน struct AutoMigrate จะเพิ่ม column ใหม่และทิ้ง column เก่าไว้ ข้อมูลจะไม่ถูก migrate ให้อัตโนมัติ
- กรณีเหล่านี้ต้องรัน SQL เองก่อน แล้วค่อยแก้ struct ตาม
- ห้าม commit ไฟล์ .env ขึ้น repository

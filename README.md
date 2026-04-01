# middle-project

Single source of truth สำหรับ shared data models และ auto migration ข้าม microservices

## Project Structure

```
middle-project/
├── database/
│   ├── oms/
│   │   └── models/      # GORM structs สำหรับ OMS database
│   └── wms/
│       └── models/      # GORM structs สำหรับ WMS database
├── migrate/
│   ├── main.go          # Entry point สำหรับรัน migration
│   └── runner/
│       ├── oms.go       # Migration logic สำหรับ OMS
│       └── wms.go       # Migration logic สำหรับ WMS (ถ้ามี)
├── shared/
│   └── model/           # Common types เช่น Pagination
├── .env
└── go.mod
```

---

## Auto Migration

### วิธีรัน

```bash
go run ./migrate/main.go
```

### เพิ่ม Table ใหม่

1. สร้าง struct ใน `database/oms/models/` เช่น `warehouse.go`

```go
package models

import (
    "time"
    "github.com/google/uuid"
)

type Warehouse struct {
    ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    WarehouseCode string    `json:"warehouse_code"`
    WarehouseName string    `json:"warehouse_name"`
    IsActive      bool      `json:"is_active"`
    CreateBy      string    `json:"create_by"`
    CreateDtm     time.Time `json:"create_dtm"`
}

func (Warehouse) TableName() string { return "warehouse" }
```

2. เปิด comment ใน `migrate/runner/oms.go`

```go
tables := []any{
    // comment table ที่ไม่ต้องการแตะ
    // &models.Order{},
    // &models.Shop{},

    // เปิดเฉพาะที่ต้องการ migrate
    &models.Warehouse{},
}
```

3. รัน migration

```bash
go run ./migrate/main.go
```

> GORM `AutoMigrate` จะ skip table ที่มีอยู่แล้ว และสร้างเฉพาะ table/column ใหม่เท่านั้น ไม่ลบ data เดิม

### หลาย Database

แต่ละ database มี runner แยกกัน ควบคุมได้จาก `migrate/main.go`

```go
// main.go
runner.MigrateOMS() // comment ถ้าไม่ต้องการแตะ OMS
runner.MigrateWMS() // comment ถ้าไม่ต้องการแตะ WMS
```

### .env

```env
database_gorm_url_prime_oms=host=<host> user=<user> password=<password> dbname=<dbname> port=5432 sslmode=disable
database_gorm_url_prime_wms=host=<host> user=<user> password=<password> dbname=<dbname> port=5432 sslmode=disable
```

---

## ใช้ Models ข้าม Repo (Private Repo)

### 1. ตั้งค่า GONOSUMCHECK และ GOFLAGS (ทำครั้งเดียว)

เนื่องจาก repo นี้เป็น private ต้องบอก Go ให้ข้าม checksum database

```bash
go env -w GONOSUMCHECK=github.com/Teeranuns000/*
go env -w GONOSUMDB=github.com/Teeranuns000/*
go env -w GOPRIVATE=github.com/Teeranuns000/*
```

### 2. ตั้งค่า Git ให้ใช้ SSH แทน HTTPS

```bash
git config --global url."git@github.com:".insteadOf "https://github.com/"
```

หรือถ้าใช้ Personal Access Token (PAT):

```bash
git config --global url."https://<YOUR_TOKEN>@github.com/".insteadOf "https://github.com/"
```

### 3. ติดตั้ง package ใน repo ปลายทาง

ถ้าทีมมีการสื่อสารและ sync กันดี สามารถดึงตรงจาก branch ได้เลยโดยไม่ต้อง tag version

```bash
go get github.com/Teeranuns000/middle-project@main
```

> ใช้ `@main` แทน `@latest` เพราะ `@latest` ใน Go ดึงเฉพาะ tag ล่าสุด ถ้าไม่มี tag จะ error

หรือถ้าต้องการ lock version เพื่อความปลอดภัย

```bash
go get github.com/Teeranuns000/middle-project@v1.0.0
```

### 4. Import และใช้งาน

```go
import (
    omsModels "github.com/Teeranuns000/middle-project/database/oms/models"
)

var shop omsModels.Shop
db.First(&shop, "shop_code = ?", "SHOP001")
```

### 5. Update Version

เมื่อมีการแก้ไข model ให้ทำตามขั้นตอนนี้

```bash
git add .
git commit -m "feat: add Warehouse model"
git tag v1.1.0
git push origin main
git push origin v1.1.0
```

แล้วใน repo ปลายทางรัน:

```bash
go get github.com/Teeranuns000/middle-project@v1.1.0
```

---

## Rules

- ห้าม define struct ซ้ำใน service ให้ import จาก repo นี้เท่านั้น
- ถ้าทีม sync กันดีสามารถใช้ `@main` ได้เลย ไม่จำเป็นต้อง tag version
- ถ้ามี relation (has many) ต้องใส่ `gorm:"foreignKey:XxxID"` ด้วยทุกครั้ง
- comment table ที่ไม่ต้องการแตะก่อนรัน migration เสมอ

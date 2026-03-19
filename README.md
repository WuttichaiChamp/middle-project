# common-models

Single source of truth for shared data models across microservices.

## Project Structure

```
common-models/
├── user/
│   ├── model/       # DB + business structs
│   ├── request/     # API request DTOs
│   └── response/    # API response DTOs
├── order/
│   ├── model/
│   ├── request/
│   └── response/
├── wms/
│   ├── model/
│   ├── request/
│   └── response/
├── oms/
│   ├── model/
│   ├── request/
│   └── response/
└── shared/
    └── model/       # Common types (Pagination, etc.)
```

## Naming Convention

| Layer | Package | Example |
|-------|---------|---------|
| DB model | `model` | `ShipmentRow` |
| Business model | `model` | `Shipment` |
| Request DTO | `request` | `CreateShipmentRequest` |
| Response DTO | `response` | `ShipmentResponse` |

## Usage

```go
import (
    omsModel    "github.com/your-org/common-models/oms/model"
    omsRequest  "github.com/your-org/common-models/oms/request"
    omsResponse "github.com/your-org/common-models/oms/response"
)

// Use model
var row omsModel.ShipmentRow
db.Get(&row, "SELECT * FROM shipments WHERE id = ?", id)

// Use request DTO
var req omsRequest.CreateShipmentRequest
json.Unmarshal(body, &req)

// Use response DTO
res := &omsResponse.ShipmentResponse{
    ID:     row.ID,
    Status: row.Status,
}
```

## Rules

- ห้าม define struct ซ้ำใน service
- ห้าม import ข้าม domain โดยไม่จำเป็น
- ต้อง lock version เสมอ ห้ามใช้ `@latest`
- เฉพาะ owner เท่านั้นที่แก้ไข domain ของตัวเองได้

## Domain Ownership

| Domain | Owner |
|--------|-------|
| user/* | user-service |
| order/* | order-service |
| wms/* | wms-service |
| oms/* | oms-service |

## Install

```bash
go get github.com/your-org/common-models@v1.0.0
```

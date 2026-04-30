# 🚛 Fleet Maintenance Management System (Golang + Next.js)

## 📌 Overview

Hệ thống WebApp quản lý bảo trì đội xe vận tải, bao gồm:

* Quản lý phương tiện
* Quản lý vật tư & kho
* Quản lý bảo trì
* Quy trình duyệt phiếu
* Phân quyền người dùng

### 🎯 Goal

Xây dựng hệ thống CRUD hoàn chỉnh + workflow duyệt + quản lý kho + dashboard.

---

## 🏗️ Tech Stack

### Backend

* Golang (Gin hoặc Fiber)
* GORM hoặc SQLX
* JWT Authentication

### Frontend

* Next.js (App Router)
* React Hook Form + Zod
* Shadcn UI
* TanStack Table
* React Query / Axios

### Database

* My sql

---

## 🧱 Architecture

### Backend (Clean Architecture)

```
/internal
  /handler     → HTTP layer
  /service     → business logic
  /repository  → database access
  /model       → entity
  /dto         → request/response
```

### Frontend

```
/app
/components
/services
/hooks
/types
```

---

## 📊 Core Modules

---

## A. 🚗 Vehicle Management

### Features

* CRUD phương tiện
* Xem lịch sử bảo trì

### Table: Vehicle

```sql
CREATE TABLE Vehicle (
  Id UNIQUEIDENTIFIER PRIMARY KEY,
  LicensePlate NVARCHAR(50),
  Model NVARCHAR(100),
  Year INT,
  Note NVARCHAR(MAX),
  CreatedAt DATETIME DEFAULT GETDATE()
);
```

---

## B. 📦 Inventory & Warehouse

### Features

* Danh mục vật tư
* Theo dõi tồn kho
* Nhập / Xuất / Điều chỉnh

### Tables

```sql
CREATE TABLE Item (
  Id UNIQUEIDENTIFIER PRIMARY KEY,
  Name NVARCHAR(255),
  Category NVARCHAR(100),
  Unit NVARCHAR(50)
);

CREATE TABLE Warehouse (
  Id UNIQUEIDENTIFIER PRIMARY KEY,
  Name NVARCHAR(255),
  Location NVARCHAR(255)
);

CREATE TABLE Stock (
  Id UNIQUEIDENTIFIER PRIMARY KEY,
  ItemId UNIQUEIDENTIFIER,
  WarehouseId UNIQUEIDENTIFIER,
  Quantity INT
);

CREATE TABLE StockTransaction (
  Id UNIQUEIDENTIFIER PRIMARY KEY,
  Type NVARCHAR(20), -- IN | OUT | ADJUST
  ItemId UNIQUEIDENTIFIER,
  WarehouseId UNIQUEIDENTIFIER,
  Quantity INT,
  ReferenceId UNIQUEIDENTIFIER,
  CreatedAt DATETIME DEFAULT GETDATE()
);
```

---

## C. 🏢 Supplier Management

```sql
CREATE TABLE Supplier (
  Id UNIQUEIDENTIFIER PRIMARY KEY,
  Name NVARCHAR(255),
  Phone NVARCHAR(50),
  Address NVARCHAR(255)
);
```

---

## D. 📄 Maintenance Request

### Workflow

```
Pending → Approved → Auto OUT stock
        → Rejected
```

### Tables

```sql
CREATE TABLE MaintenanceRequest (
  Id UNIQUEIDENTIFIER PRIMARY KEY,
  VehicleId UNIQUEIDENTIFIER,
  Status NVARCHAR(20),
  CreatedBy UNIQUEIDENTIFIER,
  CreatedAt DATETIME DEFAULT GETDATE()
);

CREATE TABLE MaintenanceRequestItem (
  Id UNIQUEIDENTIFIER PRIMARY KEY,
  RequestId UNIQUEIDENTIFIER,
  ItemId UNIQUEIDENTIFIER,
  Quantity INT
);
```

---

## E. 🔧 Maintenance Log

```sql
CREATE TABLE MaintenanceLog (
  Id UNIQUEIDENTIFIER PRIMARY KEY,
  VehicleId UNIQUEIDENTIFIER,
  Date DATETIME,
  Odometer INT,
  Description NVARCHAR(MAX)
);

CREATE TABLE MaintenanceLogItem (
  Id UNIQUEIDENTIFIER PRIMARY KEY,
  LogId UNIQUEIDENTIFIER,
  ItemId UNIQUEIDENTIFIER,
  Quantity INT
);
```

---

## F. 👤 User & Role

```sql
CREATE TABLE [User] (
  Id UNIQUEIDENTIFIER PRIMARY KEY,
  Username NVARCHAR(100),
  PasswordHash NVARCHAR(255),
  Role NVARCHAR(50)
);
```

### Roles

* Admin
* Technician
* Warehouse

---

## 🔐 Authentication

* JWT Authentication

### Flow

1. Login → nhận Access Token
2. Gửi request kèm header:

```
Authorization: Bearer <token>
```

---

## 🔄 API Design

### Vehicle

```
GET    /api/vehicles
POST   /api/vehicles
GET    /api/vehicles/:id
PUT    /api/vehicles/:id
DELETE /api/vehicles/:id
```

---

### Maintenance Request

```
POST   /api/maintenance-requests
GET    /api/maintenance-requests
POST   /api/maintenance-requests/:id/approve
POST   /api/maintenance-requests/:id/reject
```

---

## ⚙️ Business Rules

### 1. Không cho xuất kho nếu không đủ tồn

```go
if stock.Quantity < requestQty {
    return errors.New("not enough stock")
}
```

---

### 2. Approve Request

* Update status = Approved
* Tạo StockTransaction (OUT)
* Trừ tồn kho

---

### 3. Maintenance Log

* Ghi nhận vật tư đã sử dụng
* Liên kết kho

---

## 📈 Dashboard

* Tổng số xe
* Tổng tồn kho
* Số phiếu Pending
* Lịch sử bảo trì gần nhất

---

## 🎨 Frontend Pages

```
/dashboard
/vehicles
/inventory
/requests
/maintenance
/users
```

---

## 🧪 Validation

### Frontend

* Zod schema

### Backend

* Validate struct (binding / validator)

---

## 🚀 Deployment

### Backend

```bash
go mod tidy
go build -o app
./app
```

### Frontend

```bash
npm install
npm run build
npm start
```

---

## 📦 Deliverables

* Source code (Frontend + Backend)
* SQL Script
* ERD Diagram
* Swagger API Docs
* Deployment Guide

---

## 🛠️ Suggested Libraries

### Golang

* gin-gonic/gin
* gorm.io/gorm
* golang-jwt/jwt

### Next.js

* react-hook-form
* zod
* shadcn/ui
* @tanstack/react-query

---

## 🔧 Future Improvements

* Barcode / QR vật tư
* Mobile App
* Notification (Email, Zalo)
* Predictive Maintenance (AI)

---

## 🤖 Notes for AI Agent

* Sử dụng Clean Architecture
* Không viết logic trong handler
* Tách DTO rõ ràng
* Validate đầy đủ input
* API chuẩn REST
* Code dễ mở rộng, maintain

---


---
Mỗi lần chạy đều review code và check lỗi để fix lỗi 

# IAM Domain (Identity & Access Management)

---

## 1. Overview

IAM Domain thuộc **Platform Core Layer**.

IAM chịu trách nhiệm:

- Xác thực (Authentication)
- Quản lý danh tính (Identity)
- Định nghĩa Permission chuẩn hệ thống
- Thực thi Authorization Engine

IAM **không biết business là F&B hay Retail**.  
IAM chỉ biết:

- Ai là ai?
- User đang hoạt động trong tenant nào?
- User có permission gì trong context đó?

---

## 2. Vai trò trong kiến trúc

IAM là **Security Core** của toàn hệ thống.

Tất cả request phải đi qua:

1. Authentication (xác thực token)
2. Authorization (kiểm tra permission theo context)

Không có IAM → không có kiểm soát truy cập.

---

## 3. Core Entities

> ⚠ Lưu ý: IAM không chứa Business Role như “Thu ngân”.

---

### 3.1 User (Aggregate Root)

Đại diện cho một identity toàn hệ thống.

Thuộc tính:

- id
- email
- phone
- passwordHash
- status (ACTIVE / BLOCKED / DELETED)
- createdAt
- lastLoginAt

User tồn tại độc lập với Organization.

---

### 3.2 Permission (Core Entity)

Permission là định nghĩa kỹ thuật của quyền truy cập.

Permission gồm:

- resource (STORE_OWNER, UNIT, ORDER, ORGANIZATION…)
- action (CREATE, UPDATE, DELETE, VIEW…)
- scope (GLOBAL / TENANT / UNIT / OWN)

Ví dụ:

| id | resource      | action | scope  |
|----|--------------|--------|--------|
| 1  | ORDER        | CREATE | UNIT   |
| 2  | ORDER        | REFUND | UNIT   |
| 3  | ORGANIZATION | DELETE | TENANT |

Permission là global, không phụ thuộc tenant.

---

### 3.3 PlatformRole (Optional)

Dành cho hệ thống nội bộ.

Ví dụ:

- SUPER_ADMIN
- SUPPORT
- AUDITOR

---

### 3.4 PlatformRolePermission

Mapping giữa PlatformRole và Permission.

---

### 3.5 Session / Token

Bao gồm:

- accessToken (JWT)
- refreshToken
- expiredAt
- currentTenantId (optional context)

---

## 4. Authorization Model

IAM sử dụng **Context-based Authorization**.

Permission check dựa vào:
(userId, organizationId, resource, action, resourceId)

Không chỉ đơn giản là: user có permission X?

Mà là: user có permission X trong tenant Y không?


---

## 5. Multi-tenant Design

User có thể:

- Thuộc nhiều Organization
- Có vai trò khác nhau ở mỗi Organization

Ví dụ:

User A:
- OWNER ở Org A
- STAFF ở Org B

IAM không lưu “STAFF”.

IAM chỉ lưu:

- User
- Permission definitions
- Authorization engine

Role cụ thể của Organization nằm ở Organization Domain.

---

## 6. Responsibilities

### ✔ Register

- Tạo User
- Hash password
- Validate email / phone

---

### ✔ Login

- Verify password
- Sinh JWT
- Tạo refreshToken
- Thiết lập tenant context

---

### ✔ Token validation

- Verify chữ ký
- Kiểm tra expired
- Resolve user context

---

### ✔ Authorization Engine

- Nhận resource + action
- Resolve permission key
- Check permission theo tenant context
- Check scope (GLOBAL / TENANT / UNIT / OWN)
- Optional: check ownership

---

### ✔ User lifecycle

- Block
- Unblock
- Soft delete

---

## 7. IAM KHÔNG làm gì

IAM không:

- Không biết Organization business gì
- Không chứa Business Role (Thu ngân, Nhân viên…)
- Không quản lý Unit
- Không xử lý Order
- Không xử lý Payment
- Không quản lý Subscription
- Không quyết định cấu trúc tenant

IAM chỉ là Security Engine.

---

## 8. Quan hệ với Domain khác

| Domain | Quan hệ |
|--------|---------|
| Organization | Gán User vào Organization + Business Role |
| Unit | IAM chỉ check permission, không quản lý Unit |
| Subscription | Có thể giới hạn số user qua policy |
| Business Vertical | Interceptor gọi IAM để authorize |

---

## 9. Phân biệt IAM vs Organization

| IAM | Organization |
|-----|--------------|
| Quản lý User toàn hệ thống | Quản lý tenant |
| Định nghĩa Permission | Định nghĩa Business Role |
| Authorization engine | Gán User vào Role |
| Context-based access | Cấu trúc tổ chức |
| Không chứa nghiệp vụ | Chứa logic phân quyền nội bộ |

---

## 10. Authorization Flow

Interceptor pseudo code:

```go
opts := proto.GetMessageOptions(msg)

if opts.skip_auth {
    return next()
}

resource := opts.resource
action := opts.action

tenantId := extractTenant(ctx)

authorize(userId, tenantId, resource, action, resourceId)

IAM chỉ thực thi authorize.

Mapping user ↔ role ↔ permission một phần nằm ở Organization Domain.

---

## 11. Enterprise Authorization Structure
Hai tầng:
- Platform-level: SUPER_ADMIN, SUPPORT
- Tenant-level (Organization Domain): OWNER, ADMIN, STAFF, THU NGÂN, NHÂN VIÊN F&B
Organization Role → map tới Permission (định nghĩa ở IAM).
IAM không định nghĩa Business Role.

## 12. Boundary quan trọng
IAM:
- Không truy cập dữ liệu business
- Không phụ thuộc vertical
- Không chứa role theo ngành
- Chỉ cung cấp engine và permission model

## 13. Summary

IAM Domain là:
- Trung tâm Identity
- Trung tâm Permission Definition
- Trung tâm Authorization Engine
- Multi-tenant aware
- Context-based access control
- Organization Domain là:
- Tenant container
- Business Role definition
- User membership management
- Không có IAM → hệ thống không an toàn.
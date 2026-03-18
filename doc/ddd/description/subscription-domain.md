# Subscription Domain

## 1. Overview

Subscription Domain thuộc **Platform Core Layer**.

Subscription đại diện cho **gói dịch vụ SaaS mà một Organization đăng ký sử dụng**.

Nó kiểm soát:

- Organization có được sử dụng hệ thống hay không
- Được sử dụng bao nhiêu **Unit**
- Bao nhiêu **User**
- Được phép sử dụng **Module / Feature nào**

Subscription **không xử lý thanh toán tiền**.

Thanh toán thuộc **Billing Domain**.

---

# 2. Vai trò trong kiến trúc

Subscription là cơ chế **kiểm soát quyền sử dụng hệ thống SaaS**.

Mỗi Organization có **một Subscription active tại một thời điểm**.

Organization (1) ---- (1) Subscription


Subscription xác định:

- Organization có quyền truy cập **module nào**
- Organization có **bao nhiêu tài nguyên**

---

# 3. Entity bên trong Subscription Context

## 3.1 Subscription (Aggregate Root)

Đại diện cho **gói dịch vụ đang active của Organization**.

### Thuộc tính

- `id`
- `organizationId`
- `planId`
- `status` (TRIAL / ACTIVE / EXPIRED / CANCELLED)
- `startDate`
- `expireDate`
- `autoRenew`
- `enabledModules`
- `createdAt`

`enabledModules` xác định Organization được phép sử dụng module nào.

### Ví dụ

```yaml
enabledModules:
  - FNB
  - POS
  - INVENTORY
```

## 3.2 Plan

Plan định nghĩa gói dịch vụ chuẩn của hệ thống.

Thuộc tính

id

name (Basic / Pro / Enterprise)

maxUnit

maxUser

includedModules

price

Ví dụ

Plan: FNB_PRO

maxUnit: 5
maxUser: 20

includedModules:
  - FNB
  - POS
  - INVENTORY

Plan chỉ là template.

Khi Subscription được activate → modules sẽ được copy từ Plan.

## 3.3 Module

Module đại diện cho một capability hoặc sản phẩm của hệ thống.

Thuộc tính

- id
- code
- name
- description

Modules
-------
FNB
RETAIL
SPA
POS
INVENTORY
CRM
LOYALTY

Modules cho phép hệ thống mở rộng vertical sau này.

## 3.4 Usage (Optional)

Usage theo dõi mức sử dụng thực tế của Organization.

Thuộc tính:
- currentUnitCount
- currentUserCount
- storageUsed
- apiCallUsed

Usage có thể được lưu trong:
- Database
- Analytics service
- Usage tracking service

# 4. Trách nhiệm (Responsibilities)

Subscription Domain chịu trách nhiệm:

✔ Kích hoạt Subscription
- Tạo Subscription khi Organization đăng ký
- Gán Plan
- Copy modules từ Plan
- Set thời gian hiệu lực

✔ Kiểm tra giới hạn tài nguyên
- Organization có vượt quá Unit limit không
- Có vượt quá User limit không

✔ Kiểm tra quyền sử dụng Module
- Organization có quyền sử dụng module nào
- Module đó có được enable không

✔ Kiểm tra trạng thái Subscription
- Subscription còn hạn không
- Subscription có đang bị cancel không

✔ Gia hạn / Hết hạn
- Chuyển trạng thái EXPIRED
- Emit domain event


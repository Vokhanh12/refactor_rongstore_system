# Member Domain

## 1. Overview

Member Domain thuộc **Platform Core Layer**.

Member đại diện cho mối quan hệ giữa:

User (IAM) ↔ Organization ↔ Unit

Member KHÔNG phải là User.
User là danh tính toàn hệ thống.
Member là vai trò của User trong một Organization cụ thể.

---

## 2. Tại sao cần Member Domain?

Một User có thể:

- Thuộc nhiều Organization
- Có role khác nhau ở mỗi Organization
- Có quyền khác nhau ở mỗi Unit

Ví dụ:

User A:
- OWNER ở Organization A
- STAFF ở Organization B

Member Domain giải quyết bài toán multi-tenant authorization.

---

## 3. Phân biệt User vs Member

| User (IAM) | Member |
|------------|--------|
| Danh tính toàn hệ thống | Quan hệ trong Organization |
| Có email / password | Có role |
| Không biết business | Biết context Organization |
| Global identity | Tenant identity |

---

## 4. Entity bên trong Member Context

### 4.1 Member (Aggregate Root)

Thuộc tính ví dụ:

- id
- userId
- organizationId
- role (OWNER / ADMIN / STAFF)
- status (ACTIVE / INVITED / SUSPENDED)
- joinedAt

---

### 4.2 UnitAssignment (Optional)

Nếu phân quyền theo Unit:

- memberId
- unitId
- unitRole (MANAGER / CASHIER / WAITER)

---

### 4.3 Invitation

Dùng khi mời user vào Organization:

- email
- organizationId
- role
- token
- expiredAt

---

## 5. Trách nhiệm (Responsibilities)

Member Domain chịu trách nhiệm:

### ✔ Thêm user vào Organization

- Kiểm tra Subscription limit
- Gán role
- Tạo Member record

---

### ✔ Xóa member khỏi Organization

- Kiểm tra quyền người thực hiện
- Soft delete member

---

### ✔ Thay đổi role

- OWNER không thể bị xóa nếu là owner cuối cùng
- Validate quyền

---

### ✔ Gán quyền theo Unit

- Assign member vào Unit
- Giới hạn phạm vi thao tác

---

## 6. Hành vi chính (Core Behaviors)

### 6.1 Add Member
addMember(userId, organizationId, role)


- Validate Organization
- Validate subscription
- Tạo Member

---

### 6.2 Change Role
changeRole(memberId, newRole)

- Validate permission
- Update role

---

### 6.3 Assign Unit
assignToUnit(memberId, unitId, unitRole)

- Validate unit thuộc organization
- Lưu assignment

---

### 6.4 Remove Member
removeMember(memberId)

- Không cho phép xóa owner cuối cùng
- Update status

---

## 7. Member KHÔNG làm gì

- Không xác thực user (IAM làm)
- Không xử lý business F&B
- Không quản lý Subscription
- Không xử lý thanh toán

Member chỉ xử lý tenant-level relationship.

---

## 8. Quan hệ với Domain khác

| Domain | Quan hệ |
|--------|---------|
| IAM | Member liên kết tới User |
| Organization | Member thuộc Organization |
| Unit | Member có thể được assign vào Unit |
| Subscription | Giới hạn số member |
| Business Vertical | Kiểm tra role khi thao tác |

---

## 9. Flow ví dụ thực tế

### Case: Owner mời nhân viên

1. Owner gọi addMember
2. Member Domain kiểm tra subscription limit
3. Tạo Invitation
4. Khi user accept → tạo Member ACTIVE

---

### Case: Nhân viên thao tác trong Unit

1. Hệ thống check:
   (userId, organizationId, unitId, action)
2. IAM xác thực
3. Member Domain kiểm tra role
4. Nếu hợp lệ → cho phép

---

## 10. Boundary quan trọng

Member:

- Không truy cập Organization khác
- Không cho phép user thao tác ngoài Organization của mình
- Không cho phép vượt subscription limit

Member là lớp kiểm soát tenant-level access.

---

## 11. Vì sao Member thuộc Platform Core

Vì:

- Không phụ thuộc ngành nghề
- Dùng chung cho F&B, Retail, Service
- Là phần bắt buộc của hệ thống multi-tenant

Nếu mai bạn thêm vertical mới → Member không đổi.

---

## 12. Tóm tắt

Member Domain là:

- Cầu nối giữa User và Organization
- Quản lý role tenant-level
- Hỗ trợ phân quyền theo Unit
- Giải quyết bài toán multi-tenant
- Không chứa business logic ngành

Không có Member → không thể phân quyền đúng trong SaaS multi-tenant.
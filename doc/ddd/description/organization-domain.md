# Organization Domain

## 1. Overview

Organization Domain thuộc **Platform Core Layer**.

Nó đại diện cho thực thể pháp lý hoặc chủ thể sở hữu hệ thống, ví dụ:

- Công ty
- Chuỗi cửa hàng
- Doanh nghiệp cá nhân
- Thương hiệu

Organization KHÔNG phải là chi nhánh.
Organization là "gốc" của toàn bộ hệ thống business phía dưới.

---

## 2. Vai trò trong kiến trúc

Organization là root aggregate của toàn bộ hệ thống multi-tenant.

Mọi thứ đều xoay quanh Organization:

- Unit thuộc Organization
- Member thuộc Organization
- Subscription gắn với Organization
- Billing gắn với Organization

Nếu không có Organization → hệ thống không tồn tại tenant.

---

## 3. Entity bên trong Organization Context

### 3.1 Organization (Aggregate Root)

Đại diện cho một tổ chức.

Thuộc tính ví dụ:

- id
- name
- legalName
- taxCode
- status (ACTIVE / SUSPENDED / DELETED)
- createdAt

---

### 3.2 OrgMember

Đại diện cho người dùng thuộc Organization.

Thuộc tính:

- userId
- organizationId
- role (OWNER / ADMIN / STAFF)
- joinedAt
- status

---

### 3.3 Subscription

Đại diện cho gói dịch vụ Organization đang sử dụng.

Thuộc tính:

- planId
- startDate
- expireDate
- status
- limits (maxUnit, maxUser, etc)

---

## 4. Trách nhiệm (Responsibilities)

Organization Domain chịu trách nhiệm:

### ✔ Tạo Organization
- Đăng ký tổ chức mới
- Gán OWNER đầu tiên

### ✔ Quản lý thành viên
- Thêm member
- Xóa member
- Thay đổi role
- Kiểm tra quyền

### ✔ Quản lý trạng thái Organization
- Kích hoạt
- Tạm khóa
- Đóng

### ✔ Kiểm tra Subscription
- Organization có còn hiệu lực không?
- Có vượt giới hạn không?

---

## 5. Hành vi chính (Core Behaviors)

### 5.1 Create Organization
createOrganization(name, ownerUserId)

- Tạo organization
- Gán owner
- Khởi tạo subscription mặc định

---

### 5.2 Add Member
addMember(userId, role)

- Kiểm tra quyền người thêm
- Kiểm tra subscription limit
- Thêm member

---

### 5.3 Suspend Organization
suspend()

- Đổi trạng thái
- Ngăn mọi Unit hoạt động

---

### 5.4 Check Permission
canPerform(userId, action)

- Kiểm tra role
- Trả về true/false

---

## 6. Organization KHÔNG làm gì

- Không quản lý sản phẩm
- Không xử lý order
- Không xử lý thanh toán khách hàng
- Không quản lý menu
- Không hiển thị bản đồ

Những thứ đó thuộc Business Vertical (F&B, Retail, Service…)

---

## 7. Quan hệ với các Domain khác

| Domain | Quan hệ |
|--------|---------|
| IAM | Xác thực user |
| Unit | Unit thuộc Organization |
| Subscription | Gắn với Organization |
| Billing | Tính phí theo Organization |
| F&B | Hoạt động dưới Organization |

---

## 8. Tại sao Organization thuộc Platform Core

Vì:

- Nó không phụ thuộc ngành nghề
- Dùng chung cho F&B, Retail, Service, HR
- Là nền tảng multi-tenant

Nếu đổi ngành → Organization vẫn giữ nguyên.

---

## 9. Quy tắc quan trọng

Organization là boundary phân tách dữ liệu.

Không được:
- Query chéo Organization
- Cho phép Unit của Org A truy cập dữ liệu Org B

Organization chính là ranh giới bảo mật và dữ liệu.

---

## 10. Tóm tắt

Organization Domain là:

- Gốc của hệ thống multi-tenant
- Quản lý cấu trúc tổ chức
- Quản lý member và quyền
- Quản lý subscription
- Bảo vệ boundary dữ liệu

Không chứa business của ngành cụ thể.
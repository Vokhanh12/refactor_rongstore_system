# Unit Domain

## 1. Overview

Unit Domain thuộc **Platform Core Layer**.

Unit đại diện cho một đơn vị vận hành cụ thể thuộc Organization.

Ví dụ:

- Một chi nhánh cửa hàng
- Một nhà hàng trong chuỗi
- Một cửa tiệm
- Một văn phòng
- Một kho hàng

Nếu Organization là "công ty mẹ"  
thì Unit là "điểm vận hành thực tế".

---

## 2. Vai trò trong kiến trúc

- Unit là cấp hoạt động thực tế của business
- Mọi Order, Menu, Staff, Payment… đều diễn ra ở Unit
- Unit thuộc về duy nhất một Organization

Quan hệ:

Organization (1) ------ (N) Unit

---

## 3. Entity bên trong Unit Context

### 3.1 Unit (Aggregate Root)

Thuộc tính ví dụ:

- id
- organizationId
- name
- code
- address
- latitude
- longitude
- status (ACTIVE / CLOSED / SUSPENDED)
- createdAt

---

### 3.2 UnitSettings

Cấu hình riêng của từng Unit

Ví dụ:

- timezone
- currency
- openHours
- serviceMode (DINE_IN / TAKE_AWAY / DELIVERY)
- taxRate

---

### 3.3 UnitStaff (tùy chọn nếu không dùng chung OrgMember)

- userId
- unitId
- role (MANAGER / CASHIER / WAITER)
- status

---

## 4. Trách nhiệm (Responsibilities)

Unit Domain chịu trách nhiệm:

### ✔ Tạo Unit

- Kiểm tra Organization tồn tại
- Kiểm tra Subscription limit
- Lưu thông tin cơ bản

---

### ✔ Quản lý trạng thái Unit

- Mở / Đóng cửa
- Tạm ngưng hoạt động

---

### ✔ Quản lý vị trí

- Lưu latitude, longitude
- Validate địa chỉ
- Gọi Map Adapter để lấy tọa độ (nếu cần)

---

### ✔ Kiểm tra quyền theo Unit

Ví dụ:
- Nhân viên chỉ được thao tác trong Unit của mình
- Không được truy cập Unit khác

---

## 5. Hành vi chính (Core Behaviors)

### 5.1 Create Unit
createUnit(organizationId, name, address)

- Kiểm tra Organization hợp lệ
- Kiểm tra Subscription limit
- Lưu Unit
- (Optional) Gọi Map Adapter để lấy tọa độ

---

### 5.2 Activate / Suspend Unit
activate()
suspend()
close()

- Đổi trạng thái
- Emit domain event

---

### 5.3 Update Location

updateLocation(address)

- Gọi Map Adapter
- Lưu lat/lng
- Cập nhật địa chỉ chuẩn hóa

---

### 5.4 Check Unit Access

canAccess(userId)


- Kiểm tra user có thuộc Organization
- Kiểm tra user có role hợp lệ

---

## 6. Unit KHÔNG làm gì

- Không xử lý Order
- Không quản lý Menu
- Không tính tiền
- Không xử lý thanh toán khách hàng
- Không làm Loyalty

Những thứ đó thuộc Business Vertical (F&B, Retail…)

---

## 7. Quan hệ với Domain khác

| Domain | Quan hệ |
|--------|---------|
| Organization | Unit thuộc Organization |
| IAM | Xác thực user |
| Subscription | Giới hạn số Unit |
| F&B | Business chạy trong Unit |
| Map Adapter | Hỗ trợ lấy tọa độ |

---

## 8. Vì sao Unit thuộc Platform Core

Vì:

- Mọi ngành đều có khái niệm "điểm vận hành"
- F&B có chi nhánh
- Retail có cửa hàng
- Service có văn phòng
- HR có office

Unit là khái niệm nền tảng, không phụ thuộc ngành.

---

## 9. Boundary quan trọng

Unit phải:

- Không truy cập dữ liệu Unit khác
- Không truy cập Organization khác
- Luôn validate organizationId trước khi thao tác

Unit là boundary cấp 2 sau Organization.

---

## 10. Phân biệt Organization vs Unit

| Organization | Unit |
|-------------|------|
| Chủ thể pháp lý | Điểm vận hành |
| Có Subscription | Không có Subscription riêng |
| Quản lý Member | Quản lý hoạt động thực tế |
| Multi-tenant root | Operational root |

---

## 11. Tóm tắt

Unit Domain là:

- Điểm vận hành thực tế
- Thuộc Organization
- Chứa thông tin vị trí và trạng thái
- Không chứa business ngành
- Là nơi Business Vertical chạy vào

Không có Unit → không có hoạt động thực tế.
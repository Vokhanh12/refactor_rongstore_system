# POS App (Point Of Sale Application)

## 1. Overview

POS App thuộc **Application Layer**.

POS không phải Business Domain.
POS là ứng dụng dùng để:

- Thực hiện bán hàng tại điểm bán
- Tương tác với nhân viên thu ngân
- Gọi các Domain xử lý nghiệp vụ

POS có thể dùng cho:

- F&B
- Retail
- Service

POS chỉ điều phối — không chứa business rule lõi.

---

## 2. Vai trò trong kiến trúc

POS App đóng vai trò:

- Giao diện thao tác bán hàng
- Điều phối use case
- Kết nối Payment
- Hiển thị trạng thái giao dịch

Kiến trúc:

Presentation (UI)
        ↓
POS App (Application Layer)
        ↓
Domain Layer (F&B / Retail)
        ↓
Business Support (Order Payment)
        ↓
Integration Layer

---

## 3. Trách nhiệm (Responsibilities)

### ✔ Nhận thao tác người dùng

- Scan barcode
- Chọn món
- Thêm vào giỏ hàng
- Chỉnh số lượng
- Áp dụng giảm giá

---

### ✔ Điều phối nghiệp vụ

Ví dụ F&B:

- createOrder()
- addItem()
- confirmOrder()

Ví dụ Retail:

- createSale()
- checkStock()
- reserveInventory()

---

### ✔ Kích hoạt thanh toán

- Gọi Order Payment Domain
- Hiển thị trạng thái thanh toán
- Nhận kết quả

---

### ✔ Quản lý phiên bán hàng (Sale Session)

- Mở ca
- Đóng ca
- Tổng kết cuối ca

---

### ✔ In hóa đơn

- Gửi lệnh tới printer adapter
- Hiển thị bản xem trước hóa đơn

---

## 4. Hành vi chính (Core Behaviors)

### 4.1 Start Sale
- Tạo order mới (F&B hoặc Retail)
- Load cấu hình unit

---

### 4.2 Add Item
addItem(productId, quantity)


- Gọi Domain để validate
- Cập nhật UI

---

### 4.3 Apply Discount
applyDiscount(code)


- Gọi domain validate
- Cập nhật tổng tiền

---

### 4.4 Checkout
checkout(paymentMethod)


Flow:

1. Validate order
2. Gọi Order Payment
3. Nhận kết quả
4. Complete order
5. In hóa đơn

---

### 4.5 Open / Close Shift
openShift()
closeShift()


- Lưu tiền đầu ca
- Tính tổng cuối ca

---

## 5. POS KHÔNG làm gì

- Không tính toán tồn kho
- Không quyết định business rule
- Không quản lý subscription
- Không kiểm soát permission (IAM làm)
- Không gọi trực tiếp payment gateway (thông qua Business Support)

---

## 6. Phân biệt POS vs Domain

| POS App | Domain |
|----------|--------|
| Điều phối | Chứa luật |
| Gọi use case | Thực thi business rule |
| Hiển thị UI | Kiểm tra invariant |
| Không lưu logic ngành | Lưu logic ngành |

---

## 7. POS trong F&B vs Retail

### POS trong F&B

- Quản lý bàn
- Tách bill
- Gộp bill
- Gửi xuống bếp

---

### POS trong Retail

- Scan barcode
- Kiểm tra tồn kho
- Trả hàng
- Đổi hàng

---

## 8. POS nâng cao (Offline Mode)

Nếu hỗ trợ offline:

- Lưu local order
- Đồng bộ khi online
- Resolve conflict

Lúc này POS có thể có Local Sync Service.

---

## 9. Boundary quan trọng

POS chỉ hoạt động trong phạm vi:

- unitId
- organizationId
- saleSession

POS không được phép:

- Truy cập unit khác
- Truy cập organization khác

---

## 10. Tóm tắt

POS App là:

- Ứng dụng bán hàng tại điểm bán
- Thuộc Application Layer
- Điều phối nghiệp vụ
- Không chứa business rule lõi
- Có thể dùng cho nhiều vertical

POS là cửa vào của business,
Domain mới là bộ não xử lý logic.
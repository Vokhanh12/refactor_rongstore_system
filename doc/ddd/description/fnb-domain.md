# F&B Domain (Food & Beverage)

## 1. Overview

F&B Domain thuộc **Business Vertical Layer**.

Đây là domain chứa toàn bộ logic nghiệp vụ ngành Food & Beverage.

F&B chỉ hoạt động bên trong một Unit.

Organization / Unit thuộc Platform Core.
F&B là business chạy phía trên Platform Core.

---

## 2. Phạm vi của F&B Domain

F&B chịu trách nhiệm:

- Menu
- Product
- Table
- Order
- Kitchen workflow
- Pricing
- Discount
- Tax
- Order lifecycle

F&B không xử lý:

- IAM
- Subscription
- Billing SaaS
- Payment Gateway integration (chỉ gọi qua Business Support)

---

## 3. Aggregate chính trong F&B

### 3.1 Menu

Thuộc tính:

- id
- unitId
- name
- status (ACTIVE / INACTIVE)
- availableTimeRange

---

### 3.2 Product

- id
- menuId
- name
- price
- category
- options (size, topping…)
- status

---

### 3.3 Table (nếu là dine-in)

- id
- unitId
- tableNumber
- status (AVAILABLE / OCCUPIED / RESERVED)

---

### 3.4 Order (Aggregate Root quan trọng nhất)

- id
- unitId
- tableId (nullable)
- orderType (DINE_IN / TAKE_AWAY / DELIVERY)
- status
- totalAmount
- createdAt

---

### 3.5 OrderItem

- productId
- quantity
- price
- note
- subtotal

---

## 4. Trách nhiệm của F&B Domain

### ✔ Quản lý Menu

- Tạo / sửa / xóa menu
- Bật / tắt sản phẩm

---

### ✔ Quản lý Order lifecycle

- Create order
- Add item
- Remove item
- Calculate total
- Confirm order
- Complete order
- Cancel order

---

### ✔ Tính toán giá

- Áp dụng discount
- Áp dụng tax
- Tính subtotal
- Tính service charge

---

### ✔ Quản lý trạng thái bàn

- Check-in bàn
- Close bàn

---

## 5. Hành vi chính (Core Behaviors)

### 5.1 Create Order
createOrder(unitId, orderType, tableId?)

- Validate unit ACTIVE
- Tạo order status = CREATED

---

### 5.2 Add Item
addItem(productId, quantity)

- Validate product active
- Tính subtotal
- Cập nhật total

---

### 5.3 Confirm Order
confirm()

- Lock order
- Gửi xuống bếp (emit event)

---

### 5.4 Complete Order
complete()


- Set status = COMPLETED

---

### 5.5 Cancel Order

cancel(reason)


- Set status = CANCELLED

---

## 6. Order State Machine (Quan trọng)

CREATED  
→ CONFIRMED  
→ PREPARING  
→ SERVED  
→ COMPLETED  

Hoặc  

CREATED  
→ CANCELLED  

F&B chịu trách nhiệm quản lý state machine này.

---

## 7. Quan hệ với các Layer khác

| Layer | Vai trò |
|-------|----------|
| Platform Core | Cung cấp Organization, Unit |
| Business Support | Gọi Order Payment |
| Integration Layer | Gửi notification, in bill |
| IAM | Kiểm tra permission |

---

## 8. Phân biệt F&B vs Order Payment

| F&B | Order Payment |
|-----|---------------|
| Tạo order | Thanh toán order |
| Tính tiền | Thực hiện giao dịch |
| Quản lý bàn | Gọi payment gateway |
| Quản lý trạng thái | Xử lý webhook |

F&B không xử lý tiền.
F&B chỉ biết totalAmount.

---

## 9. Boundary quan trọng

F&B:

- Không truy cập Organization khác
- Không truy cập Unit khác
- Không biết Subscription
- Không xử lý SaaS billing

F&B chỉ hoạt động trong phạm vi:

(unitId)

---

## 10. Flow thực tế

### Case: Khách quét QR tại bàn

1. User chọn Organization
2. Chọn Unit
3. F&B load menu
4. User tạo Order
5. F&B tính total
6. Gọi Order Payment
7. Payment thành công
8. F&B complete order

---

## 11. Vì sao F&B là Business Vertical

Vì:

- Chứa logic ngành cụ thể
- Retail sẽ có logic khác
- Service sẽ có logic khác

F&B không dùng lại cho Retail.

---

## 12. Tóm tắt

F&B Domain là:

- Logic ngành Food & Beverage
- Chứa Order lifecycle
- Quản lý menu, sản phẩm, bàn
- Tính tiền
- Không xử lý thanh toán thực tế
- Chạy trong phạm vi Unit

Không có F&B → hệ thống không có nghiệp vụ ngành.
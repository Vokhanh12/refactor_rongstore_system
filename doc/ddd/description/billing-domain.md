# Billing Domain

## 1. Overview

Billing Domain thuộc **Platform Core Layer**.

Billing chịu trách nhiệm:

- Tính phí
- Tạo hóa đơn (Invoice)
- Theo dõi thanh toán
- Ghi nhận giao dịch

Billing KHÔNG kiểm soát quyền sử dụng hệ thống.
Việc đó thuộc Subscription Domain.

---

## 2. Vai trò trong kiến trúc

Billing là domain xử lý tiền ở mức SaaS (thu tiền Organization).

Luồng tổng quát:

Subscription -> Billing tạo Invoice -> Thanh toán -> Cập nhật Subscription

Billing không biết business F&B hay Retail.
Billing chỉ biết:

- Organization nợ bao nhiêu tiền
- Đã thanh toán hay chưa

---

## 3. Entity bên trong Billing Context

### 3.1 Invoice (Aggregate Root)

Đại diện cho hóa đơn thanh toán.

Thuộc tính ví dụ:

- id
- organizationId
- amount
- currency
- status (PENDING / PAID / FAILED / CANCELLED)
- dueDate
- createdAt

---

### 3.2 InvoiceItem

Chi tiết dòng tiền trong hóa đơn.

Ví dụ:

- description
- quantity
- unitPrice
- totalAmount

---

### 3.3 PaymentTransaction

Giao dịch thanh toán thực tế.

- id
- invoiceId
- paymentMethod
- externalTransactionId
- status
- paidAt

---

## 4. Trách nhiệm (Responsibilities)

Billing Domain chịu trách nhiệm:

### ✔ Tạo hóa đơn

- Khi Subscription activate / renew
- Khi upgrade plan
- Khi mua add-on

---

### ✔ Theo dõi trạng thái thanh toán

- Pending
- Paid
- Failed

---

### ✔ Ghi nhận transaction

- Lưu externalTransactionId
- Lưu phương thức thanh toán

---

### ✔ Đồng bộ với Subscription

- Nếu invoice PAID -> kích hoạt subscription
- Nếu invoice FAILED -> giữ nguyên trạng thái

---

## 5. Hành vi chính (Core Behaviors)

### 5.1 Create Invoice
createInvoice(organizationId, items)

- Tính tổng tiền
- Set status = PENDING
- Lưu invoice

---

### 5.2 Mark as Paid
markAsPaid(transactionData)


- Update invoice status = PAID
- Lưu PaymentTransaction
- Emit domain event

---

### 5.3 Mark as Failed
markAsFailed(reason)


- Update status
- Ghi log lỗi

---

### 5.4 Calculate Amount
calculateTotal(items)


- Sum invoice items

---

## 6. Billing KHÔNG làm gì

- Không gọi trực tiếp Payment Gateway
- Không validate thẻ ngân hàng
- Không xử lý API Stripe / Momo

Những việc đó thuộc Integration Layer (Payment Gateway Adapter).

Billing chỉ xử lý logic nội bộ.

---

## 7. Quan hệ với Domain khác

| Domain | Quan hệ |
|--------|---------|
| Subscription | Billing tạo invoice cho subscription |
| Organization | Invoice thuộc Organization |
| Payment Gateway Adapter | Thực hiện thanh toán |
| IAM | Kiểm tra quyền trước khi thanh toán |

---

## 8. Flow thực tế

### Case: Organization gia hạn gói

1. Subscription yêu cầu gia hạn
2. Billing tạo Invoice
3. Frontend gọi Payment Gateway
4. Gateway trả webhook
5. Billing cập nhật invoice PAID
6. Subscription được activate

---

## 9. Billing thuộc Platform Core vì

- Không phụ thuộc ngành nghề
- Dùng chung cho mọi vertical
- Là cơ chế thu tiền SaaS

Nếu mai thêm Retail, Service → Billing không đổi.

---

## 10. Phân biệt Billing vs Payment Support Domain

| Billing | Order Payment (Business Support) |
|----------|----------------------------------|
| Thu tiền Organization | Thu tiền khách hàng |
| Liên quan Subscription | Liên quan Order |
| Hóa đơn SaaS | Thanh toán đơn hàng |
| Platform level | Business level |

---

## 11. Boundary quan trọng

Billing:

- Không truy cập dữ liệu business (Order, Menu)
- Không xử lý tiền mặt tại quầy
- Không xử lý loyalty

Chỉ xử lý tiền giữa Organization và Platform.

---

## 12. Tóm tắt

Billing Domain là:

- Hệ thống thu tiền SaaS
- Quản lý hóa đơn
- Ghi nhận giao dịch
- Đồng bộ với Subscription
- Không phụ thuộc ngành nghề

Không có Billing -> không thể vận hành SaaS thương mại.
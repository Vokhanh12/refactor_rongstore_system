# Member Domain

## 1. Overview

Member Domain thuộc **Platform Core Layer**.

Member đại diện cho **quan hệ giữa User và Organization trong một context cụ thể**.

User (IAM) ↔ Member ↔ Organization ↔ Unit


- **User** = danh tính toàn hệ thống (global identity)  
- **Member** = danh tính trong tenant (tenant identity)  

👉 Một User có thể có nhiều Member (mỗi Organization là 1 Member)

---

## 2. Mục đích của Member Domain

Giải quyết bài toán:

- Multi-tenant authorization  
- User thuộc nhiều Organization  
- Phân quyền theo Organization / Unit / Domain  

Ví dụ:

- User A:
   + OWNER ở Organization A
   + STAFF ở Organization B
   + MANAGER ở Unit X (thuộc Organization A)

---

## 3. Phân biệt User vs Member

| User (IAM) | Member |
|------------|--------|
| Global identity | Tenant identity |
| Email / Password | Role / Permission |
| Không biết Organization | Biết Organization / Unit |
| Auth (xác thực) | AuthZ (phân quyền) |

---

## 4. Aggregate Design

👉 **Member là Aggregate Root**
Member (Aggregate Root)
├── RoleAssignment
├── UnitScopedRole (optional)
└── Invitation (có thể tách riêng nếu cần scale)

---

## 5. Entity trong Member Context

### 5.1 Member (Aggregate Root)
id
userId
organizationId
status (ACTIVE / INVITED / SUSPENDED)
joinedAt


❗ Không nên đặt role trực tiếp tại đây.

---

### 5.2 RoleAssignment
id
memberId
scopeType (ORG / UNIT)
scopeId (organizationId / unitId)
role (OWNER / ADMIN / STAFF / MANAGER / ...)
domain (optional: F&B / HR / ...)


👉 Cho phép:

- 1 Member có nhiều role  
- Role theo Organization  
- Role theo Unit  
- Role theo domain  

---

### 5.3 UnitScopedRole (Optional)

- Có thể bỏ nếu RoleAssignment đã đủ  
- Dùng khi muốn optimize read model  

---

### 5.4 Invitation
id
email
organizationId
role (initial role)
token
expiredAt

---

## 6. Responsibilities

### ✔ Manage Membership

- Add member vào Organization  
- Remove member  
- Suspend / activate member  

---

### ✔ Manage Role & Permission

- Gán role theo Organization  
- Gán role theo Unit  
- Gán role theo domain  

---

### ✔ Authorization

- Kiểm tra quyền user  
- Scope theo Organization / Unit / Domain  

---

### ✔ Invitation Flow

- Tạo invitation  
- Accept / reject  

---

## 7. Core Behaviors

### 7.1 Add Member
addMember(userId, organizationId)

- Validate Organization  
- Check subscription limit  
- Tạo Member  

---

### 7.2 Assign Role
assignRole(memberId, scopeType, scopeId, role, domain)

- Validate scope  
- Validate permission  
- Tạo RoleAssignment  

---

### 7.3 Remove Member
removeMember(memberId)

- Không xóa owner cuối cùng  
- Soft delete  

---

### 7.4 Authorize
canPerform(userId, orgId, unitId, action, domain)


- Lookup Member  
- Lookup RoleAssignment  
- Check permission  

---

## 8. Member KHÔNG làm gì

- Không authenticate user (IAM làm)  
- Không quản lý Organization lifecycle  
- Không xử lý Subscription billing  
- Không chứa logic business (F&B, HR)  

---

## 9. Quan hệ với Domain khác

| Domain | Quan hệ |
|--------|--------|
| IAM | dùng userId |
| Organization | Member thuộc Organization |
| Unit | RoleAssignment scoped theo Unit |
| Subscription | giới hạn số member |
| Business Modules | check permission qua Member |

---

## 10. Boundary quan trọng

- Không cross Organization  
- Role luôn có scope rõ ràng  
- Không truy cập data ngoài tenant  

👉 Member = security boundary của tenant  

---

## 11. Insight quan trọng

### 1. Role không nằm trong Member

Role phải là entity riêng (RoleAssignment)

---

### 2. Member không thuộc Unit

Member thuộc Organization  
Unit chỉ là scope của Role  

---

### 3. Không có sub-root

- Member là Aggregate Root  
- Unit là Aggregate Root  
- Liên kết bằng ID  

---

## 12. Tóm tắt

Member Domain:

- Cầu nối giữa User và Organization  
- Quản lý Authorization (AuthZ)  
- Quản lý Role theo Organization / Unit / Domain  
- Là nền tảng multi-tenant SaaS  

---

## 🔥 Một câu chốt

**Member là trung tâm của Authorization trong hệ thống multi-tenant**
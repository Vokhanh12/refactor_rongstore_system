# F&B Role & Permission Example Data

## 1️⃣ USERS

| ID            | username       | full_name           |
|---------------|---------------|-------------------|
| user1-uuid    | admin         | Super Admin        |
| user2-uuid    | rest_manager  | Restaurant Manager |
| user3-uuid    | street_vendor | Street Vendor      |
| user4-uuid    | cafe_operator | Café Operator      |

---

## 2️⃣ ROLES

| ID            | CODE               | NAME                | TYPE           | DESCRIPTION                  |
|---------------|------------------|-------------------|----------------|-----------------------------|
| role1-uuid    | SUPER_ADMIN       | Super Admin       | PLATFORM       | Full access to system       |
| role2-uuid    | RESTAURANT_MANAGER| Restaurant Manager| ORGANIZATION   | Manage a restaurant unit    |
| role3-uuid    | CAFE_OPERATOR     | Café Operator     | ORGANIZATION   | Operate a café              |
| role4-uuid    | FB_VENDOR         | Street Vendor     | ORGANIZATION   | Operate a mobile vendor     |

---

## 3️⃣ PERMISSIONS

| ID            | CODE             | RESOURCE    | ACTION     | NAME                  |
|---------------|----------------|------------|------------|----------------------|
| perm1-uuid    | ORDER_CREATE    | ORDER      | CREATE     | Create Order          |
| perm2-uuid    | ORDER_UPDATE    | ORDER      | UPDATE     | Update Order          |
| perm3-uuid    | ORDER_COMPLETE  | ORDER      | COMPLETE   | Complete Order        |
| perm4-uuid    | PRODUCT_MANAGE  | PRODUCT    | MANAGE     | Manage Product        |
| perm5-uuid    | PAYMENT_COLLECT | PAYMENT    | COLLECT    | Collect Payment       |
| perm6-uuid    | POST_CREATE     | POST       | CREATE     | Create Post           |

---

## 4️⃣ ROLE_PERMISSIONS

| ROLE_ID       | PERMISSION_ID   |
|---------------|----------------|
| role1-uuid    | perm1-uuid     |
| role1-uuid    | perm2-uuid     |
| role1-uuid    | perm3-uuid     |
| role1-uuid    | perm4-uuid     |
| role1-uuid    | perm5-uuid     |
| role1-uuid    | perm6-uuid     |
| role2-uuid    | perm1-uuid     |
| role2-uuid    | perm2-uuid     |
| role2-uuid    | perm3-uuid     |
| role2-uuid    | perm4-uuid     |
| role2-uuid    | perm5-uuid     |
| role3-uuid    | perm1-uuid     |
| role3-uuid    | perm2-uuid     |
| role3-uuid    | perm3-uuid     |
| role3-uuid    | perm4-uuid     |
| role3-uuid    | perm5-uuid     |
| role4-uuid    | perm1-uuid     |
| role4-uuid    | perm3-uuid     |
| role4-uuid    | perm4-uuid     |
| role4-uuid    | perm5-uuid     |

---

## 5️⃣ ROLE_ASSIGNEMENTS

| USER_ID       | ROLE_ID         | SCOPE_TYPE | SCOPE_ID           |
|---------------|----------------|------------|------------------|
| user1-uuid    | role1-uuid      | GLOBAL     | NULL             |
| user2-uuid    | role2-uuid      | UNIT       | unit-uuid-123    |
| user3-uuid    | role4-uuid      | OWN        | vendor-unit-uuid |
| user4-uuid    | role3-uuid      | UNIT       | cafe-uuid-456    |

---

### 🔹 Giải thích nhanh

- `SUPER_ADMIN` (user1) → toàn hệ thống (GLOBAL)  
- `Restaurant Manager` (user2) → quyền trên 1 unit nhà hàng cụ thể  
- `Street Vendor` (user3) → chỉ quyền của chính user (OWN)  
- `Café Operator` (user4) → quyền ở quán café cụ thể (UNIT)
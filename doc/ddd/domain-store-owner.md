# Store Owner Domain

## Aggregate Root

StoreOwner

---

## Responsibility

StoreOwner đại diện cho **chủ cửa hàng gắn với vị trí địa lý** trong hệ thống.

Aggregate này chịu trách nhiệm:

* đảm bảo vị trí (geo + tile) luôn hợp lệ
* quản lý vòng đời dữ liệu của store owner
* cung cấp dữ liệu đúng cho việc tìm kiếm theo tile (query side)

StoreOwner **không chịu trách nhiệm search**, chỉ đảm bảo dữ liệu đúng.
---

## Key Attributes
- StoreOwnerId
- Location (lat, lng)
- TileIndex (tileX, tileY)
- LogoUrl (optional)
- CreatedBy


---

## Invariants (Luật bất biến)

Những luật này **luôn luôn đúng** trong suốt vòng đời StoreOwner:

* `latitude ∈ [-90, 90]`
* `longitude ∈ [-180, 180]`
* `tileX ≥ 0`
* `tileY ≥ 0`
* StoreOwner luôn có vị trí hợp lệ

Nếu bất kỳ invariant nào bị vi phạm → StoreOwner **không hợp lệ và không được persist**.

---

## Behaviors (Hành vi nghiệp vụ)

StoreOwner **không cho phép chỉnh field trực tiếp**.
Mọi thay đổi phải thông qua hành vi.

### Danh sách hành vi

* `Create()`
* `ChangeLocation()`
* `UpdateLogo()`
* `ChangeOwner()` *(nếu domain cho phép)*
* `Deactivate()` *(optional – soft delete / lifecycle)*

---

## Behavior Rules

### Create

* Latitude / Longitude phải hợp lệ
* Tile phải không âm
* `CreateBy` bắt buộc
* Khởi tạo `CreateDate` và `UpdateDate`

---

### ChangeLocation

* Vị trí mới phải hợp lệ (lat/lng)
* Tile mới phải không âm
* Cập nhật `UpdateDate`
* Không được phá invariant hiện tại

---

### UpdateLogo

* Logo URL có thể null
* Không ảnh hưởng invariant vị trí

---

### (Optional) Deactivate

* StoreOwner bị deactivate không được phép update location
* Dùng cho soft delete hoặc lifecycle control

---

## Validation Strategy

| Layer           | Trách nhiệm                              |
| --------------- | ---------------------------------------- |
| Proto           | Validate format & range cơ bản           |
| Domain (Entity) | Enforce invariant & behavior rules       |
| Repository      | Tin entity, không validate business      |
| Database        | Enforce data integrity (CHECK, NOT NULL) |
| Query side      | Validate tile range theo zoom            |

---

## Database Responsibility

Database **không quyết định business rule**, chỉ bảo vệ dữ liệu khỏi bị hỏng.

### CHECK constraints (ví dụ)

```sql
CHECK (LAT BETWEEN -90 AND 90)
CHECK (LNG BETWEEN -180 AND 180)
CHECK (TILE_X >= 0)
CHECK (TILE_Y >= 0)
```

---

(End)

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

## Notes

* Tile **không biết zoom** trong domain
* Zoom & tile upper bound là **query concern**, không phải invariant
* SearchByTile là **Query Use Case**, không phải behavior của StoreOwner

---

## Example Behavior (Code)

```go
func (s *StoreOwner) ChangeLocation(
    lat, lng float64,
    tileX, tileY int32,
    updatedBy string,
) error {
    if lat < -90 || lat > 90 {
        return ErrInvalidLatitude
    }
    if lng < -180 || lng > 180 {
        return ErrInvalidLongitude
    }
    if tileX < 0 || tileY < 0 {
        return ErrInvalidTile
    }

    s.Lat = lat
    s.Lng = lng
    s.TileX = tileX
    s.TileY = tileY
    s.UpdateBy = updatedBy
    s.UpdateDate = time.Now()

    return nil
}
```

---

## Design Principles

* Entity mang **hành vi**, không chỉ là data
* Repository chỉ làm persistence
* Domain không phụ thuộc DB / ORM
* Query và Command tách biệt (CQRS)

---

> **StoreOwner là đối tượng sống trong domain, không phải bản ghi DB**

(End)

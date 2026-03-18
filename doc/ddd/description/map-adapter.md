# Map Adapter Documentation

## 1. Mục đích

Map Adapter là một thành phần thuộc **Integration Layer**.

Nhiệm vụ chính:
- Trừu tượng hóa (abstract) việc làm việc với các dịch vụ bản đồ bên ngoài.
- Tách Domain logic khỏi provider cụ thể (Google Maps, Mapbox, OSRM, OpenStreetMap…).
- Cho phép thay đổi provider mà không ảnh hưởng đến Business / Domain.

Map Adapter KHÔNG chứa business logic.

---

## 2. Vai trò trong hệ thống

Map Adapter đóng vai trò:

- Chuyển đổi request từ Domain sang format của Map Provider
- Gọi API bên ngoài
- Chuẩn hóa response về format nội bộ của hệ thống

Nó là một "cổng giao tiếp" giữa hệ thống và dịch vụ bản đồ.

---

## 3. Map Adapter KHÔNG làm gì

- Không lưu dữ liệu cửa hàng
- Không quyết định business rule
- Không validate nghiệp vụ
- Không chứa logic Organization / Unit

---

## 4. Ví dụ sử dụng

### 4.1 Tạo Unit có vị trí

Flow:

1. User nhập địa chỉ
2. Organization/Unit Service gọi Map Adapter
3. Map Adapter gọi Geocoding API
4. Map Adapter trả về lat/lng
5. Domain lưu vào Unit

Pseudo flow:
UnitService.createUnit(name, address):
coordinates = mapAdapter.geocode(address)
unit = Unit(name, address, coordinates)
unitRepository.save(unit)

---

### 4.2 Tìm cửa hàng gần nhất

Flow:

1. User gửi vị trí hiện tại
2. Domain lấy danh sách unit
3. Có thể gọi Map Adapter để:
   - Tính khoảng cách
   - Hoặc gọi routing service (OSRM)

---

## 5. OSRM Adapter là gì?

OSRM Adapter là một loại Map Adapter chuyên dùng để:

- Tính đường đi (routing)
- Tính khoảng cách thực tế
- Tính thời gian di chuyển

Ví dụ:
distance = osrmAdapter.calculateRoute(from, to)

OSRM Adapter vẫn thuộc Integration Layer.

---

## 6. Thiết kế Interface

Nên định nghĩa interface chung:

```java
public interface MapAdapter {
    Coordinates geocode(String address);
    Route calculateRoute(Coordinates from, Coordinates to);
}
```

## 7. Lợi ích kiến trúc

✔ Không phụ thuộc provider
✔ Dễ thay đổi dịch vụ
✔ Test dễ dàng (mock adapter)
✔ Domain sạch sẽ

## 8. Tóm tắt

Thành phần	Vai trò
Organization	Quản lý công ty
Unit	Quản lý chi nhánh
Map Adapter	Lấy và xử lý dữ liệu bản đồ
OSRM Adapter	Tính đường đi

## 9. Kết luận

Map Adapter là công cụ hỗ trợ kỹ thuật (technical tool) cho Domain.

Domain quyết định:

Cần tọa độ để làm gì

Map Adapter chỉ:

Cung cấp tọa độ

Hoặc tính toán bản đồ theo yêu cầu

Không chứa business logic.
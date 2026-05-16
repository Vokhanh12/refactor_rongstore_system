# Architecture Mindset — Layer Responsibility & Error Classification

Updated: 2026-05-16

---

# 1. Core Mindset

Kiến trúc không phải là:

> chia folder đẹp

Mà là:

> phân chia responsibility đúng boundary

Một hệ thống maintainable cần:

* business logic không phụ thuộc technical concern
* transport không leak vào domain
* infra không điều khiển business
* parser không nằm trong validator domain
* application không chứa persistence detail

---

# 2. Layer Classification

| Layer       | Responsibility                     | Ví dụ                  |
| ----------- | ---------------------------------- | ---------------------- |
| DOMAIN      | Business invariant & business rule | Role name required     |
| APPLICATION | Usecase orchestration              | Create role workflow   |
| TRANSPORT   | Request/DTO/protobuf/http/grpc     | Parse request body     |
| INFRA       | DB/network/cache/external service  | Postgres timeout       |
| PARSER      | Technical conversion/parsing       | Parse UUID             |
| SECURITY    | Authentication/authorization       | JWT invalid            |
| SYSTEM      | Unexpected/internal/fallback       | Panic/internal failure |

---

# 3. DOMAIN Layer

## Purpose

Domain tồn tại để bảo vệ business invariant.

Domain KHÔNG nên biết:

* database
* http
* grpc
* json
* protobuf
* redis
* sqlc
* parser detail
* request payload

---

## DOMAIN should contain

### Entity

```go
Role
Permission
User
```

### Value Object

```go
Email
PhoneNumber
RoleKey
```

### Domain Validator

```go
Required
MinLen
Pattern
Enum
Range
```

---

## DOMAIN Validation = Business Meaning

Ví dụ:

```go
roleCode cannot be empty
```

đây là invariant.

---

## DOMAIN should NOT contain

❌ split string

❌ parse uuid

❌ decode json

❌ db duplicate detection

❌ transport DTO validation

❌ request malformed handling

---

# 4. APPLICATION Layer

## Purpose

Điều phối usecase.

Application layer xử lý:

* orchestration
* transaction flow
* repository interaction
* policy coordination
* existence check
* conflict handling

---

## APPLICATION should contain

### Usecase

```go
CreateRoleUsecase
AssignPermissionUsecase
```

### Application Validation

```go
Duplicate
Conflict
NotFound
AlreadyExists
DependencyMissing
ForbiddenState
```

---

## Example

```go
exists, err := repo.Exists(ctx, roleKey)

v := validation.New().
    AlreadyExists("role", exists)

if err := v.Err(); err != nil {
    return nil, err
}
```

Đây KHÔNG phải domain invariant.

Đây là orchestration concern.

---

# 5. TRANSPORT Layer

## Purpose

Nhận request từ bên ngoài và convert sang application input.

---

## TRANSPORT should contain

### grpc/http handler

```go
AuthzHandler
```

### DTO / Request object

```go
CreateRoleRequest
```

### Request validation

```go
missing request field
invalid payload structure
protobuf malformed
```

---

## Example

```go
if req == nil {
    return nil, errInvalidPayload
}
```

Đây không phải domain.

---

# 6. INFRA Layer

## Purpose

Kết nối external system.

---

## INFRA should contain

### Database

```go
postgres
sqlc
gorm
```

### Cache

```go
redis
```

### External services

```go
jwt
s3
rabbitmq
```

---

## INFRA Error Examples

```go
DB_TIMEOUT
POSTGRES_UNAVAILABLE
REDIS_UNAVAILABLE
JSON_SERIALIZATION_FAILED
```

---

# 7. PARSER Layer

## Purpose

Technical conversion.

Parser KHÔNG phải domain.

---

## PARSER should contain

### Parsing

```go
ParseUUID
ParseRoleKey
ParseJSON
ParseEnum
```

---

## Example

```go
func ParseRoleKey(value string) (RoleKey, error)
```

### Technical concern

```go
strings.Split()
uuid.Parse()
json.Unmarshal()
```

---

## PARSER Error Examples

```go
INVALID_UUID
INVALID_ROLE_KEY_FORMAT
JSON_PARSE_FAILED
```

---

# 8. SECURITY Layer

## Purpose

Authentication & authorization concern.

---

## SECURITY should contain

```go
JWT validation
Permission denied
Token expired
Unauthorized
```

---

## Examples

```go
TOKEN_EXPIRED
JWT_INVALID
UNAUTHORIZED
ROLE_FORBIDDEN
```

---

# 9. SYSTEM Layer

## Purpose

Unexpected fallback.

---

## SYSTEM Examples

```go
panic
unknown internal failure
unmapped exception
```

---

# 10. Validation Philosophy

## Important Principle

Không phải mọi validate đều là domain validation.

---

## Domain Validation

```go
email required
role name max length
percentage range
```

---

## Technical Validation

```go
invalid uuid
json malformed
invalid protobuf
```

---

## Application Validation

```go
already exists
dependency missing
conflict
```

---

# 11. Refactor Validator Strategy

## BEFORE

```text
validator/
 ├── Required
 ├── MinLen
 ├── ParseUUID
 ├── Conflict
 ├── NotFound
 ├── Duplicate
```

Problem:

* mixed responsibility
* leaking boundary
* difficult maintenance
* domain polluted by infra concern

---

## AFTER

```text
core/
├── domain/
│   └── validator/
│       ├── required.go
│       ├── pattern.go
│       └── range.go
│
├── application/
│   └── validation/
│       ├── conflict.go
│       ├── existence.go
│       └── dependency.go
│
├── parser/
│   ├── uuid.go
│   ├── role_key.go
│   └── json.go
│
├── infra/
│   ├── postgres/
│   ├── redis/
│   └── jwt/
│
└── errors/
```

---

# 12. Error Classification Mindset

## DOMAIN error

Business invalid.

```text
ROLE_NAME_REQUIRED
EMAIL_INVALID
```

---

## APPLICATION error

Workflow invalid.

```text
ROLE_ALREADY_EXISTS
DEPENDENCY_MISSING
```

---

## PARSER error

Technical conversion failed.

```text
INVALID_UUID
INVALID_ROLE_KEY_FORMAT
```

---

## INFRA error

External system failed.

```text
DB_TIMEOUT
REDIS_UNAVAILABLE
```

---

## SECURITY error

Authentication/authorization failure.

```text
TOKEN_EXPIRED
PERMISSION_DENIED
```

---

# 13. ParseRoleKey Correct Mindset

## WRONG

```go
validator.Format(...)
```

Vì parsing != domain invariant.

---

## CORRECT

```go
parts := strings.Split(value, ":")

if len(parts) != 2 {
    return errInvalidRoleKeyFormat
}
```

---

# 14. Repository Error Translation

Repository layer nên translate infra error sang app error boundary.

---

## Example

```go
if err != nil {
    return dberr.TranslateDBError(err, s.dberr)
}
```

Đây là đúng direction.

Infrastructure detail không leak ra ngoài.

---

# 15. Final Architecture Philosophy

## Domain protects meaning

## Application coordinates workflow

## Transport receives external request

## Parser converts technical representation

## Infra talks to external systems

## Security protects access

## System handles unexpected failure

---

# 16. Golden Rule

Nếu logic cần business knowledge → DOMAIN

Nếu logic cần orchestration → APPLICATION

Nếu logic cần technical conversion → PARSER

Nếu logic cần external system → INFRA

Nếu logic cần request/response → TRANSPORT

Nếu logic cần access control → SECURITY

Nếu logic là fallback/unexpected → SYSTEM

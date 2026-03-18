Tổng quan Use Case Map
IAM
├── Identity  → Quản lý user & trạng thái bảo mật
├── Auth      → Xác thực & token
└── AuthZ     → Phân quyền & policy

1️⃣ Identity Context – Use cases cốt lõi

Identity = nguồn sự thật của User

🧑‍💻 User Lifecycle
1. CreateUser

Tạo user mới (email / phone)

Set password ban đầu hoặc external provider

Init security state

👉 dùng khi:

Signup

Admin create user

2. ActivateUser / DeactivateUser

Disable user

Soft delete

Block tạm thời

3. UpdateIdentifiers

Add / remove / change:

Email

Phone number

⚠️ invariant:

Không trùng identifier

Ít nhất 1 identifier active

🔐 Credential & Security
4. ChangePassword

Verify password cũ

Update password mới

Update passwordChangedAt

5. RequestForgotPassword

Generate OTP

Set expiredAt

Send OTP (email/SMS)

6. VerifyForgotPasswordOTP

Verify code

Check expired

Check retry count

7. ResetPassword

Reset password bằng OTP

Clear OTP

Unlock user (nếu cần)

8. RecordLoginSuccess

Reset failed count

Update lastLoginAt, IP

9. RecordLoginFailure

Increase failed count

Auto lock nếu vượt threshold

10. LockUser / UnlockUser

Lock manual hoặc auto

Unlock sau thời gian

🔍 Query (Read)
11. FindUserByIdentifier

email / phone → User snapshot

2️⃣ Auth Context – Use cases xác thực

Auth = orchestrator, không sở hữu User

🔑 Authentication
1. Login (CORE)

Resolve identity

Verify credential

Get authorization

Issue access token

2. LoginWithOTP

Verify OTP

Skip password

Issue token

3. LoginWithSocial

Google / Facebook / Apple

Map external identity → User

4. RefreshToken

Verify refresh token

Issue new access token

5. Logout

Revoke refresh token

Blacklist access token (optional)

6. VerifyAccessToken

Validate signature

Check expiration

Parse claims

7. IntrospectToken (optional)

Used by API Gateway / Envoy

Active / inactive

🧾 Token Management
8. IssueAccessToken

Embed:

sub

roles

permissions

scopes

9. RevokeToken

Manual revoke

Admin force logout

3️⃣ AuthZ Context – Use cases phân quyền

AuthZ = policy & authorization rule

🎭 Role Management
1. CreateRole

System role

Business role

2. UpdateRole

Name

Description

Active / inactive

3. DeleteRole

Soft delete

Cannot delete system role

🔑 Permission Management
4. CreatePermission

resource + action

Example: order:read

5. UpdatePermission
6. DeletePermission
🔗 Role ↔ Permission
7. AssignPermissionToRole

Many-to-many

8. RemovePermissionFromRole
👤 User ↔ Role
9. AssignRoleToUser

With scope (global / tenant / org)

10. RevokeRoleFromUser
11. CheckUserPermission

Input:

userID

resource

action

Output: allow / deny

👉 thường dùng bởi:

API Gateway

Policy Engine

12. GetAuthorizationSnapshot (CORE)

userID →

roles

permissions
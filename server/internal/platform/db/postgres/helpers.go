package postgres

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func timestamptzFromTime(t time.Time) pgtype.Timestamptz {
	var ts pgtype.Timestamptz
	ts.Scan(t)
	return ts
}

func timeFromTimestamptz(ts pgtype.Timestamptz) time.Time {
	if ts.Valid {
		return ts.Time
	}
	return time.Time{}
}

func numericFromFloat64(f float64) pgtype.Numeric {
	var n pgtype.Numeric
	// Convert float64 to string first, then scan
	n.Scan(fmt.Sprintf("%.2f", f))
	return n
}

func float64FromNumeric(n pgtype.Numeric) float64 {
	f, _ := n.Float64Value()
	return f.Float64
}

func TimestamptzFromTime(t time.Time) pgtype.Timestamptz {
	var ts pgtype.Timestamptz
	_ = ts.Scan(t)
	return ts
}

func TimeFromTimestamptz(ts pgtype.Timestamptz) time.Time {
	if ts.Valid {
		return ts.Time
	}
	return time.Time{}
}

func TimePtrFromTimestamptz(ts pgtype.Timestamptz) *time.Time {
	if !ts.Valid {
		return nil
	}
	return &ts.Time
}

func NumericFromFloat64(f float64) pgtype.Numeric {
	var n pgtype.Numeric
	_ = n.Scan(fmt.Sprintf("%.2f", f))
	return n
}

func Float64FromNumeric(n pgtype.Numeric) float64 {
	f, _ := n.Float64Value()
	return f.Float64
}

func Float64PtrFromNumeric(n pgtype.Numeric) *float64 {
	if !n.Valid {
		return nil
	}
	f, _ := n.Float64Value()
	return &f.Float64
}

func StringFromUUID(u pgtype.UUID) string {
	if !u.Valid {
		return ""
	}
	return u.String()
}

func StringPtrFromUUID(u pgtype.UUID) *string {
	if !u.Valid {
		return nil
	}
	s := u.String()
	return &s
}

func PgUUIDFromUUIDPtr(u *uuid.UUID) pgtype.UUID {
	if u == nil {
		return pgtype.UUID{Valid: false}
	}

	return pgtype.UUID{
		Bytes: *u,
		Valid: true,
	}
}

func UUIDPtrFromPgUUID(u pgtype.UUID) *uuid.UUID {
	if !u.Valid {
		return nil
	}

	id := uuid.UUID(u.Bytes)
	return &id
}

func StringFromText(t pgtype.Text) string {
	if !t.Valid {
		return ""
	}
	return t.String
}

func StringPtrFromText(t pgtype.Text) *string {
	if !t.Valid {
		return nil
	}
	return &t.String
}

func BoolFromBool(b pgtype.Bool) bool {
	if !b.Valid {
		return false
	}
	return b.Bool
}

func BoolPtrFromBool(b pgtype.Bool) *bool {
	if !b.Valid {
		return nil
	}
	return &b.Bool
}

func Int32FromInt(i pgtype.Int4) int32 {
	if !i.Valid {
		return 0
	}
	return i.Int32
}

func Int32PtrFromInt(i pgtype.Int4) *int32 {
	if !i.Valid {
		return nil
	}
	return &i.Int32
}

func Int64FromInt(i pgtype.Int8) int64 {
	if !i.Valid {
		return 0
	}
	return i.Int64
}

func Int64PtrFromInt(i pgtype.Int8) *int64 {
	if !i.Valid {
		return nil
	}
	return &i.Int64
}

func TextFromStringPtr(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{
		String: *s,
		Valid:  true,
	}
}

func Int4FromInt32(i int32) pgtype.Int4 {
	return pgtype.Int4{
		Int32: i,
		Valid: true,
	}
}

func PgInt4FromUint8(i uint8) pgtype.Int4 {
	return pgtype.Int4{
		Int32: int32(i),
		Valid: true,
	}
}

func Uint8FromPgInt4(i pgtype.Int4) uint8 {
	if !i.Valid {
		return 0
	}

	if i.Int32 < 0 {
		return 0
	}

	if i.Int32 > 255 {
		return 255
	}

	return uint8(i.Int32)
}

func Uint8FromInt32(i int32) uint8 {
	if i < 0 {
		return 0
	}

	if i > 255 {
		return 255
	}

	return uint8(i)
}

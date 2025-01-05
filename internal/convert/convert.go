package convert

import (
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// MustStringToInt64 converts a string value to int64 values,
// it causes panic if s is not a valid int64.
func MustStringToInt64(s string) int64 {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return num
}

// PgTimestampToTime converts a pgtype.Timestamptz value to *time.Time value,
// it returns nil if ts is not a valid pgtype.Timestamptz value.
func PgTimestampToTime(ts pgtype.Timestamptz) *time.Time {
	if ts.Valid {
		return &ts.Time
	}

	return nil
}

// MustPgTimestampToTime converts a pgtype.Timestamptz value to *time.Time value,
// it causes panic if ts is not a valid pgtype.Timestamptz value.
func MustPgTimestampToTime(ts pgtype.Timestamptz) time.Time {
	if !ts.Valid {
		panic("nullable time")
	}

	return ts.Time
}

// TimeToPgTimestamp converts a time.Time instance to pgtype.Timestamptz
func TimeToPgTimestamp(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:  t,
		Valid: true,
	}
}

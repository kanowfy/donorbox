package convert

import (
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func MustStringToInt64(s string) int64 {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return num
}

func PgTimestampToTime(ts pgtype.Timestamptz) *time.Time {
	if ts.Valid {
		return &ts.Time
	}

	return nil
}

func MustPgTimestampToTime(ts pgtype.Timestamptz) time.Time {
	if !ts.Valid {
		panic("nullable time")
	}

	return ts.Time
}

func TimeToPgTimestamp(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:  t,
		Valid: true,
	}
}

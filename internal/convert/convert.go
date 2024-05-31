package convert

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func StringToPgxUUID(s string) (pgtype.UUID, error) {
	var uuid pgtype.UUID
	err := uuid.Scan(s)
	if err != nil {
		return pgtype.UUID{}, err
	}
	return uuid, nil
}

func PgxUUIDToString(uuid pgtype.UUID) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid.Bytes[0:4], uuid.Bytes[4:6], uuid.Bytes[6:8], uuid.Bytes[8:10], uuid.Bytes[10:16])
}

func MustStringToPgxUUID(s string) pgtype.UUID {
	uuid, err := StringToPgxUUID(s)
	if err != nil {
		panic("parsing validated invalid UUID")
	}

	return uuid
}

func MustStringToInt64(s string) int64 {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return num
}

func MustTimeToPgxTimestamp(t time.Time) pgtype.Timestamptz {
	var time pgtype.Timestamptz
	err := time.Scan(t)
	if err != nil {
		panic("parsing validated invalid timestamp")
	}

	return time
}

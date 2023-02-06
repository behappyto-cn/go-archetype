package dateUtil

import (
	"time"
)

const (
	formatYyyyMMdd = "2006-01-02"
)

// 时间格式化
func FormatYyyyMmDd(date time.Time) string {
	return date.Format(formatYyyyMMdd)
}

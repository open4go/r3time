package r3time

import (
	"fmt"
	"time"
)

// CurrentTime 当前时间
// 格式 2006.01.02 15:04:05
func CurrentTime() string {
	dt := time.Now()
	t := fmt.Sprintf("%v", dt.Format(time.DateTime))
	return t
}

// Today 获取当天日期
// 格式 20060102
func Today() string {
	// Using time.Now() function.
	dt := time.Now()
	t := fmt.Sprintf("%v", dt.Format("20060102"))
	return t
}

// Timestamp2ReadTime 获取可读时间
// 格式 2006.01.02 15:04:05
func Timestamp2ReadTime(timestamp int64) string {
	// Convert timestamp to time.Time
	dt := time.Unix(timestamp, 0)
	t := fmt.Sprintf("%v", dt.Format(time.DateTime))
	return t
}

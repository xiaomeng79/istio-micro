package utils

import (
	"strconv"
	"time"
)

// 生成毫秒时间戳字符串
func GenMicTimeStr() string {
	return strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
}

// 毫秒时间戳转字符串
func MicTimeToFormatStr() string {
	tm := time.Now()
	return tm.Format("2006-01-02 15:04:05")
}

// 生成毫秒时间戳
func GenMicTime() int64 {
	return time.Now().UnixNano() / 1e6
}

// 毫秒时间戳转字符串
func MicTimeToStr(i64 int64) string {
	tm := time.Unix(i64/1e3, 0)
	return tm.Format("2006-01-02 15:04:05")
}

// 毫秒时间戳转字符串
func StrMicTimeToStr(s string) (string, error) {
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return "", err
	}
	tm := time.Unix(i64/1e3, 0)
	return tm.Format("2006-01-02 15:04:05"), nil
}

// 字符串转毫秒时间戳
func StrToMicTime(s string) string {
	// 获取本地location   // 待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                    // 转化所需模板
	loc, _ := time.LoadLocation("Local")                   // 重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, s, loc) // 使用模板在对应时区转化为time.time类型
	return strconv.FormatInt(theTime.Unix()*1e3, 10)
}

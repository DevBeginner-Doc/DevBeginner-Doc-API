package utils

import (
	"sort"
	"time"
)

func In(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	if index < len(str_array) && str_array[index] == target {
		return true
	}
	return false
}

func UnixToTime(unix int64) time.Time {
	temp := time.Unix(unix, 0)
	return temp
}

func StringToTime(timeStr string) time.Time {
	temp, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	return temp
}

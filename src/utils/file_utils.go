package utils

import (
	"strconv"
)

func CalculateSize(size int64) string {
	switch {
	case size < 1024:
		return strconv.FormatInt(size, 10) + "B"
	case size < 1024*1024:
		return strconv.FormatInt(size/1024, 10) + "KB"
	case size < 1024*1024*1024:
		return strconv.FormatInt(size/1024/1024, 10) + "MB"
	case size < 1024*1024*1024*1024:
		return strconv.FormatInt(size/1024/1024/1024, 10) + "GB"
	case size < 1024*1024*1024*1024*1024:
		return strconv.FormatInt(size/1024/1024/1024/1024, 10) + "TB"
	default:
		return "Too big"
	}
}

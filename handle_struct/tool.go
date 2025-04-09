package handlestruct

import (
	"strconv"
	"strings"
	"time"
)

// 时间转化
func dateToStr(t time.Time) string {
	layout := "2006-01-02"
	format := t.Format(layout)
	split := strings.Split(format, " ")
	if split[0] == "0001-01-01" {
		return ""
	}
	return split[0]
}

func timeToStr(t time.Time) string {
	layout := "2006-01-02 15:04:05"
	return t.Format(layout)
}

func strToDate(t string) time.Time {
	layout := "2006-01-02"
	if t == "" {
		return time.Time{}
	}
	parse, err := time.Parse(layout, t)
	if err != nil {
		return time.Time{}
	}
	return parse
}

func strToTime(t string) time.Time {
	layout := "2006-01-02 15:04:05"
	parse, err := time.Parse(layout, t)
	if err != nil {
		return time.Time{}
	}
	return parse
}

func strInt(s string) int {
	if s == "" {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
func strInt32(s string) int32 {
	if s == "" {
		return 0
	}
	si, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0
	}
	return int32(si)
}

func strFloat64(s string) float64 {
	if s == "" {
		return 0
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

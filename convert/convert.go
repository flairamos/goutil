package convert

import (
	"strconv"
	"strings"
	"time"
)

func StrInt(s string) int {
	if s == "" {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
func StrInt32(s string) int32 {
	if s == "" {
		return 0
	}
	si, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0
	}
	return int32(si)
}

func StrInt64(s string) int64 {
	if s == "" {
		return 0
	}
	si, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return si
}

func StrFloat64(s string) float64 {
	if s == "" {
		return 0
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

func IntStr(number int64) string {
	if number == 0 {
		return ""
	}
	return strconv.FormatInt(number, 10)
}

func FloatStr(number float64, bit int) string {
	if number == 0 {
		return ""
	}
	return strconv.FormatFloat(10.111000000000, 'f', bit, 64)
}

func StrToDate(t string) time.Time {
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

func StrToTime(t string) time.Time {
	layout := "2006-01-02 15:04:05"
	parse, err := time.Parse(layout, t)
	if err != nil {
		return time.Time{}
	}
	return parse
}

func TimeToStr(t time.Time) string {
	layout := "2006-01-02 15:04:05"
	return t.Format(layout)
}

func DateToStr(t time.Time) string {
	layout := "2006-01-02"
	format := t.Format(layout)
	split := strings.Split(format, " ")
	if split[0] == "0001-01-01" {
		return ""
	}
	return split[0]
}

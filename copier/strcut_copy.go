package copier

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
	"time"

	. "github.com/flairamos/goutil/convert"
)

// 结构体的赋值方法，适配大多数基本数据类型，不可以结构体的嵌套
// The assignment method of structs is suitable for most basic data types, and cannot be nested by structs
func StructCopy[T any, M any](param T) M {
	var result M
	modelType := reflect.TypeOf(param)
	valueType := reflect.ValueOf(param)
	resultType := reflect.TypeOf(result)
	resultPtr := reflect.New(resultType).Elem()
	// 依次赋值
	// 遍历结构体字段
	for i := 0; i < modelType.NumField(); i++ {
		// 在新建的结构体参数寻找一样的字段赋值
		for j := 0; j < resultType.NumField(); j++ {
			if modelType.Field(i).Name == resultType.Field(j).Name {
				// 判断类型是否一样
				// 类型不一样将参数类型转为返回值类型
				if modelType.Field(i).Type.Name() == resultType.Field(j).Type.Name() {
					resultPtr.Field(j).Set(reflect.ValueOf(valueType.Field(i).Interface()))
				} else {
					if slices.Contains([]string{"int", "int64", "int32", "int16", "int8"}, resultType.Field(j).Type.Name()) {
						// 数字类型
						if slices.Contains([]string{"int", "int64", "int32", "int16", "int8"}, modelType.Field(i).Type.Name()) {
							resultPtr.Field(j).Set(reflect.ValueOf(valueType.Field(i).Interface()))
						} else if slices.Contains([]string{"uint", "uint64", "uint32", "uint16", "uint8"}, modelType.Field(i).Type.Name()) {
							value := valueType.Field(i).Interface()
							resultPtr.Field(j).SetInt(StrInt64(fmt.Sprintf("%d", value)))
						} else if slices.Contains([]string{"string"}, modelType.Field(i).Type.Name()) {
							resultPtr.Field(j).SetInt(StrInt64(valueType.Field(i).Interface().(string)))
						} else if slices.Contains([]string{"bool"}, modelType.Field(i).Type.Name()) {
							flag := valueType.Field(i).Interface().(bool)
							if flag {
								resultPtr.Field(j).SetInt(1)
							} else {
								resultPtr.Field(j).SetInt(0)
							}
						} else { // 其他类型无法转换
							resultPtr.Field(j).SetInt(0)
						}
					} else if slices.Contains([]string{"float32, float64"}, resultType.Field(j).Type.Name()) {
						if slices.Contains([]string{"int", "int64", "int32", "int16", "int8"}, modelType.Field(i).Type.Name()) {
							value := fmt.Sprintf("%d", valueType.Field(i).Interface())
							resultPtr.Field(j).SetFloat(StrFloat64(value))
						} else if slices.Contains([]string{"string"}, modelType.Field(i).Type.Name()) {
							resultPtr.Field(j).SetFloat(float64(StrInt32(valueType.Field(i).Interface().(string))))
						} else { // 其他类型无法转换
							resultPtr.Field(j).SetFloat(0)
						}
					} else if slices.Contains([]string{"uint", "uint64", "uint32", "uint16", "uint8"}, resultType.Field(j).Type.Name()) {
						if slices.Contains([]string{"int", "int64", "int32", "int16", "int8"}, modelType.Field(i).Type.Name()) {
							valueStr := fmt.Sprintf("%d", valueType.Field(i).Interface())
							if strings.HasPrefix(valueStr, "-") {
								resultPtr.Field(j).SetUint(0)
							} else {
								resultPtr.Field(j).SetUint(uint64(StrInt32(valueStr)))
							}
						} else if slices.Contains([]string{"string"}, modelType.Field(i).Type.Name()) {
							// string先转int判断正负
							value := valueType.Field(i).Interface().(string)
							if StrInt32(value) < 0 {
								resultPtr.Field(j).SetUint(uint64(-StrInt32(value)))
							} else {
								resultPtr.Field(j).SetUint(uint64(StrInt32(value)))
							}
						} else if slices.Contains([]string{"uint", "uint64", "uint32", "uint16", "uint8"}, modelType.Field(i).Type.Name()) {
							resultPtr.Field(j).SetUint(valueType.Field(i).Interface().(uint64))
						} else { // 其他类型无法转换
							resultPtr.Field(j).SetUint(0)
						}
					} else if slices.Contains([]string{"bool"}, resultType.Field(j).Type.Name()) {
						if slices.Contains([]string{"bool"}, modelType.Field(i).Type.Name()) {
							resultPtr.Field(j).SetBool(valueType.Field(i).Interface().(bool))
						} else if slices.Contains([]string{"string"}, modelType.Field(i).Type.Name()) {
							resultPtr.Field(j).SetBool(StrBool(valueType.Field(i).Interface().(string)))
						} else if slices.Contains([]string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8"}, modelType.Field(i).Type.Name()) {
							var varible = make(map[string]any)
							varible["one"] = 1
							if valueType.Field(i).Interface() == varible["one"] {
								resultPtr.Field(j).SetBool(true)
							} else {
								resultPtr.Field(j).SetBool(false)
							}
						} else { // 其他类型无法转换
							resultPtr.Field(j).SetBool(false)
						}
					} else if slices.Contains([]string{"string"}, resultType.Field(j).Type.Name()) {
						if slices.Contains([]string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8"}, modelType.Field(i).Type.Name()) {
							resultPtr.Field(j).SetString(fmt.Sprintf("%v", valueType.Field(i).Interface()))
						} else if slices.Contains([]string{"bool"}, modelType.Field(i).Type.Name()) {
							resultPtr.Field(j).SetString(fmt.Sprintf("%t", valueType.Field(i).Interface().(bool)))
						} else if slices.Contains([]string{"Time"}, modelType.Field(i).Type.Name()) {
							t := valueType.Field(i).Interface().(time.Time)
							if t.IsZero() {
								resultPtr.Field(i).SetString("")
							} else if strings.Contains(t.String()[11:20], "00:00:00") {
								resultPtr.Field(i).SetString(TimeStandardDate(t))
							} else {
								resultPtr.Field(i).SetString(TimeStandardStr(t))
							}
						} else if slices.Contains([]string{"float32, float64"}, modelType.Field(i).Type.Name()) {
							resultPtr.Field(j).SetString(fmt.Sprintf("%.3f", valueType.Field(i).Interface().(float64)))
						} else {
							resultPtr.Field(j).SetString(fmt.Sprintf("%v", valueType.Field(i).Interface()))
						}
					} else if slices.Contains([]string{"Time"}, resultType.Field(j).Type.Name()) {
						if modelType.Field(i).Type.Name() == "Time" {
							resultPtr.Field(j).Set(reflect.ValueOf(valueType.Field(i).Interface().(time.Time)))
						} else {
							resultPtr.Field(j).Set(reflect.ValueOf(time.Time{}))
						}
					} else {
						continue
					}
				}
			}
		}
	}
	return resultPtr.Interface().(M)
}
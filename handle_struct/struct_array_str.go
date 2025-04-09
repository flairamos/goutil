package handlestruct

import (
	"reflect"
	"time"
)

type ReflectModel struct {
	Name     string
	TypeName string
	Value    any
}

// 泛型传入一个结构体返回成员名称，类型和值的结构体数组
func ReflectInfo[T any](model T) []ReflectModel {
	var result = make([]ReflectModel, 0)
	getType := reflect.TypeOf(model)
	for i := 0; i < getType.NumField(); i++ {
		fieldType := getType.Field(i)
		name := fieldType.Name
		typeName := fieldType.Type.String()
		value := reflect.ValueOf(model).Field(i)
		if typeName == "time.Time" {
			result = append(result, ReflectModel{name, typeName, timeToStr(time.Now())})
		} else {
			result = append(result, ReflectModel{name, typeName, value.Interface()})
		}
	}
	return result
}

// 将字符串数组转化为一个结构体
// 字符串数组内是具体的值，值与实际的结构体类型可以转换
// 方法根据结构体的值将字符串数组的次序值转化
// 字符串数组的值必须与结构体的成员总数一致
func MakeStrcutByArray[T any](p []string) T {
	var param T
	info := ReflectInfo(param)
	structType := reflect.TypeOf(param)
	// 使用反射创建结构体实例
	structValue := reflect.New(structType).Elem()
	for i, model := range info {
		// 设置结构体字段的值
		field1 := structValue.FieldByName(model.Name)
		if field1.IsValid() && field1.CanSet() {
			if model.TypeName == "time.Time" {
				// 判断是日期还是时间
				// 注意时间的格式只支持`2024-10-10` 和`2024-10-10 12:30:00`两种。
				if len(p[i]) == 10 {
					field1.Set(reflect.ValueOf(strToDate(p[i])))
				} else if len(p[i]) == 19 {
					field1.Set(reflect.ValueOf(strToTime(p[i])))
				} else {
					field1.Set(reflect.ValueOf(strToTime("")))
				}
			} else if model.TypeName == "int" {
				field1.Set(reflect.ValueOf(strInt(p[i])))
			} else if model.TypeName == "int32" {
				field1.Set(reflect.ValueOf(strInt32(p[i])))
			} else if model.TypeName == "float32" || model.TypeName == "float64" {
				field1.Set(reflect.ValueOf(strFloat64(p[i])))
			} else {
				field1.Set(reflect.ValueOf(p[i]))
			}
		}
	}
	return structValue.Interface().(T)
}

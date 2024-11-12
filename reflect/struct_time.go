package reflect

import (
	"reflect"
	"strings"
	"time"
)

// 结构体时间转换方法
// 对于存在time.Time类型的结构体来说，对于时间的转换比较麻烦，类似 2024-10-11 12:10:30 与time.Time的转化，
// 或者  2024-10-11 与time.Time的转化需要区分。
// 使用如下的方法可避免上述的区分与转化，只需要两个结构体，一个是存在time.Time的结构体，另一个是string类型的时间
// 使用StrcutTimeCopy方法传递两个泛型，一个time.Time类型一个string时间的结构体，方法会自动区分日期与时间。
// T是具有time.Time的结构体 ， M是具有string的结构体
func StrcutTimeConvert[T any, M any](param T) M {
	var result M
	modelType := reflect.TypeOf(param)
	valueType := reflect.ValueOf(param)
	resultType := reflect.TypeOf(result)
	resultPtr := reflect.New(resultType).Elem()

	for i := 0; i < modelType.NumField(); i++ {
		for j := 0; j < resultType.NumField(); j++ {
			if modelType.Field(i).Name == resultType.Field(i).Name {
				if modelType.Field(i).Type.Name() == "Time" {
					t := valueType.Field(i).Interface().(time.Time)
					if strings.Contains(t.String()[11:20], "00:00:00") {
						resultPtr.Field(i).SetString(dateToStr(t))
					} else {
						resultPtr.Field(i).SetString(timeToStr(t))
					}
				} else {
					resultPtr.Field(i).Set(reflect.ValueOf(valueType.Field(i).Interface()))
				}
				continue
			}
		}
	}
	return resultPtr.Interface().(M)
}

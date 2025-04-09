package handlestruct

import (
	"reflect"
	"time"
)

// StructToMap converts a struct to a map while omitting empty fields and formatting time.Time fields.
// which can omit default value with each type
// the map key name is the property of the struct
func StructToMapOmitZero(input interface{}) map[string]interface{} {
    result := make(map[string]interface{})
    v := reflect.ValueOf(input)
    t := v.Type()

    for i := 0; i < v.NumField(); i++ {
        field := v.Field(i)
        fieldType := t.Field(i)

        if !field.IsZero() { // Check if the field is not empty
            if field.Type() == reflect.TypeOf(time.Time{}) {
                result[fieldType.Name] = field.Interface().(time.Time).Format("2006-01-02 15:04:05")
            } else {
                result[fieldType.Name] = field.Interface()
            }
        }
    }
    return result
}


// transfer the grom tag name to the map key
// omit the boolean type value
func StructToMapNoBool(input interface{}) map[string]interface{} {
    result := make(map[string]interface{})
    v := reflect.ValueOf(input)

    origin := reflect.TypeOf(input)

    for i := 0; i < v.NumField(); i++ {
        field := v.Field(i)
        originField := origin.Field(i)
        tag := originField.Tag.Get("gorm")
        if field.Kind() == reflect.Bool {
            result[tag] = field.Interface()
        } else if field.Type() == reflect.TypeOf(time.Time{}) && !field.IsZero() {
            result[tag] = field.Interface().(time.Time).Format("2006-01-02 15:04:05")
        } else if !field.IsZero() {
            result[tag] = field.Interface()
        } else {

        }
    }
    return result
}


func StructToMapNoBoolCustomTag(input interface{}, tagName string) map[string]interface{} {
    result := make(map[string]interface{})
    v := reflect.ValueOf(input)

    origin := reflect.TypeOf(input)

    for i := 0; i < v.NumField(); i++ {
        field := v.Field(i)
        originField := origin.Field(i)
        tag := originField.Tag.Get(tagName)
        if field.Kind() == reflect.Bool {
            result[tag] = field.Interface()
        } else if field.Type() == reflect.TypeOf(time.Time{}) && !field.IsZero() {
            result[tag] = field.Interface().(time.Time).Format("2006-01-02 15:04:05")
        } else if !field.IsZero() {
            result[tag] = field.Interface()
        } else {
            // do nothing
        }
    }
    return result
}


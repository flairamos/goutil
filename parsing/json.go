package parsing

import (
	"encoding/json"
	"fmt"
)

// Json2List
// @params jsonData: json string
// @description Parse `[1,1,1]` to []int{1,1,1}  the parameter can be any or interface{}
// @author xvwen
// @return []T
func JsonToList[T interface{}](jsonData string) (*[]T, error) {
	var data []T
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Json2Struct 解析json字符串为结构体
// @auther xvwen
// @param jsonData json字符串 `{"name","xvwen","blog":"www.baidu.com"}`
// @desc 解析json字符串为结构体
// @return *T 结构体指针, error
func JsonToStruct[T interface{}](jsonData string) (*T, error) {
	// 定义一个变量用于存储解析后的数据
	var items T
	// 使用json.Unmarshal进行解析
	err := json.Unmarshal([]byte(jsonData), &items)
	if err != nil {
		fmt.Println("解析JSON时发生错误:", err)
		return nil, err
	}
	return &items, err
}

// Json2StructList 解析json字符串为结构体切片
// @auther xvwen
// @param jsonData json字符串 `{"name","xvwen","blog":"www.baidu.com"}`
// @desc 解析json字符串为结构体切片
// @return *[]T 结构体切片指针, error
func JsonToStructList[T interface{}](jsonData string) (*[]T, error) {
	// 定义一个切片用于存储解析后的数据
	var items []T
	// 使用json.Unmarshal进行解析
	err := json.Unmarshal([]byte(jsonData), &items)
	if err != nil {
		fmt.Println("解析JSON时发生错误:", err)
		return nil, err
	}
	return &items, err
}

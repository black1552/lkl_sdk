package common

import (
	"reflect"

	"github.com/gogf/gf/v2/encoding/gjson"
)

// CleanJSON 清理JSON字符串中的空值和0值字段
func CleanJSON(jsonStr string) ([]byte, error) {
	// 解析为JSON
	json, err := gjson.DecodeToJson(jsonStr, gjson.Options{StrNumber: true})
	if err != nil {
		return nil, err
	}

	// 递归清理数据
	cleaned := cleanData(json)

	// 转换回JSON字符串
	result, err := gjson.Marshal(cleaned)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// cleanData 递归清理map或slice中的空值和0值
func cleanData(data interface{}) interface{} {
	if data == nil {
		return nil
	}

	v := reflect.ValueOf(data)

	switch v.Kind() {
	case reflect.Map:
		cleanedMap := make(map[string]interface{})
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key)
			if val.Interface() != nil && !isEmptyValue(val) {
				cleanedVal := cleanData(val.Interface())
				// 只有清理后的值不为空才添加
				if !isCleanedValueEmpty(cleanedVal) {
					cleanedMap[key.String()] = cleanedVal
				}
			}
		}
		return cleanedMap

	case reflect.Slice:
		cleanedSlice := make([]interface{}, 0)
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i)
			if item.Interface() != nil && !isEmptyValue(item) {
				cleanedItem := cleanData(item.Interface())
				if !isCleanedValueEmpty(cleanedItem) {
					cleanedSlice = append(cleanedSlice, cleanedItem)
				}
			}
		}
		// 如果slice为空，则返回nil而不是空slice
		if len(cleanedSlice) == 0 {
			return nil
		}
		return cleanedSlice

	default:
		return data
	}
}

// isEmptyValue 检查值是否为空值或0值
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Slice, reflect.Map:
		return v.Len() == 0
	case reflect.String:
		return v.String() == "" || v.String() == "0"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	}

	// 特殊处理nil值
	if !v.IsValid() {
		return true
	}

	return false
}

// isCleanedValueEmpty 检查清理后的值是否为空
func isCleanedValueEmpty(val interface{}) bool {
	if val == nil {
		return true
	}

	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Map:
		return v.Len() == 0
	case reflect.Slice:
		return v.Len() == 0
	case reflect.String:
		return val == ""
	}

	return false
}

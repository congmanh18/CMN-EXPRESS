package gorm

import "reflect"

func OmitFields(
	data interface{},
	callBack func(fieldValue reflect.Value) bool,
) []string {
	var omitFields []string

	// Lấy giá trị thực tế, giải tham chiếu nếu là con trỏ
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Đảm bảo rằng v là struct trước khi xử lý
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			fieldName := v.Type().Field(i).Name
			fieldValue := v.Field(i)

			if callBack(fieldValue) {
				omitFields = append(omitFields, fieldName)
			}
		}
	}

	return omitFields
}

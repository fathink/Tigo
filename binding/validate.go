package binding

import (
	"reflect"
	"regexp"
	"strconv"
)

// checkObjBinding 检查实例的字段
func checkObjBinding(element reflect.Type, vElement reflect.Value) error {
	for i := 0; i < element.NumField(); i++ {
		isRequired := false
		fieldType := element.Field(i).Type.Name()
		fieldValue := vElement.Field(i).Interface()
		fieldName := element.Field(i).Name
		required, isRequiredExisted := element.Field(i).Tag.Lookup("required")
		if !isRequiredExisted || (required != "true" && required != "TRUE") {
			continue
		} else {
			isRequired = true
		}
		defaultValue, isDefaultExisted := element.Field(i).Tag.Lookup("default")
		regexStr, isRegexExisted := element.Field(i).Tag.Lookup("regex")
		switch fieldType {
		case "bool":
			break
		case "string":
			value := fieldValue.(string)
			if isRequired && value == "" && (!isDefaultExisted || defaultValue == "") {
				return RequiredErr(fieldName)
			}
			if value == "" && defaultValue != "" {
				value = defaultValue
			}
			if isRegexExisted && !isMatchRegex(value, regexStr) {
				return RegexErr(fieldName)
			}
			vElement.Field(i).SetString(value)
			break
		case "int":
			value := fieldValue.(int)
			if isRequired && value == 0 && (!isDefaultExisted || defaultValue == "") {
				return RequiredErr(fieldName)
			}
			if value == 0 && defaultValue != "" {
				var err error
				value, err = strconv.Atoi(defaultValue)
				if err != nil {
					return DefaultErr(fieldName)
				}
			}
			if isRegexExisted {
				valueStr := strconv.FormatInt(int64(value), 10)
				if !isMatchRegex(valueStr, regexStr) {
					return RegexErr(fieldName)
				}
			}
			vElement.Field(i).SetInt(int64(value))
			break
		case "int8":
			value := fieldValue.(int8)
			if isRequired && value == 0 && (!isDefaultExisted || defaultValue == "") {
				return RequiredErr(fieldName)
			}
			if value == 0 && defaultValue != "" {
				val, err := strconv.Atoi(defaultValue)
				if err != nil {
					return DefaultErr(fieldName)
				}
				value = int8(val)
			}
			if isRegexExisted {
				valueStr := strconv.FormatInt(int64(value), 10)
				if !isMatchRegex(valueStr, regexStr) {
					return RegexErr(fieldName)
				}
			}
			vElement.Field(i).SetInt(int64(value))
			break
		case "int16":
			value := fieldValue.(int16)
			if isRequired && value == 0 && (!isDefaultExisted || defaultValue == "") {
				return RequiredErr(fieldName)
			}
			if value == 0 && defaultValue != "" {
				val, err := strconv.Atoi(defaultValue)
				if err != nil {
					return DefaultErr(fieldName)
				}
				value = int16(val)
			}
			if isRegexExisted {
				valueStr := strconv.FormatInt(int64(value), 10)
				if !isMatchRegex(valueStr, regexStr) {
					return RegexErr(fieldName)
				}
			}
			vElement.Field(i).SetInt(int64(value))
			break
		case "int32":
			value := fieldValue.(int32)
			if isRequired && value == 0 && (!isDefaultExisted || defaultValue == "") {
				return RequiredErr(fieldName)
			}
			if value == 0 && defaultValue != "" {
				val, err := strconv.Atoi(defaultValue)
				if err != nil {
					return DefaultErr(fieldName)
				}
				value = int32(val)
			}
			if isRegexExisted {
				valueStr := strconv.FormatInt(int64(value), 10)
				if !isMatchRegex(valueStr, regexStr) {
					return RegexErr(fieldName)
				}
			}
			vElement.Field(i).SetInt(int64(value))
			break
		case "int64":
			value := fieldValue.(int64)
			if isRequired && value == 0 && (!isDefaultExisted || defaultValue == "") {
				return RequiredErr(fieldName)
			}
			if value == 0 && defaultValue != "" {
				val, err := strconv.Atoi(defaultValue)
				if err != nil {
					return DefaultErr(fieldName)
				}
				value = int64(val)
			}
			if isRegexExisted {
				valueStr := strconv.FormatInt(int64(value), 10)
				if !isMatchRegex(valueStr, regexStr) {
					return RegexErr(fieldName)
				}
			}
			vElement.Field(i).SetInt(int64(value))
			break
		case "uint":
			value := fieldValue.(uint)
			if isRequired && value == 0 && (!isDefaultExisted || defaultValue == "") {
				return RequiredErr(fieldName)
			}
			if value == 0 && defaultValue != "" {
				val, err := strconv.Atoi(defaultValue)
				if err != nil {
					return DefaultErr(fieldName)
				}
				value = uint(val)
			}
			if isRegexExisted {
				valueStr := strconv.FormatUint(uint64(value), 10)
				if !isMatchRegex(valueStr, regexStr) {
					return RegexErr(fieldName)
				}
			}
			vElement.Field(i).SetUint(uint64(value))
			break
		case "uint8":
			value := fieldValue.(uint8)
			if isRequired && value == 0 && (!isDefaultExisted || defaultValue == "") {
				return RequiredErr(fieldName)
			}
			if value == 0 && defaultValue != "" {
				val, err := strconv.Atoi(defaultValue)
				if err != nil {
					return DefaultErr(fieldName)
				}
				value = uint8(val)
			}
			if isRegexExisted {
				valueStr := strconv.FormatUint(uint64(value), 10)
				if !isMatchRegex(valueStr, regexStr) {
					return RegexErr(fieldName)
				}
			}
			vElement.Field(i).SetUint(uint64(value))
			break
		case "uint16":
			value := fieldValue.(uint16)
			if isRequired && value == 0 && (!isDefaultExisted || defaultValue == "") {
				return RequiredErr(fieldName)
			}
			if value == 0 && defaultValue != "" {
				val, err := strconv.Atoi(defaultValue)
				if err != nil {
					return DefaultErr(fieldName)
				}
				value = uint16(val)
			}
			if isRegexExisted {
				valueStr := strconv.FormatUint(uint64(value), 10)
				if !isMatchRegex(valueStr, regexStr) {
					return RegexErr(fieldName)
				}
			}
			vElement.Field(i).SetUint(uint64(value))
			break
		case "uint32":
			value := fieldValue.(uint32)
			if isRequired && value == 0 && (!isDefaultExisted || defaultValue == "") {
				return RequiredErr(fieldName)
			}
			if value == 0 && defaultValue != "" {
				val, err := strconv.Atoi(defaultValue)
				if err != nil {
					return DefaultErr(fieldName)
				}
				value = uint32(val)
			}
			if isRegexExisted {
				valueStr := strconv.FormatUint(uint64(value), 10)
				if !isMatchRegex(valueStr, regexStr) {
					return RegexErr(fieldName)
				}
			}
			vElement.Field(i).SetUint(uint64(value))
			break
		case "uint64":
			value := fieldValue.(uint64)
			if isRequired && value == 0 && (!isDefaultExisted || defaultValue == "") {
				return RequiredErr(fieldName)
			}
			if value == 0 && defaultValue != "" {
				val, err := strconv.Atoi(defaultValue)
				if err != nil {
					return DefaultErr(fieldName)
				}
				value = uint64(val)
			}
			if isRegexExisted {
				valueStr := strconv.FormatUint(uint64(value), 10)
				if !isMatchRegex(valueStr, regexStr) {
					return RegexErr(fieldName)
				}
			}
			vElement.Field(i).SetUint(uint64(value))
			break
		case "float32":
			value := fieldValue.(float32)
			if isRequired && value == 0 && (!isDefaultExisted || defaultValue == "") {
				return RequiredErr(fieldName)
			}
			if value == 0 && defaultValue != "" {
				val, err := strconv.ParseFloat(defaultValue, 64)
				if err != nil {
					return DefaultErr(fieldName)
				}
				value = float32(val)
			}
			if isRegexExisted {
				valueStr := strconv.FormatFloat(float64(value), 'E', -1, 32)
				if !isMatchRegex(valueStr, regexStr) {
					return RegexErr(fieldName)
				}
			}
			vElement.Field(i).SetFloat(float64(value))
			break
		case "float64":
			value := fieldValue.(float64)
			if isRequired && value == 0 && (!isDefaultExisted || defaultValue == "") {
				return RequiredErr(fieldName)
			}
			if value == 0 && defaultValue != "" {
				val, err := strconv.ParseFloat(defaultValue, 64)
				if err != nil {
					return DefaultErr(fieldName)
				}
				value = float64(val)
			}
			if isRegexExisted {
				valueStr := strconv.FormatFloat(float64(value), 'E', -1, 32)
				if !isMatchRegex(valueStr, regexStr) {
					return RegexErr(fieldName)
				}
			}
			vElement.Field(i).SetFloat(float64(value))
			break
		default:
			field := vElement.Field(i)
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				if field.CanAddr() && field.Addr().CanInterface() {
					attrType := field.Type()
					attrValue := field
					return checkObjBinding(attrType, attrValue)
				}
			}
			if fieldKind == reflect.Ptr {
				if field.CanAddr() && field.Addr().CanInterface() {
					attrType := field.Type().Elem()
					attrValue := field.Elem()
					return checkObjBinding(attrType, attrValue)
				}
			}
		}
	}
	return nil
}

// ValidateInstance 检查结构体实例化是否有效
func ValidateInstance(obj interface{}) error {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	element := t.Elem()
	vElement := v.Elem()
	return checkObjBinding(element, vElement)
}

// 判断是否符合正则规则
func isMatchRegex(str, regex string) bool {
	rgx := regexp.MustCompile(regex)
	return rgx.MatchString(str)
}

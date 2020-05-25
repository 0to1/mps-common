package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

func Interface2String(value interface{}) string {
	var val string
	switch value.(type) {
	case string:
		val = value.(string)
	case int:
		val = strconv.Itoa(value.(int))
	case float64:
		val = strconv.FormatFloat(value.(float64), 'f', -1, 64)
	case bool:
		val = strconv.FormatBool(value.(bool))
	default:
		val = ""
	}

	return val
}

func GetInterfaceType(value interface{}) string {
	return reflect.TypeOf(value).String()
}

func String2Interface(value string, vtype string) (interface{}, error) {
	var val interface{}
	var err error

	switch vtype {
	case "string":
		val = value
	case "int":
		val, _ = strconv.Atoi(value)
	case "float64":
		val, err = strconv.ParseFloat(value, 64)
	case "bool":
		val, err = strconv.ParseBool(value)
	default:
		val = value
		err = fmt.Errorf("未知类型%s", vtype)
	}

	return val, err
}
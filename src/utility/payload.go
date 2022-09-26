package utility

import (
	"reflect"
	"strconv"
	"strings"
)

func ConstructPayload(keyword string, model interface{}) (string, []interface{}) {
	isInt := true
	isFirstField := true
	iKeyword, err := strconv.Atoi(keyword)
	if err != nil {
		isInt = false
	}
	isBool := false
	if keyword == "true" || keyword == "false" {
		isBool = true
	}
	var payload string
	numField := reflect.TypeOf(model).Elem().NumField()
	args := make([]interface{}, 0)
	for i := 0; i < numField; i++ {
		field := reflect.TypeOf(model).Elem().Field(i)
		jsonTag := field.Tag.Get("json")
		typeName := field.Type.Name()
		if !((typeName == "uint" && isInt) || typeName == "string" || (typeName == "bool" && isBool)) {
			continue
		}
		if strings.ContainsRune(jsonTag, '-') {
			continue
		}
		if jsonTag == "-" {
			continue
		}
		if !isFirstField {
			payload += " OR "
		} else {
			isFirstField = false
		}
		if typeName == "uint" {
			payload += strings.Replace(jsonTag, "-", "_", -1) + " = ?"
			args = append(args, iKeyword)
		} else if typeName == "bool" {
			payload += strings.Replace(jsonTag, "-", "_", -1) + " = ?"
			args = append(args, keyword)
		} else {
			payload += strings.Replace(jsonTag, "-", "_", -1) + " like ?"
			args = append(args, "%"+keyword+"%")
		}
	}
	return payload, args
}

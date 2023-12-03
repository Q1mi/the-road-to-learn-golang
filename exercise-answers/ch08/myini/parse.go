package myini

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

/*
编写代码利用反射实现一个ini文件的解析器程序。
*/

// Parse 解析ini文件
func Parse(data []byte, result interface{}) (err error) {
	lineSlice := strings.Split(string(data), "\n")
	// 判断用户传进来的result是否是一个结构体指针
	ptrInfo := reflect.TypeOf(result)
	structInfo := ptrInfo.Elem()
	if ptrInfo.Kind() != reflect.Ptr {
		err = errors.New("please pass into a struct ptr")
		return
	}
	if structInfo.Kind() != reflect.Struct {
		err = errors.New("please pass into a struct ptr")
		return
	}

	// ini配置文件解析
	var structFieldName string //节
	// 解析语法错误
	for index, line := range lineSlice {
		// 判断是否是注释
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if line[0] == ';' || line[0] == '#' {
			continue
		}
		if line[0] == '[' { // 按照节来处理
			structFieldName, err = parseSection(line, index+1, structInfo)
			if err != nil {
				return err
			}
			continue
		} else { // 按照属性处理
			err := parseItem(structFieldName, line, index+1, result)
			if err != nil {
				return err
			}
		}

	}
	return
}

func parseSection(line string, lineNo int, structInfo reflect.Type) (structFieldName string, err error) {
	// 判断 节 格式是否有效
	if len(line) <= 2 || line[len(line)-1] != ']' {
		err = fmt.Errorf("syntax error, invalid section:%s line:%d", line, lineNo)
		return
	}
	//解析到节的内容
	sectionName := strings.TrimSpace(line[1 : len(line)-1])
	if len(sectionName) == 0 {
		err = fmt.Errorf("syntax error, invalid section:%s line:%d", line, lineNo)
		return
	}
	// 根据节的名字去结构体中按照tag获取对应字段
	for i := 0; i < structInfo.NumField(); i++ {
		field := structInfo.Field(i)
		tagName := field.Tag.Get("ini")
		if sectionName == tagName {
			structFieldName = field.Name
			fmt.Println(sectionName)
		}
		continue
	}
	return
}

// 解析item
func parseItem(structFieldName string, line string, lineNo int, result interface{}) (err error) {
	index := strings.Index(line, "=")
	if index == -1 { // 没有等号说明配置文件格式有误
		err = fmt.Errorf("syntax error, line:%d", lineNo)
		return
	}
	key := strings.TrimSpace(line[:index])
	value := strings.TrimSpace(line[index+1:])
	if len(key) == 0 {
		err = fmt.Errorf("syntax errror, line:%d", lineNo)
		return
	}
	resultValue := reflect.ValueOf(result)
	// 从结构体中取到节对应的那个嵌入结构体
	sectionValue := resultValue.Elem().FieldByName(structFieldName)
	sectionType := sectionValue.Type()
	if sectionType.Kind() != reflect.Struct {
		err = fmt.Errorf("%s field must be a struct", structFieldName)
		return
	}
	keyFieldName := ""
	for i := 0; i < sectionType.NumField(); i++ {
		field := sectionType.Field(i)
		tagVal := field.Tag.Get("ini")
		if tagVal == key {
			keyFieldName = field.Name
			break
		}
	}
	if len(keyFieldName) == 0 {
		return
	}
	fieldValue := sectionValue.FieldByName(keyFieldName)
	switch fieldValue.Type().Kind() {
	case reflect.String:
		fieldValue.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, parseErr := strconv.ParseInt(value, 10, 64)
		if parseErr != nil {
			return parseErr
		}
		fieldValue.SetInt(intVal)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		intVal, parseErr := strconv.ParseUint(value, 10, 64)
		if parseErr != nil {
			return parseErr
		}
		fieldValue.SetUint(intVal)
	case reflect.Float32, reflect.Float64:
		floatVal, parseErr := strconv.ParseFloat(value, 64)
		if parseErr != nil {
			return parseErr
		}
		fieldValue.SetFloat(floatVal)
	case reflect.Bool:
		boolVal, parseErr := strconv.ParseBool(value)
		if parseErr != nil {
			return parseErr
		}
		fieldValue.SetBool(boolVal)
	default:
		err = fmt.Errorf("unsupport type:%v", fieldValue.Type().Kind())
		return
	}
	return
}

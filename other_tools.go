// Package go_toolkit coding=utf-8
// @Project : go-toolkit
// @Time    : 2023/12/20 15:21
// @Author  : chengxiang.luo
// @Email   : chengxiang.luo@foxmail.com
// @File    : other_tools.go
// @Software: GoLand
package go_toolkit

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetInterfaceToString(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case time.Time:
		t, _ := value.(time.Time)
		key = t.String()
		// 2022-11-23 11:29:07 +0800 CST  这类格式把尾巴去掉
		key = strings.Replace(key, " +0800 CST", "", 1)
		key = strings.Replace(key, " +0000 UTC", "", 1)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

func ReadRecordsFromCsv(csvFilePath string) [][]string {
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		panic(err)
	}
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			panic(err)
		}
	}(csvFile)

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	return records
}

func GetInterfaceToInt(v interface{}) int {
	var r int
	switch v.(type) {
	case uint:
		r = int(v.(uint))
		break
	case int8:
		r = int(v.(int8))
		break
	case uint8:
		r = int(v.(uint8))
		break
	case int16:
		r = int(v.(int16))
		break
	case uint16:
		r = int(v.(uint16))
		break
	case int32:
		r = int(v.(int32))
		break
	case uint32:
		r = int(v.(uint32))
		break
	case int64:
		r = int(v.(int64))
		break
	case uint64:
		r = int(v.(uint64))
		break
	case float32:
		r = int(v.(float32))
		break
	case float64:
		r = int(v.(float64))
		break
	case string:
		r, _ = strconv.Atoi(v.(string))
		if r == 0 && len(v.(string)) > 0 {
			f, _ := strconv.ParseFloat(v.(string), 64)
			r = int(f)
		}
		break
	case nil:
		r = 0
		break
	case json.Number:
		t3, _ := v.(json.Number).Int64()
		r = int(t3)
		break
	default:
		r = v.(int)
		break
	}
	return r
}

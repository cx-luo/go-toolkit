// Package go_toolkit coding=utf-8
// @Project : go-toolkit
// @Time    : 2023/12/20 15:19
// @Author  : chengxiang.luo
// @Email   : chengxiang.luo@foxmail.com
// @File    : convert_charset_to_utf8.go
// @Software: GoLand
package go_toolkit

import (
	"github.com/qiniu/iconv"
)

// ConvertCharsetToUtf8 convert other coding to utf-8
func ConvertCharsetToUtf8(s string, fromCharset string) string {
	converter, err := iconv.Open("utf-8", fromCharset)
	if err != nil {
		panic(err)
	}
	defer func(converter iconv.Iconv) {
		err := converter.Close()
		if err != nil {
			panic(err)
		}
	}(converter)

	output := converter.ConvString(s)
	return output
}

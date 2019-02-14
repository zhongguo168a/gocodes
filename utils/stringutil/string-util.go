// 字符串工具集合
package stringutil

import "strings"

// 按照大小写分组
func SplitUpper(target string) []string {
	chars := []byte(target)
	result := []string{}
	str := ""
	for i := 0; i < len(target); i++ {
		char := chars[i]
		if i == 0 {

		} else if char >= 65 && char <= 90 {
			result = append(result, str)
			str = ""
		}

		str += string(char)
	}

	result = append(result, str)
	return result
}

//SplitLast 获取字符串被"."分离后的字符串数组中
// ，最后一个字符串。
func SplitLast(str string, sep string) string {
	strArr := strings.Split(str, sep)
	if strArr != nil && len(strArr) > 0 {
		return strArr[len(strArr)-1]
	}

	return str
}

// LowerFirst 首字母小写
func LowerFirst(str string) string {
	first := str[:1]
	return strings.Replace(str, first, strings.ToLower(first), 1)
}

// UpperFirst 首字母大写
func UpperFirst(str string) string {
	if len(str) == 0 {
		return str
	}

	first := str[:1]

	return strings.Replace(str, first, strings.ToUpper(first), 1)
}

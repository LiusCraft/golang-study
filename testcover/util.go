package testcover

import "strings"

func toCamelCase(word string) string {
	// 将单词转为首字母大写，其余字母小写的形式
	camelCase := strings.Title(strings.ToLower(word))
	// 将空格或下划线替换为空字符串
	camelCase = strings.ReplaceAll(camelCase, " ", "")
	camelCase = strings.ReplaceAll(camelCase, "_", "")
	return camelCase
}

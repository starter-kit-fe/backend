package utils

import "regexp"

// 验证电子邮件格式的函数
func IsValidEmail(email string) bool {
	// 使用正则表达式验证电子邮件格式
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// MD5哈希(小写)
// 常用于存储密码或生成唯一的摘要值
// 其易受两个不同的输入生成相同的哈希值，不推荐用于安全敏感的场景。
func Md5Encode(data string) string {
	m := md5.New()
	m.Write([]byte(data))              //将数据转换为字节数组写入哈希对象
	tempStr := m.Sum(nil)              //获取哈希值字节数组
	return hex.EncodeToString(tempStr) //将字节数组转为十六进制字符串
}

// MD5哈希(大写)
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// 加盐哈希(加密)
func MakePassword(password, salt string) string {
	return Md5Encode(password + salt)
}

// 解密
func ValidPassword(password, salt string, passwordDB string) bool {
	return Md5Encode(password+salt) == passwordDB
}

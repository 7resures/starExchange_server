package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 检查文件后缀名是否符合要求
func IsSuffix(str string, list []string) bool {
	for _, item := range list {
		if str == item {
			return true
		}
	}
	return false
}

// 写入数据库时MD5加密
func MD5(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

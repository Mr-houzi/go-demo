package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// md5 一种摘要算法，经过md5之后可以得到一个128 bit的由0/1组成的二进制串。

// Encode 返回加密后的32字符；
//
// 通常，以字符串的形式返回它的16进制字符表示，即 Encode 方法；每4位2进制可以表示1位16进制，即转换后共32位。
func Encode(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Encode16 返回加密后的16字符；
//
// 通常，所谓的16字符的md5是去掉前8位、去掉后8位而得到。
func Encode16(str string) string {
	return Encode(str)[8:24]
}

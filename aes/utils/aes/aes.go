package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)
// AES/CBC/PKCS5Padding
func AesEncode(plaintext []byte, encryptionKey []byte, iv []byte) ([]byte, error) {
	//bPlaintext := PKCS5Padding(plaintext, aes.BlockSize)
	bPlaintext := PKCS5Padding(plaintext)
	block, err := aes.NewCipher(encryptionKey)

	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, bPlaintext)

	return ciphertext, nil
}

func AesDecode(cipherText []byte, encryptionKey []byte, iv []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	return PKCS7UnPadding(cipherText), nil
}

// PKCS7Padding [PKCS7 / PKCS5 填充算法](https://segmentfault.com/a/1190000019793040)
// PKCS7Padding [三种填充模式的区别(PKCS7Padding/PKCS5Padding/ZeroPadding)](https://blog.csdn.net/xiongya8888/article/details/84947232)
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	// 填充的长度
	paddingLen := blockSize - len(cipherText)%blockSize
	// 待填充的字符串
	paddingChar := []byte{byte(paddingLen)}
	padText := bytes.Repeat(paddingChar, paddingLen)

	return append(cipherText, padText...)
}

// PKCS5Padding pkcs5作为pkcs7的子集算法，概念上没有什么区别，只是在blockSize上固定为 8 bytes
func PKCS5Padding(cipherText []byte) []byte {
	return PKCS7Padding(cipherText, 8)
}

// PKCS7UnPadding PKCS7 和 PKCS5 去重填充是一个规则
func PKCS7UnPadding(plantText []byte) []byte {
	length   := len(plantText)
	// 填充的长度，最后一个字节肯定为填充数据的长度
	paddingLen := int(plantText[length-1])
	return plantText[:(length - paddingLen)]
}
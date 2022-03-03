package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"strings"
)

// https://gist.github.com/awadhwana/9c95377beba61293390c5fd23a3bb1df

func main() {
	EncryptTest("yey")
	EncryptTest("uwow")
	EncryptTest("hello world")
	EncryptTest("a c")

}

func EncryptTest(string2 string) {
	Key, _ := GenerateRandomBytes(32)
	encryptAES, IV, _ := Aes256Encode([]byte(string2), Key)
	decryptAES, _ := Aes256Decode(encryptAES, Key, IV)

	fmt.Println("Key:", Key)
	fmt.Println("Key len:", len(Key))
	fmt.Println("IV:", IV)
	fmt.Println("result: " + string(decryptAES))
	fmt.Println("result length: ", len(decryptAES))
	fmt.Println("---------")
}

func Aes256Encode(content []byte, encryptionKey []byte) (encryptedContent []byte, IV []byte, err error) {
	bPlaintext := PKCS5Padding(content, aes.BlockSize)
	block, err := aes.NewCipher(encryptionKey)

	if err != nil {
		return nil, nil, err
	}

	IV, _ = GenerateRandomBytes(block.BlockSize())
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, IV)
	mode.CryptBlocks(ciphertext, bPlaintext)

	return ciphertext, IV, err
}

func Aes256Decode(cipherText []byte, encryptionKey []byte, IV []byte) (decryptedContent []byte, err error) {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, IV)
	mode.CryptBlocks(cipherText, cipherText)

	cutTrailingSpaces := []byte(strings.TrimSpace(string(cipherText)))
	return cutTrailingSpaces, err
}

func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipherText, padText...)
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}


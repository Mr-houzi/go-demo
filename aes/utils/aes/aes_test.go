package aes

import (
	crand "crypto/rand"
	"encoding/hex"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// go test -run=TestAesEncode -v
func TestAesEncode(t *testing.T) {
	//// 生成密钥
	//key, err := GenerateRandomBytes(32)
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//// 生成偏移量
	//iv, err := GenerateRandomBytes(16)
	//if err != nil {
	//	t.Error(err)
	//}

	// 生成对称密钥
	key := GenrateRandomString(32)
	// 生成偏移量
	//iv := GenrateRandomString(16) // iv 的长度需要 16 个字节，16个字母字符组成的字符串等于16字节（即[]byte 长度位16）
	iv := GenrateRandomString(14) // iv 的长度需要 16 个字节，16个字母字符组成的字符串等于16字节（即[]byte 长度位16）

	// 待加密的明文
	plainText := "hello world"

	t.Log("plainText（明文）:", plainText)

	t.Log("key:", fmt.Sprintf("%v", key))
	t.Log("iv:", fmt.Sprintf("%v", iv))

	ciphertext, err := AesEncode([]byte(plainText), []byte(key), []byte(iv))
	if err != nil {
		t.Error(err)
	}

	t.Log("ciphertext byte:", ciphertext)
	t.Log("ciphertext:", string(ciphertext))
	t.Log("ciphertext hex(加密后的密文16进制表示):", fmt.Sprintf("%x", ciphertext))
}

func TestAesDecode(t *testing.T) {
	ciphertextHex := "ae5c4b54f33904d0dddcf2bbccdb2c02"
	ciphertextByte, err := hex.DecodeString(ciphertextHex)
	if err != nil {
		t.Error(err)
	}

	// 对称密钥
	key := "IRCMQGSDPDVVYYJVVJLULIJAXBNFPHYT"
	// 向量
	iv := "IRCMQGSDPDVVYYJV"

	plaintextByte, err := AesDecode(ciphertextByte, []byte(key), []byte(iv))
	if err != nil {
		t.Error(err)
	}

	plaintext := string(plaintextByte)
	t.Log("解密结果：", plaintext)

	if plaintext == "hello world" {
		t.Log("eq")
	} else {
		t.Error("not eq")
	}
}

// 生成随机字节
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := crand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

// 生成全大写字母的随机字符串
func GenrateRandomString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
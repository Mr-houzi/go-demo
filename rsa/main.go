package main

import (
	"encoding/hex"
	"fmt"
	"godemo/rsa/utils"
)

func main() {
	//rsa 密钥文件产生
	fmt.Println("-------------------------------获取RSA公私钥-----------------------------------------")
	prvKey, pubKey, err := utils.GenRsaKey()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(prvKey))
	fmt.Println(string(pubKey))

	fmt.Println("-------------------------------进行签名与验证操作-----------------------------------------")
	var data = "卧了个槽，这么神奇的吗？？！！！  ԅ(¯﹃¯ԅ) ！！！！！！）"
	fmt.Println("对消息进行签名操作...")
	signData, err := utils.RsaSignWithSha256([]byte(data), prvKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("消息的签名信息： ", hex.EncodeToString(signData))
	fmt.Println("\n对签名信息进行验证...")
	if ok, err := utils.RsaVerySignWithSha256([]byte(data), signData, pubKey); !ok {
		fmt.Println(err)
	}else {
		fmt.Println("签名信息验证成功，确定是正确私钥签名！！")
	}

	fmt.Println("-------------------------------进行加密解密操作-----------------------------------------")
	ciphertext, err := utils.RsaEncrypt([]byte(data), pubKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("公钥加密后的数据：", hex.EncodeToString(ciphertext))
	sourceData, err := utils.RsaDecrypt(ciphertext, prvKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("私钥解密后的数据：", string(sourceData))
}


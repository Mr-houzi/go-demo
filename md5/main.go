package main

import (
	"fmt"
	"godemo/md5/utils/md5"
)

func main()  {
	plaintext := "hello word"

	ciphertext := md5.Encode(plaintext)
	fmt.Println("md5 Encode:", ciphertext)

	ciphertext16 := md5.Encode16(plaintext)
	fmt.Println("md5 Encode 16:", ciphertext16)
}



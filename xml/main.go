package main

import (
	"encoding/xml"
	"fmt"
)

type Msg struct {
	Header Header `xml:"header"`
	Body Body `xml:"body"`
}

type Header struct {
	Version string `xml:"version"`
}

type Body struct {
	Name string `xml:"name"`
	Age  int32 `xml:"age"`
	Sex string `xml:"sex"`
}

func main()  {
	struc2xml()
	xml2struct()
}

func struc2xml() {
	msg := Msg{
		Header{
			Version: "1.0",
		},
		Body{
			Name: "xiaoming",
			Age: 16,
			//Sex: "male", // 当值不存在时，golang暂时无法解析成自闭合标签</sex> ; https://github.com/golang/go/issues/21399
		},
	}

	b, err := xml.Marshal(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func xml2struct() {
	xmlStr := `<Msg>
    <header>
        <version>1.0</version>
    </header>
    <body>
        <name>xiaoming</name>
        <age>16</age>
        <sex></sex>
    </body>
</Msg>`

	msgStruct := Msg{}
	err := xml.Unmarshal([]byte(xmlStr), &msgStruct)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", msgStruct)
}
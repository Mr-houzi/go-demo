package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	//Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}

var validate *validator.Validate

func main()  {


	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
	}

	validate = validator.New()
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	//注册翻译器
	if err := zh_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		fmt.Println(err)
	}

	if err := validate.Struct(user); err != nil {
		fmt.Println(err)

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Translate(trans))//Age必须大于18
			return
		}
	}
}

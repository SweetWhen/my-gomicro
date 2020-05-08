package main

import (
    "fmt"
    "go-micro-study/microTools/AppLib"
    "gopkg.in/go-playground/validator.v9"
    "log"
    "testing"
)

type Users struct {
    Username string `validate:"required,min=6,max=20" vmsg:"用户名必须大于6位以上"`
    Userpwd string `validate:"required,min=6,max=18" vmsg:"用户密码必须大于6位以上"`
    //TestName string `validate:"required,username" vmsg:"用户名不正确"`
    Usertags []string `validate:"required,min=1,max=5,unique,dive,usertag" vmsg:"用户标签不合法"`
}

func TestSimpleStruct(t *testing.T)  {
    user := &Users{Username:"aaaaa", Userpwd:"123"}

    valid := validator.New()
    err := valid.Struct(user)
    if err != nil {
        if errs, ok := err.(validator.ValidationErrors); ok {
            for _, e := range errs {
                fmt.Println(e.Value())
                fmt.Println(e.Field())
                fmt.Println(e.Tag())
                AppLib.GetValidMsg(user, e.Field())
            }

        }
        log.Fatal(err)
    }

    fmt.Println("验证成功")

}

//自定义正则匹配
func TestStruct(t *testing.T) {
   userTags := []string{"aa", "bb","cc","dd","ee"}
   valid := validator.New()
   _ = AppLib.AddRegexTag("usertag", "^[\u4e00-\u9fa5a-zA-Z0-9]{2,4}$", valid)
   //_ = AppLib.AddRegexTag("username", "[a-zA-Z]\\w{5,19}", valid)
   u := &Users{Username:"ishenyi", Userpwd:"12ddd3e",
       //TestName:"1jjfta",
       Usertags: userTags,
   }
   err := AppLib.ValidErrMsg(u, valid.Struct(u))
   if err != nil {
      log.Fatal(err)
   }

   log.Printf("验证成功")
}



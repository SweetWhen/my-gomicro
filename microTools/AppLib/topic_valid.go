package AppLib

import (
    "fmt"
    "gopkg.in/go-playground/validator.v9"
    "reflect"
    "regexp"
)

type Topic struct {
    TopicID int `json:"id"`
    TopicTitle string `json:"title" binding:"min=4,max=20"`
    //nefield表示与该结构体中的某一个字段不能够相等,当然我们也可以定义 efield表示与某一个字段相等用于两次输入密码确认
    TopicShortTitle string `json:"stitle" binding:"required,nefield=TopicTitle"`
    UserIp string `json:"ip" binding:"ipv4"` //ipv4地址格式
    TopicScore int `json:"score" binding:"omitempty,gt=5"` //要么不填，填就要大于5
    TopicUrl string `json:"url" binding:"omitempyt,topicurl"` //使用自定义的tag: topicurl
}

type Topics struct {
    TopicList []Topic `json:"topics" binding:"gt=0,lte=3,topics,dive"` //让验证器进入Topic里面继续验证
    TopicSize int `json:"size"`
}

//topicurl
func TopicUrl(v *validator.Validate, topStruct reflect.Value, //顶层struct
    currentStructOrField reflect.Value,  //嵌套struct中当前struct
    field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
       _, ok1 := topStruct.Interface().(*Topics)
       _, ok2 := topStruct.Interface().(*Topic)
       if ok1 || ok2 { //确定是我们定义的Topic结构体
            fmt.Println(field.String()) //这个就是客户端给的东西
            getValue := field.String()
            if m, _ := regexp.MatchString(`^\w{5,12}$`, getValue); m {
                //合法
                return true
            } else {
                return false //验证不通过
            }
        }

    return false //验证不通过

}

//topics
func TopicsValidate(v *validator.Validate, topStruct reflect.Value, //顶层struct
    currentStructOrField reflect.Value,  //嵌套struct中当前struct
    field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
    topics, ok1 := topStruct.Interface().(*Topics)
    if ok1 && topics.TopicSize == len(topics.TopicList) { //确定是我们定义的Topic结构体
            //合法
            return true
    }

    return false //验证不通过
}
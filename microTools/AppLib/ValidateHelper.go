package AppLib

import (
    "fmt"
    "gopkg.in/go-playground/validator.v9"
    "reflect"
    "regexp"
)

func GetValidMsg(obj interface{}, field string)  {
    getObj := reflect.TypeOf(obj)
    if f, exist := getObj.Elem().FieldByName(field); exist {
        fmt.Println(f.Tag.Get("vmsg"))
    }
}

func ValidErrMsg(obj interface{}, err error) error  {
    getObj := reflect.TypeOf(obj)
    if err != nil {
        if errs, ok := err.(validator.ValidationErrors); ok {
            for _, e := range errs {
                if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
                    if value, ok := f.Tag.Lookup("vmsg"); ok {
                        return fmt.Errorf("%s", value)
                    } else {
                        return fmt.Errorf("%s", e)
                    }
                } else {
                    return fmt.Errorf("%s", e)
                }
            }
        }
    }

    return err
}

func AddRegexTag(tagName string, pattern string, v*validator.Validate) error  {
    return v.RegisterValidation(tagName, func(fl validator.FieldLevel) bool {
        match, _ := regexp.MatchString(pattern, fl.Field().String())

        return match
    }, false)
}
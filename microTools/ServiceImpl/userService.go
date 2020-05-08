package ServiceImpl

import (
    "context"
    "fmt"
    Service "go-micro-study/Services"
    "go-micro-study/microTools/AppInit"
    "go-micro-study/microTools/DBModels"
    "time"
)

type UserService struct {
    
}

func (service *UserService) Call(ctx context.Context, req *Service.TestRequest, resp*Service.TestResponse) error {
    resp.Data = fmt.Sprintf("Id:%d, 在Call me", req.Id)
    fmt.Println("Call invoke....")
    return nil
}

func (*UserService)UserReg(ctx context.Context, in *Service.UserModel, out *Service.RegResponse) (err error) {
    users := DBModels.Users{UserName:in.UserName,
        UserPw:in.UserPwd, UserDate:time.Now()}
    err = AppInit.GetDb().Create(&users).Error
    if err != nil {
        out.Message = err.Error()
        out.Status = "error"
        return
    } else {
        out.Status = "success"
    }
    return
}

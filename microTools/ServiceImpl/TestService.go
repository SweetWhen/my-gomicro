package ServiceImpl

import (
    "context"
    Service "go-micro-study/Services"
    "strconv"
)

type TestService struct {
    
}

func (h *TestService) Call(ctx context.Context, in *Service.TestRequest, out *Service.TestResponse) error {
    out.Data = "test"+strconv.Itoa(int(in.Id))

    return nil
}
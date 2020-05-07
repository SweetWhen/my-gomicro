package wrappers

import (
    "context"
    "github.com/afex/hystrix-go/hystrix"
    "github.com/micro/go-micro/client"
    grpcSvc "go-micro-study/Services"
    "strconv"
)

type ProdsWrapper struct {
    client.Client
}

func NewProdsWrapper(c client.Client) client.Client {
    return &ProdsWrapper{c}
}

func (l *ProdsWrapper)Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error  {
    cmdName := req.Service()+"."+req.Endpoint()
    //1.配置configuration
    configA := hystrix.CommandConfig{
        Timeout: 1000,
        RequestVolumeThreshold: 2, //阈值：意思是有20个请求才进行错误百分比计算
        ErrorPercentThreshold:50, //错误百分比 20% 的错误发生之后，就直接执行降级方法
        SleepWindow: 5000, //过多少毫秒之后重新尝试后端被降级的服务
    }
    //2.配置command
    hystrix.ConfigureCommand(cmdName, configA)

    return hystrix.Do(cmdName,
        func() error {
            return l.Client.Call(ctx,req, rsp, opts...)
        },
        func(e error) error {
            //defaultProds(rsp)
            defaultData(rsp)
            return nil
        })
}

//通用降级方法
func defaultData(rsp interface{})  {
    switch t := rsp.(type) {
    case *grpcSvc.ProdListResp:
        defaultProds(rsp)
    case *grpcSvc.ProdDetailtResp:
        t.Data = &grpcSvc.ProdModel{ProdId:int32(10), ProdName:"降级商品详情"+strconv.Itoa(10)}
    default:

    }
}

//商品列表详情
func defaultProds(rsp interface{})  {
    data := make([]*grpcSvc.ProdModel, 0)
    for i := 0; i < 5; i++ {
        data = append(data, &grpcSvc.ProdModel{ProdId:int32(20+i), ProdName:"prodname"+strconv.Itoa(i)})
    }
    resp := rsp.(*grpcSvc.ProdListResp)
    resp.Data = data
    return
}

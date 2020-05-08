package main

import (
    "context"
    "fmt"
    "github.com/micro/go-micro"
    "github.com/micro/go-micro/client"
    "github.com/micro/go-micro/metadata"
    "github.com/micro/go-micro/registry"
    "github.com/micro/go-micro/web"
    "github.com/micro/go-plugins/registry/consul"
    grpcSvc "go-micro-study/Services"
    "go-micro-study/webLib"
    "go-micro-study/wrappers"
)

type logWrapper struct {
    client.Client
}

func (l *logWrapper)Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error  {
    fmt.Println("调用接口")
    md, _ := metadata.FromContext(ctx)
    fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
    return l.Client.Call(ctx, req, rsp, opts...)
}

func NewLogWrapper(c client.Client) client.Client  {
    return &logWrapper{c}
}

func main()  {
    //consul启动： consul.exe agent -server -bootstrap -ui -client 0.0.0.0 -bind 192.168.1.101 -data-dir=F:/consul_data
    Reg := consul.NewRegistry(
       registry.Addrs("192.168.1.101:8500"), //服务发现地址，也就是前
       // 面启动的consul
    )
    //Reg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
    mySvc := micro.NewService(
        micro.Name("prodservice.client"),
        micro.WrapClient(NewLogWrapper),
        micro.WrapClient(wrappers.NewProdsWrapper),
        )
    prodSvc := grpcSvc.NewProdService("prodservice", mySvc.Client())
    httpSvc := web.NewService(
        web.Name("httpprodservice"), //服务名称
        web.Address(":8001"), //监听端口
        web.Handler(webLib.NewGinRouter(prodSvc)), //将gin引入
        web.Registry(Reg), //将consul引入
     )


    httpSvc.Init() //在使用到命令行参数的时候使用到： go run prod_main.go --server_address :8001
    httpSvc.Run() //启动
}
package main

import (
    "github.com/micro/go-micro"
    "github.com/micro/go-micro/registry"
    "github.com/micro/go-plugins/registry/consul"
    Service "go-micro-study/Services"
    _ "go-micro-study/microTools/AppInit"
    "go-micro-study/microTools/ServiceImpl"
)

func main()  {
    Reg := consul.NewRegistry(registry.Addrs("192.168.1.101:8500"))
    //Reg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
    mySvc := micro.NewService(
        micro.Name("api.jtthink.com.myapp"),
        micro.Address(":8001"),
        micro.Registry(Reg),
        )

    Service.RegisterUserServiceHandler(mySvc.Server(), new(ServiceImpl.UserService))
    mySvc.Run()
}

package main

import (
    "github.com/micro/go-micro"
    "github.com/micro/go-micro/registry"
    "github.com/micro/go-plugins/registry/consul"
    Service "go-micro-study/Services"
    "go-micro-study/grpcServer/ServiceImpl"
)
func main()  {
    Reg := consul.NewRegistry(
           registry.Addrs("192.168.1.101:8500"),
       )
    //Reg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
    svc := micro.NewService(
        micro.Name("prodservice"),
        micro.Address(":8011"),
        micro.Registry(Reg),
        )

    svc.Init()
    Service.RegisterProdServiceHandler(svc.Server(), new(ServiceImpl.ProdService))
    svc.Run()
}

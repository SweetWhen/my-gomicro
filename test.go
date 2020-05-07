package main

import (
    "context"
    "fmt"
    "github.com/micro/go-micro/client"
    "github.com/micro/go-micro/client/selector"
    "github.com/micro/go-micro/registry"
    "github.com/micro/go-micro/registry/etcd"
    myhttp "github.com/micro/go-plugins/client/http"
    "github.com/micro/go-plugins/registry/consul"
    Models "go-micro-study/bak/models"
    "io/ioutil"
    "log"
    "net/http"
)

func callApi2(s selector.Selector)  {
    myClient := myhttp.NewClient(client.Selector(s),
        client.ContentType("application/json"),
        )
    request := myClient.NewRequest("prodservice", "/v1/prods", Models.ProdsReq{Size: 1})
    var resp Models.ProdListResp

    err := myClient.Call(context.Background(), request, &resp)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(resp.Data)

    return
}

func callApi(addr string, path string, method string) (string, error)  {
    req, err := http.NewRequest(method, "http://"+addr+path, nil)
    if err != nil {
    }

    client := http.DefaultClient
    response, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer response.Body.Close()

   buff, _ := ioutil.ReadAll(response.Body)
    return string(buff), nil
}

func main()  {
    //consulReg := consul.NewRegistry(
    //    registry.Addrs("192.168.1.101:8500"),
    //    )
    etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
    mySeletor := selector.NewSelector(
        selector.Registry(etcdReg),
        selector.SetStrategy(selector.Random),
        )

    callApi2(mySeletor)

    //for  {
    //    services, err := consulReg.GetService("prodservice")
    //    if err != nil {
    //        log.Fatal(err)
    //    }
    //
    //    next := selector.RoundRobin(services)
    //    node, err := next()
    //    if err != nil{
    //        log.Fatal(err)
    //    }
    //    callres, err := callApi(node.Address, "/v1/prods", "GET")
    //    if err != nil{
    //        log.Fatal(err)
    //    }
    //    fmt.Println(node.Id, node.Address, node.Metadata, callres)
    //
    //    time.Sleep(time.Second*1)
    //}

}
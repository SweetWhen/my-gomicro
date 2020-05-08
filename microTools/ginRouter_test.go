package main

import (
    "context"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "go-micro-study/microTools/sidecar"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "testing"
)


func TestRouter(t *testing.T)  {
    ginRouter := gin.Default()
    v1 := ginRouter.Group("/v1")
    {
        v1.Handle("POST", "/test", func(context *gin.Context) {
            context.JSON(200, gin.H{"data":"test1"})
        })
    }
    server := &http.Server{Addr:":8088", Handler:ginRouter}
    svc := sidecar.NewService("api.jtthink.com.test1")
    svc.AddNode("test-"+uuid.New().String(), 8088, "127.0.0.1:8088")
    handler := make(chan error)
    go func() {
        handler <- server.ListenAndServe()
    }()
    go func() {
        //注册服务
        err := sidecar.RegService(svc)
        if err != nil {
            handler <- err
        }
    }()
    go func() {
        notify := make(chan os.Signal)
        signal.Notify(notify, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
        handler <- fmt.Errorf("%s", <-notify)

    }()
    getH := <- handler
    fmt.Println(getH.Error())
    //反注册服务
    err := sidecar.UnRegService(svc)
    if err != nil {
        fmt.Println(err)
    }
    server.Shutdown(context.Background())

}

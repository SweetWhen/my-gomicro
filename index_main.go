package main

import (
    "github.com/gin-gonic/gin"
    "github.com/micro/go-micro/web"
)

func main()  {

    ginRouter := gin.Default() //gin web框架
    ginRouter.Handle("GET", "/", func(context *gin.Context) {
        data := make([]interface{}, 0)
        context.JSON(200, gin.H{
            "data":data,
        })
    })

    server := web.NewService(
        web.Address(":8000"), //监听端口
        web.Handler(ginRouter), //将gin引入
    )

    server.Run() //启动
}

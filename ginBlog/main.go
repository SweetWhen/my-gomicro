package main

import (
    "fmt"
    "go-micro-study/ginBlog/pkg/setting"
    "go-micro-study/ginBlog/routers"
    "net/http"
)

func main()  {
    router := routers.InitRouter()

    s := &http.Server{
        Addr: fmt.Sprintf(":%d", setting.HTTPPort),
        Handler: router,
        ReadTimeout:setting.ReadTimeout,
        WriteTimeout:setting.WriteTimeout,
        MaxHeaderBytes: 1<<20,
    }

    s.ListenAndServe()
}


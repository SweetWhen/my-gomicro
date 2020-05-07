package webLib

import (
    "github.com/gin-gonic/gin"
    grpcSvc "go-micro-study/Services"
)

func NewGinRouter(prodSvc grpcSvc.ProdService) *gin.Engine {
    ginRouter := gin.Default() //gin web框架
    ginRouter.Use(InitMiddleware(prodSvc), ErrorMiddleware())
    v1Group := ginRouter.Group("/v1")
    {
        v1Group.Handle("POST", "/prods", GetProdList)
        v1Group.Handle("GET", "/prods/:pid", GetProdDetail)
    }

    return ginRouter
}

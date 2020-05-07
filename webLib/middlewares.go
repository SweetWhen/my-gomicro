package webLib

import (
    "fmt"
    "github.com/gin-gonic/gin"
    grpcSvc "go-micro-study/Services"
)

func InitMiddleware(prodSvc grpcSvc.ProdService) gin.HandlerFunc  {
    return func(context *gin.Context) {
        context.Keys = make(map[string]interface{})
        context.Keys["prodservice"] = prodSvc

        context.Next()
    }
}

func ErrorMiddleware() gin.HandlerFunc  {
    return func(context *gin.Context) {
        defer func() {
            if r := recover(); r != nil {
                context.JSON(500, gin.H{"status":fmt.Sprintf("%s", r)})
                context.Abort()
            }
            context.Next()
        }()
    }
}
package webLib

import (
    "context"
    "github.com/afex/hystrix-go/hystrix"
    "github.com/gin-gonic/gin"
    grpcSvc "go-micro-study/Services"
    "strconv"
)

func defaultProds() (resp *grpcSvc.ProdListResp, err error) {
    data := make([]*grpcSvc.ProdModel, 0)
    for i := 0; i < 5; i++ {
        data = append(data, &grpcSvc.ProdModel{ProdId:int32(20+i), ProdName:"prodname"+strconv.Itoa(i)})
    }

    resp = &grpcSvc.ProdListResp{Data:data}
    return
}

func GetProdList(ginCtx *gin.Context)  {
        var req grpcSvc.ProdsReq
        var resp *grpcSvc.ProdListResp

    prodSvc := ginCtx.Keys["prodservice"].(grpcSvc.ProdService)
        err := ginCtx.Bind(&req)
        if err != nil {
            ginCtx.JSON(500, gin.H{"status":err.Error()})
        } else {
            resp, err =  prodSvc.GetProdsList(context.Background(), &req)
            ginCtx.JSON(200, gin.H{"data": resp.Data})
            return
            //熔断代码改造
            //1.配置configuration
            configA := hystrix.CommandConfig{
                Timeout: 1000,
            }
            //2.配置command
            hystrix.ConfigureCommand("getprods", configA)
            //3. 执行，使用Do方法
           err = hystrix.Do("getprods", func() error {
                resp, err =  prodSvc.GetProdsList(context.Background(), &req)
                return err
            }, func(e error) error {
               resp, err = defaultProds()
                return e
           })
            resp, err =  prodSvc.GetProdsList(context.Background(), &req)
            if err != nil {
                ginCtx.JSON(500, gin.H{"status":err.Error()})
            }
            ginCtx.JSON(200, gin.H{"data": resp.Data})
        }
}

func PanicError(err error)  {
    if err != nil {
        panic(err)
    }
}

func GetProdDetail(ginCtx *gin.Context)  {
    var prodReq grpcSvc.ProdsReq
    PanicError(ginCtx.BindUri(&prodReq))
    prodSvc := ginCtx.Keys["prodservice"].(grpcSvc.ProdService)
    resp, _ := prodSvc.GetProdsDetail(context.Background(), &prodReq)
    ginCtx.JSON(200, gin.H{"data": resp.Data})

}
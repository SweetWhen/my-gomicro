package ServiceImpl

import (
    "context"
    "fmt"
    Service "go-micro-study/Services"
    "strconv"
    "time"
)

type ProdService struct {
}

func (service *ProdService) GetProdsDetail(ctx context.Context, req*Service.ProdsReq, resp*Service.ProdDetailtResp) error {
    time.Sleep(time.Second*3)
    fmt.Printf("req.size:%d, req.prodid:%d\n",req.Size, req.ProdId)
    resp.Data = newProd(req.ProdId, "测试商品详情")
    return nil
}

func newProd(id int32, pname string) *Service.ProdModel {
    return &Service.ProdModel{ProdId:id, ProdName:pname}
}


func (*ProdService)GetProdsList(ctx context.Context,  req *Service.ProdsReq,resp *Service.ProdListResp) error  {
    fmt.Println("GetProdsList:",req.Size)
    //time.Sleep(time.Second*3)
    models := make([]*Service.ProdModel, 0)
    var i int32
    for i = 0; i < req.Size; i++ {
        models = append(models, newProd(i+100, "prodName"+strconv.Itoa(int(i)+100)))
    }
    resp.Data = models

    return nil
}
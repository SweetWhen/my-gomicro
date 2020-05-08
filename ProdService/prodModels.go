package ProdService

import "strconv"

type ProdModel struct {
    ProdID int
    ProdName string
}

func NewProd(id int, name string) *ProdModel  {
    return &ProdModel{
        ProdID:   id,
        ProdName: name,
    }
}

func NewProdList(n int ) []*ProdModel  {
    ret := make([]*ProdModel, 0)
    for i := 0; i < n; i++ {
        ret =append(ret, NewProd(100+i, "prodName"+strconv.Itoa(i)))
    }

    return ret
}

func NewProdMap(n int ) (res map[string]interface{})  {
    res = make(map[string]interface{})
    ret := make([]*ProdModel, 0)
    for i := 0; i < n; i++ {
        ret =append(ret, NewProd(100+i, "prodName"+strconv.Itoa(i)))
    }
    res["data"] = ret

    return res
}
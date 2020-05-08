package sidecar

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type JSONReq struct {
    Jsonrpc string
    Method string
    Params []*Service
    Id int
}

func NewJSONReq(servcie *Service, endpoint string) *JSONReq  {
    return &JSONReq{
        Jsonrpc:"2.0",
        Method:endpoint,
        Params: []*Service{servcie},
        Id:1}
}
var RegistryURI = "http://localhost:8000"
func reqRegistry(jsonReq *JSONReq) error  {
    bs, err := json.Marshal(jsonReq)
    if err != nil {
        return err
    }
    rsp, err := http.Post(RegistryURI, "application/json", bytes.NewReader(bs))
    if err != nil {
        return err
    }
    defer rsp.Body.Close()

    res, err := ioutil.ReadAll(rsp.Body)
    if err != nil {
        return err
    }

    fmt.Println(string(res))

    return nil
}

func UnRegService( svc *Service) error  {
    return reqRegistry(NewJSONReq(svc, "Registry.Deregister"))
}

func RegService(svc *Service) error  {
    return reqRegistry(NewJSONReq(svc, "Registry.Register"))
}
package api

import (
	"net/http"
	"seisan/internal/model"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	req := model.NewReq(r)
	if !req.Check() {
		resp := model.NewErrorResp("签名校验失败")
		resp.Write(w)
		return
	}
	if req.Query.Kind == "" {
		resp := model.NewErrorResp("参数不合法")
		resp.Write(w)
		return
	}
	//data, err := core.GetResult(req.Query)
	//if err != nil {
	//	resp := model.NewErrorResp(err.Error())
	//	resp.Write(w)
	//	return
	//}
	//resp := model.NewResp(data)
	//resp.Write(w)
	resp := model.NewErrorResp("无")
	resp.Write(w)
	return
}

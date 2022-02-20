package api

import (
	"net/http"
	"seisan/internal/core"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	req := core.NewReq(r)
	if !req.Check() {
		//resp := core.NewErrorResp("签名校验失败")
		//resp.Write(w)
		//return
	}
	//if req.Query.Kind == "" {
	//	resp := core.NewErrorResp("参数不合法")
	//	resp.Write(w)
	//	return
	//}
	data, err := core.GetResult(req)
	if err != nil {
		resp := core.NewErrorResp(err.Error())
		resp.Write(w)
		return
	}
	resp := core.NewResp(data)
	resp.Write(w)
}

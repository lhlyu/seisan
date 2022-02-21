package api

import (
	"net/http"
	"seisan/internal/handler"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	req := handler.NewReq(r)
	if !req.Check() {
		resp := handler.NewErrorResp("签名校验失败")
		resp.Write(w)
		return
	}
	if req.Query.Kind == "" {
		resp := handler.NewErrorResp("参数不合法")
		resp.Write(w)
		return
	}
	data, err := handler.GetResult(req)
	if err != nil {
		resp := handler.NewErrorResp(err.Error())
		resp.Write(w)
		return
	}
	resp := handler.NewResp(data)
	resp.Write(w)
}

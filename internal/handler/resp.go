package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"seisan/internal/util"
)

type Resp struct {
	// 响应的结构体
	D string `json:"d"`
	// 响应的世界
	T string `json:"t"`
	// 响应的签名
	S string `json:"s"`
}

func NewResp(data map[string]interface{}) *Resp {
	if len(data) == 0 {
		return &Resp{}
	}
	d, _ := util.AesEny(key, key, data)
	now := util.Now()
	sign := util.Md5Encode(d, now, key)
	return &Resp{
		D: d,
		T: now,
		S: sign,
	}
}

func NewErrorResp(msg string) *Resp {
	return &Resp{
		S: msg,
	}
}

func (r *Resp) Write(w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json")
	b, _ := json.Marshal(r)
	io.WriteString(w, string(b))
}

type Item struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
	Gender int    `json:"gender,omitempty"`
}

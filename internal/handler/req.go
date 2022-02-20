package handler

import (
	"log"
	"net/http"
	"seisan/internal/model"
	"seisan/internal/util"
)

type Req struct {
	// 查询的请求体
	Q string `json:"q"`
	// 查询时间戳
	T string `json:"t"`
	// 签名
	S string `json:"s"`
	// 解密后的对象
	Query *model.Query `json:"-"`
}

func NewReq(r *http.Request) *Req {
	values := r.URL.Query()
	q := values.Get("q")
	t := values.Get("t")
	s := values.Get("s")
	return &Req{
		Q: q,
		T: t,
		S: s,
	}
}

func (r *Req) Check() bool {
	query := &model.Query{}
	if err := util.AesDec(key, key, r.Q, query); err != nil {
		log.Println("Check AesDec:", r.Query)
		return false
	}
	r.Query = query
	sign := util.Md5Encode(r.Q, r.T, key)
	return sign == r.S
}

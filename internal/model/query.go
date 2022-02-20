package model

type Query struct {
	// 查询的种类
	Kind string `json:"kind"`
	// 查询的参数
	Params map[string]interface{} `json:"params"`
}

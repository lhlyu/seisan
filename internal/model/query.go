package model

type Query struct {
	// 查询的种类
	Kind string `json:"kind"`
	// 生成的数量
	Number int `json:"number"`

	// 性别：0 - 随机; 1 - 男; 2 - 女
	Gender int `json:"gender,omitempty"`
	// 姓氏：0 - 随机; 1 - 单姓; 2 - 女
	Surname int `json:"surname,omitempty"`
	// 自定义前缀
	Prefix string `json:"prefix,omitempty"`
	// 自定义后缀
	Suffix string `json:"suffix,omitempty"`
	// 名字长度: 0 - 随机[1,2]
	NameLength int `json:"name_length,omitempty"`
}

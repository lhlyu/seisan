package demo

import (
	"log"
	dd "seisan/internal/dict"
)

func GetLabel(index int) string {
	log.Println("--->", index)
	v, ok := dd.D[index]
	if ok {
		return v
	}
	return "没有找到"
}

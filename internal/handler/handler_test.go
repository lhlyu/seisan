package handler

import (
	"fmt"
	"seisan/internal/model"
	"testing"
)

func TestName(t *testing.T) {
	items := GetChineseNames(&model.Query{
		Kind:       "chinese_name",
		Number:     20,
		Gender:     2,
		Surname:    1,
		Prefix:     "",
		Suffix:     "",
		NameLength: 2,
	})
	for _, item := range items {
		fmt.Printf("%s %-5s | %2d\n", item.Prefix, item.Suffix, item.Gender)
	}
}

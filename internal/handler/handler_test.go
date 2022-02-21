package handler

import (
	"fmt"
	"seisan/internal/model"
	"sort"
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
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Score > items[j].Score
	})
	for _, item := range items {
		fmt.Printf("%s %-5s | %2d | %2d\n", item.Prefix, item.Suffix, item.Gender, item.Score)
	}
}

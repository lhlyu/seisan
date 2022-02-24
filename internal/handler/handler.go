package handler

import (
	"bytes"
	"errors"
	"math/rand"
	"seisan/internal/dict"
	"seisan/internal/model"
	"strings"
	"time"
)

var (
	err_not_found_handler = errors.New("处理不了")
	err_create_empty      = errors.New("系统已经努力了，但是没有结果")
)

var kinds = map[string]func(*model.Query) []*Item{
	"chinese_name": GetChineseNames,
	"potion":       GetPotions,
	"kungfu":       GetKungFus,
	"weapon":       GetWeapons,
	"org":          GetOrgs,
	"treasure":     GetTreasures,
	"place":        GetPlaces,
}

func GetResult(req *Req) (map[string]interface{}, error) {
	fn, ok := kinds[req.Query.Kind]
	if !ok {
		return nil, err_not_found_handler
	}

	req.Query.Number = 20
	req.Query.Prefix = strings.TrimSpace(req.Query.Prefix)
	req.Query.Suffix = strings.TrimSpace(req.Query.Suffix)
	if req.Query.NameLength <= 0 {
		req.Query.NameLength = 1
	}
	if req.Query.NameLength > 5 {
		req.Query.NameLength = 5
	}

	items := fn(req.Query)
	if len(items) == 0 {
		return nil, err_create_empty
	}
	return map[string]interface{}{
		"list": items,
	}, nil
}

// 获取名字
func GetChineseNames(query *model.Query) []*Item {
	distincter := make(map[string]struct{})
	items := make([]*Item, 0)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < query.Number; i++ {
		prefix := query.Prefix
		surname := query.Surname
		if surname == 0 {
			surname = rand.Intn(2) + 1
		}
		gender := query.Gender
		if gender == 0 {
			gender = rand.Intn(2) + 1
		}
		if prefix == "" {
			if surname == 1 {
				prefix = dict.Single[rand.Intn(dict.SingleLen)]
			}
			if surname == 2 {
				prefix = dict.Double[rand.Intn(dict.DoubleLen)]
			}
		}

		suffix := query.Suffix
		if suffix == "" {
			buf := bytes.Buffer{}
			for j := 0; j < query.NameLength; j++ {
				if gender == 1 {
					suffix = dict.Man[rand.Intn(dict.ManLen)]
				}

				if gender == 2 {
					suffix = dict.Woman[rand.Intn(dict.WomanLen)]
				}

				buf.WriteString(suffix)
			}
			suffix = buf.String()
		}

		name := prefix + suffix

		// 去重
		if _, has := distincter[name]; has {
			continue
		}
		items = append(items, &Item{
			Prefix: prefix,
			Suffix: suffix,
			Gender: gender,
		})

		distincter[name] = struct{}{}
	}
	return items
}

// 获取丹药
func GetPotions(query *model.Query) []*Item {
	return nil
}

// 获取功法
func GetKungFus(query *model.Query) []*Item {
	return nil
}

// 获取武器
func GetWeapons(query *model.Query) []*Item {
	return nil
}

// 获取组织
func GetOrgs(query *model.Query) []*Item {
	return nil
}

// 获取天材地宝
func GetTreasures(query *model.Query) []*Item {
	return nil
}

// 获取地名
func GetPlaces(query *model.Query) []*Item {
	return nil
}

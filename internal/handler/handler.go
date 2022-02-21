package handler

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"seisan/internal/dict"
	"seisan/internal/model"
	"sort"
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
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Score > items[j].Score
	})
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
		eles := make([][]string, 0)
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
				val := dict.Single[rand.Intn(dict.SingleLen)]
				eles = append(eles, val)
				prefix = val[0]
			}
			if surname == 2 {
				val := dict.Double[rand.Intn(dict.DoubleLen)]
				eles = append(eles, val)
				prefix = val[0]
			}
		}

		buf := bytes.Buffer{}
		for j := 0; j < query.NameLength; j++ {
			suffix := query.Suffix
			if suffix == "" {
				if gender == 1 {
					val := dict.Man[rand.Intn(dict.ManLen)]
					eles = append(eles, val)
					suffix = val[0]
				}

				if gender == 2 {
					val := dict.Woman[rand.Intn(dict.WomanLen)]
					eles = append(eles, val)
					suffix = val[0]
				}
			}
			buf.WriteString(suffix)
		}
		name := buf.String()
		// 去重
		if _, has := distincter[name]; has {
			continue
		}
		// 评分
		score := getScore(gender, eles)

		items = append(items, &Item{
			Prefix: prefix,
			Suffix: name,
			Gender: gender,
			Score:  score,
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

// 打分
func getScore(gender int, eles [][]string) int {
	fmt.Println(eles)
	if len(eles) == 0 {
		return 0
	}
	var score int
	// 声调
	sd := make(map[string]struct{})
	// 声母
	sm := make(map[string]struct{})

	for _, ele := range eles {
		for i := 1; i < len(ele); i++ {
			// 声调
			if i%2 == 0 {
				if _, ok := sd[ele[i]]; ok {
					score--
					continue
				}
				score++
				sd[ele[i]] = struct{}{}
				continue
			}
			if _, ok := sm[ele[i]]; ok {
				score--
				continue
			}
			score++
			sm[ele[i]] = struct{}{}
			continue
		}
	}
	// 男平女仄结尾加分
	lastWord := eles[len(eles)-1]
	if gender == 1 {
		if lastWord[2] == "1" || lastWord[2] == "2" {
			score--
		} else {
			score++
		}
	}
	if gender == 2 {
		if lastWord[2] == "3" || lastWord[2] == "4" {
			score--
		} else {
			score++
		}
	}
	return score
}

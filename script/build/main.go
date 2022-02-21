package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var py = pinyin.NewArgs()

func init() {
	py.Style = pinyin.Tone2
	py.Heteronym = true
}

func main() {
	updateKey()
	dictToGo()
	dictToGo2()

}

// 修改秘钥
var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func updateKey() {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 16)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	key := string(b)

	files := []string{
		"./internal/handler/const.go",
		"./src/http/index.ts",
	}

	str := `const key = "\w+"`
	reg := regexp.MustCompile(str)

	for i := range files {
		fb, err := ioutil.ReadFile(files[i])
		if err != nil {
			log.Println("updateKey ReadFile:", err)
			continue
		}
		content := reg.ReplaceAllString(string(fb), fmt.Sprintf(`const key = "%s"`, key))
		err = ioutil.WriteFile(files[i], []byte(content), fs.ModePerm)
		if err != nil {
			log.Println("updateKey WriteFile:", err)
			continue
		}
	}
}

// 将字典转go代码
// 男平女仄结尾
// b p m f d t n l g k h j q x zh ch sh r z c s y w
var sms = map[string]struct{}{
	"b":  {},
	"p":  {},
	"m":  {},
	"f":  {},
	"d":  {},
	"t":  {},
	"n":  {},
	"l":  {},
	"g":  {},
	"k":  {},
	"h":  {},
	"j":  {},
	"q":  {},
	"x":  {},
	"zh": {},
	"ch": {},
	"sh": {},
	"r":  {},
	"z":  {},
	"c":  {},
	"s":  {},
	"y":  {},
	"w":  {},
}

const tmpl = `package dict
// %s

var %sLen = len(%s)

var %s = []string{
%s
}`

func dictToGo() {
	files := [][]string{
		{"script/build/dict/丹药后缀.txt", "Potion"},
		{"script/build/dict/功法后缀.txt", "KungFu"},
		{"script/build/dict/地名后缀.txt", "Place"},
		{"script/build/dict/天材地宝后缀.txt", "Treasure"},
		{"script/build/dict/武器后缀.txt", "Weapon"},
		{"script/build/dict/组织后缀.txt", "Org"},
	}
	for _, f := range files {
		b, err := ioutil.ReadFile(f[0])
		if err != nil {
			log.Println("dictToGo:", err)
			continue
		}

		s := string(b)

		rows := strings.Split(s, "\n")

		dictinter := make(map[string]struct{})

		arr := make([]string, 0)

		for _, row := range rows {
			row = strings.TrimSpace(row)
			if _, ok := dictinter[row]; ok {
				continue
			}

			arr = append(arr, fmt.Sprintf(`    "%s",`, row))

			dictinter[row] = struct{}{}
		}

		name := filepath.Base(f[0])
		name = strings.ReplaceAll(name, ".txt", "")

		content := fmt.Sprintf(tmpl, name, f[1], f[1], f[1], strings.Join(arr, "\n"))

		filename := fmt.Sprintf("internal/dict/%s.go", name)

		ioutil.WriteFile(filename, []byte(content), fs.ModePerm)
	}
}

const tmpl2 = `package dict
// %s

var %sLen = len(%s)

var %s = [][]string{
%s
}`

func dictToGo2() {
	files := [][]string{
		{"script/build/dict/千字文.txt", "Qword"},
		{"script/build/dict/单姓.txt", "Single"},
		{"script/build/dict/复姓.txt", "Double"},
		{"script/build/dict/女名.txt", "Woman"},
		{"script/build/dict/男名.txt", "Man"},
	}

	for _, f := range files {
		b, err := ioutil.ReadFile(f[0])
		if err != nil {
			log.Println("dictToGo:", err)
			continue
		}

		s := string(b)

		rows := strings.Split(s, "\n")

		dictinter := make(map[string]struct{})

		arr := make([]string, 0)

		for _, row := range rows {
			row = strings.TrimSpace(row)
			if _, ok := dictinter[row]; ok {
				continue
			}

			arr = append(arr, getInfo(row)...)

			dictinter[row] = struct{}{}

		}

		name := filepath.Base(f[0])
		name = strings.ReplaceAll(name, ".txt", "")

		content := fmt.Sprintf(tmpl2, name, f[1], f[1], f[1], strings.Join(arr, "\n"))

		filename := fmt.Sprintf("internal/dict/%s.go", name)

		ioutil.WriteFile(filename, []byte(content), fs.ModePerm)
	}
}

// 提取信息
func getInfo(row string) []string {
	rp := pinyin.Pinyin(row, py)

	// 单个字
	if len(rp) == 1 {
		zs := rp[0]

		arr := make([]string, 0)
		for _, z := range zs {
			arr = append(arr, fmt.Sprintf(`    {"%s", "%s", "%d"},`, row, getSm(z), getSd(z)))
			break
		}

		return arr
	}

	// 双字
	z1 := rp[0]
	z2 := rp[1]
	ele := fmt.Sprintf(`    {"%s", "%s", "%d", "%s", "%d"},`, row, getSm(z1[0]), getSd(z1[0]), getSm(z2[0]), getSd(z2[0]))
	return []string{ele}
}

// 提取声调
func getSd(s string) int {
	for _, b := range []byte(s) {
		switch b {
		case 49:
			return 1
		case 50:
			return 2
		case 51:
			return 3
		case 52:
			return 4
		}
	}
	return 0
}

// 提取声母
func getSm(s string) string {
	length := len(s)
	if length < 2 {
		return ""
	}
	prefix := s[0:2]
	if _, ok := sms[prefix]; ok {
		return prefix
	}
	prefix = s[0:1]
	if _, ok := sms[prefix]; ok {
		return prefix
	}
	return ""
}

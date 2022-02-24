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

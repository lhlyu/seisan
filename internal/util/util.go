package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// aes对称加密
func AesEny(key, iv string, content interface{}) (string, error) {
	key = paddingAes(key)
	iv = paddingAes(iv)
	plaintext, err := json.Marshal(content)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	stream := cipher.NewCTR(block, []byte(iv))
	stream.XORKeyStream(plaintext, plaintext)
	return base64.StdEncoding.EncodeToString(plaintext), nil
}

// 解密
func AesDec(key, iv, content string, v interface{}) error {
	key = paddingAes(key)
	iv = paddingAes(iv)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}
	ciptext, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return err
	}
	stream := cipher.NewCTR(block, []byte(iv))
	stream.XORKeyStream(ciptext, ciptext)
	return json.Unmarshal(ciptext, v)
}

// key iv 必须 16位
// 不足在前面补充0
// 超过裁剪
func paddingAes(s string) string {
	length := len(s)
	if length == 16 {
		return s
	}
	if length < 16 {
		return fmt.Sprintf("%016s", s)
	}
	return s[:16]
}

// md5加密
// salt 盐值
func Md5Encode(val string, salt ...string) string {
	if val == "" {
		return ""
	}
	buf := bytes.NewBufferString(val)
	buf.WriteString(strings.Join(salt, ""))

	h := md5.New()
	h.Write(buf.Bytes())
	return hex.EncodeToString(h.Sum(nil))
}

func WaitGroup(fns ...func()) {
	if len(fns) == 0 {
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(fns))
	for _, fn := range fns {
		go func(fn func()) {
			defer wg.Done()
			fn()
		}(fn)
	}
	wg.Wait()
}

// 获取当前时间戳，单位: 毫秒
func Now() string {
	return strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
}

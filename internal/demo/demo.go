package demo

import "io/ioutil"

const Demo = "demo"

var S string

const a = 1

func init() {
	b, err := ioutil.ReadFile("./demo.txt")
	if err != nil {
		S = err.Error()
		return
	}
	S = string(b)
}

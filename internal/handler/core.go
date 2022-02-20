package handler

import "log"

func GetResult(req *Req) (map[string]interface{}, error) {
	log.Println(req.Query)
	return nil, nil
}

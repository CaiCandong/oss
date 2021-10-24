package objects

import (
	"log"
	"net/http"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	// 获取文件名
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	// 保存文件
	c, err := storeObject(r.Body, object)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(c)
}

package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"oss/apiServer/heartbeat"
	"oss/apiServer/locate"
	"oss/apiServer/objects"
)

func main() {
	//加载环境变量
	godotenv.Load()
	//心跳 检测哪些dataServer提供服务
	go heartbeat.ListenHeartBeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS_API_SERVER"), nil))
}

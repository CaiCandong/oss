package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"oss/dataServer/heartbeat"
	"oss/dataServer/locate"
	"oss/dataServer/objects"
)

func main() {
	//加载环境变量
	godotenv.Load()
	//心跳 向apiServer通知存在
	go heartbeat.StartHeartBeat()
	// 对象定位服务
	go locate.StartLocate()
	// 数据通过http协议获取
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS_DATA_SERVER"), nil))
}

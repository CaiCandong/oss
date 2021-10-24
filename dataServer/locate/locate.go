package locate

import (
	"fmt"
	"os"
	"oss/lib/rabbitmq"
	"strconv"
)

func Locate(name string) bool {
	// 返回文件信息
	_, err := os.Stat(name)
	// 根据error类型判断文件是否存在
	return !os.IsNotExist(err)
}

//通过消息队列 提供数据定位在哪台数据服务中
func StartLocate() {
	queue := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer queue.Close()
	queue.Bind("dataServers")
	consume := queue.Consume()
	//监听 dataServers消息队列
	for message := range consume {
		//由于经过JSON编码，所以对象名字上有一对双引号,Unquote用来去除双引号
		object, err := strconv.Unquote(string(message.Body))
		if err != nil {
			panic(err)
		}
		if Locate(os.Getenv("STORAGE_ROOT") + "/objects/" + object) {
			fmt.Println("locate success,prepare send message to ", message.ReplyTo, os.Getenv("LISTEN_ADDRESS_DATA_SERVER"))
			queue.Send(message.ReplyTo, os.Getenv("LISTEN_ADDRESS_DATA_SERVER"))
		}
	}
}

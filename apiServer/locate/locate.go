package locate

import (
	"os"
	"oss/lib/rabbitmq"
	"strconv"
	"time"
)

// 使用临时队列给数据层发送定位请求
func Locate(name string) string {
	queue := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	//通过消息队列 发起定位请求
	queue.Publish("dataServers", name)
	//等待消息队列 收到定位响应
	consume := queue.Consume()
	go func() {
		//避免死等 超时时间为1s
		time.Sleep(time.Second)
		queue.Close()
	}()
	message := <-consume
	str, _ := strconv.Unquote(string(message.Body))
	return str
}

func Exist(name string) bool {
	return Locate(name) != ""
}

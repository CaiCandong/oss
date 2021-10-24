package heartbeat

import (
	"fmt"
	"os"
	"oss/lib/rabbitmq"
	"strconv"
	"sync"
	"time"
)

// 数据服务地址->时间 如:127.0.0.1:8089->2021:10:24:19:03
var dataServers = make(map[string]time.Time)

var mutex sync.Mutex

func ListenHeartBeat() {
	queue := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer queue.Close()
	queue.Bind("apiServers")
	consume := queue.Consume()
	go removeExpiredDataServer()
	for message := range consume {
		dataServer, err := strconv.Unquote(string(message.Body))
		if err != nil {
			fmt.Println(err)
		}
		mutex.Lock()
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}
}

func removeExpiredDataServer() {
	for {
		time.Sleep(5 * time.Second)
		mutex.Lock()
		for server, timer := range dataServers {
			if timer.Add(10 * time.Second).Before(time.Now()) {
				delete(dataServers, server)
			}
		}
		mutex.Unlock()
	}
}

//进行深拷贝
func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()
	dataServer := make([]string, 0)
	for server, _ := range dataServers {
		dataServer = append(dataServer, server)
	}
	return dataServer
}

package heartbeat

import (
	"os"
	"oss/lib/rabbitmq"
	"time"
)

func StartHeartBeat() {
	queue := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer queue.Close()
	for {
		queue.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		//log.Printf("send heart beat at %s", time.Now())
		time.Sleep(5 * time.Second)
	}
}

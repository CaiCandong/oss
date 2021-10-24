package heartbeat

import (
	"log"
	"math/rand"
)

func ChooseRandomDataServer() string {
	dataServers := GetDataServers()
	count := len(dataServers)
	log.Println("count:", count)
	if count == 0 {
		return ""
	}
	idx := rand.Intn(count)
	log.Println("index:", idx)
	return dataServers[idx]
}

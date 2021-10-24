package objects

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"oss/apiServer/locate"
	"oss/lib/objectStream"
	"strings"
)

func get(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	stream, err := getStream(object)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	io.Copy(w, stream)
}

func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("objects %s locate fail", object)
	}
	return objectStream.NewGetStream(server, object)
}

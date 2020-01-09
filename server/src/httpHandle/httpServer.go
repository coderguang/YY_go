package httpHandle

import (
	"net/http"

	"github.com/coderguang/GameEngine_go/sglog"
)

type web_server struct{}

func (h *web_server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	flag := make(chan bool)
	//go ReceiveClientData(w,r,flag)
	go ReceiveMultiClientData(w, r, flag)
	//go TryReceiveOne(w, r, flag)
	<-flag
}

func NewWebServer(port string) {
	//return
	http.Handle("/", &web_server{})
	port = "0.0.0.0:" + port
	sglog.Info("start web server.listen port:", port)

	http.ListenAndServe(port, nil)
}

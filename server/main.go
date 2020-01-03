package main

import (
	"server/src/httpHandle"
	"github.com/coderguang/GameEngine_go/sgnet/sghttp"
	"log"

	"github.com/coderguang/GameEngine_go/sgcmd"
	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgserver"
)


func TestPost(cmd []string){
	filename:="./log/20200103.log"
	if len(cmd)>1{
		filename=cmd[1]
	}
	url:="http://localhost:7000"
	_,err:=sghttp.PostFile(filename,url,"sg")
	if err!=nil{
		sglog.Error("post file error",err);
	}
}




func main() {
	sgserver.StartServer(sgserver.ServerTypeLog, "debug", "./log/", log.LstdFlags, true)

	sglog.Info("start YY listen server...")
	go httpHandle.NewWebServer("7000");

	sgcmd.RegistCmd("TestPost","[\"TestPost\"]",TestPost)

	sgcmd.StartCmdWaitInputLoop()
}

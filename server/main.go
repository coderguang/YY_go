package main

import (
	"server/src/httpHandle"
	"log"

	"github.com/coderguang/GameEngine_go/sgcmd"
	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgserver"
)

func main() {
	sgserver.StartServer(sgserver.ServerTypeLog, "debug", "./log/", log.LstdFlags, true)

	sglog.Info("start YY listen server...")


	go httpHandle.NewWebServer("7000");

	sgcmd.StartCmdWaitInputLoop()
}

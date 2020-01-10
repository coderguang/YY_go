package main

import (
	"github.com/coderguang/GameEngine_go/sgnet/sghttp"
	"github.com/coderguang/GameEngine_go/sgthread"
	"io/ioutil"
	"log"
	"server/src/httpHandle"

	"github.com/coderguang/GameEngine_go/sgcmd"
	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgserver"
)

func TestPost(cmd []string) {
	filename := "./log/20200103.log"
	if len(cmd) > 1 {
		filename = cmd[1]
	}
	url := "http://localhost:7000"
	resp, err := sghttp.PostFile(filename, url, "sg")
	if err != nil {
		sglog.Error("post file error", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	sglog.Info("return str:", string(body))
}

func TestMultiPost(cmd []string) {
	fileMap := make(map[string]string)
	fileMap["f1"] = "./log/20200103.log"
	fileMap["f2"] = "./log/20200102.log"
	fileMap["pic"] = "./log/123.png"

	keyMap := make(map[string]string)
	keyMap["k1"] = "k1111"
	keyMap["k2"] = "k2222"
	keyMap["k3"] = "k3333"

	url := "http://localhost:7000"
	resp, err := sghttp.PostMultiFormFile(url, fileMap, keyMap)
	if err != nil {
		sglog.Error("post file error", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	sglog.Info("return str:", string(body))
}

func main() {
	sgserver.StartServer(sgserver.ServerTypeLog, "debug", "./log/", log.LstdFlags, true)

	sglog.Info("start YY listen server...")
	go httpHandle.NewWebServer("7000")

	sgcmd.RegistCmd("TestPost", "[\"TestPost\"]", TestPost)
	sgcmd.RegistCmd("TestMultiPost", "[\"TestMultiPost\"]", TestMultiPost)

	sgthread.SleepBySecond(2)

	TestMultiPost([]string{})

	sgcmd.StartCmdWaitInputLoop()
}

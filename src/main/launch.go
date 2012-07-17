package main

import (
	"common"
	"event"
	"fmt"
	"log"
	"os"
	"paas"
	"path/filepath"
	//"util"
)

func main() {
	path, err := filepath.Abs(os.Args[0])
	if nil != err {
		fmt.Println(err)
		return
	}
	common.Home, _ = filepath.Split(path)
	common.InitLogger()
	common.InitConfig()
	event.Init()
	var gae paas.GAE
	gae.Init()
	common.LoadRootCA()
	//	var req event.HTTPRequestEvent
	//	req.SetHash(1)
	//	req.Url = "https://twitter.com"
	//	req.Method = "GET"
	//	conn, err := gae.GetRemoteConnection(nil)
	//	err,res  := conn.Request(nil, &req)
	//	fmt.Println(string(res.(*event.HTTPResponseEvent).Content.Bytes()))
	log.Println("=============Start GSnova " + common.Version + "=============")
	addr, exist := common.Cfg.GetProperty("LocalServer", "Listen")
	if !exist {
		log.Fatalln("No config [LocalServer]->Listen found")
	}
	startLocalProxyServer(addr)
}
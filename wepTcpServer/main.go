package main

import (
	"wepTcpServer/start"
)

func main() {
	//启动udp,开启协程
	go start.StartUdp("0.0.0.0", 9090)
	start.Start("0.0.0.0:8080")

}

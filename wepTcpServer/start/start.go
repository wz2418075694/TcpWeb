package start

import (
	"log"
	"net"
	"wepTcpServer/server"
)

// 监听地址+端口，开启用户协程
func Start(addr string) {

	//监听端口,传入协议+地址
	Listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("服务器启动失败！", err)
	}
	defer Listener.Close()
	//启动成功
	log.Printf("服务器启动成功！正在监听端口:%s", addr)
	log.Println("等待客户端连接>>>")
	//循环接受客户端连接
	for {

		//等待连接，卡住
		conn, err := Listener.Accept()
		//RemoteAddr()方法可以获取客户端的地址+端口
		addr := conn.RemoteAddr()
		if err != nil {
			log.Printf("客户端:[%s]连接失败！", addr)
			continue
		}
		//开启协程
		go server.HandleClient(conn)

	}
}

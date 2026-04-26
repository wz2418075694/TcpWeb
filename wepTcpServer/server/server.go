package server

import (
	"bufio"
	"log"
	"net"
	"os"
	"path/filepath"
	"wepTcpServer/config"
	"wepTcpServer/send"
)

func init() {
	//设置日志打印格式,LstdFlags日志标准格式
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	//创建文件目录,os.MkdirAll()如果文件夹不存在就创建，如果存在就什么都不做
	_ = os.MkdirAll(config.FileDir, 0755)
	//mk, _ := os.Getwd()
	//fmt.Println("工作目录:", mk)

}

// 处理单个客户端
func HandleClient(conn net.Conn) {
	
	//关闭连接
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	log.Printf("客户端:[%s]连接成功！", addr)
	//创建读取器,读取客户端发送过来的消息
	reader := bufio.NewReader(conn)

	for {
		//用循环阻塞等待客户端发送消息
		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("客户端:[%s]断开连接！", addr)
			return
		}
		//去除掉末尾的换行符号
		cmd = cmd[0 : len(cmd)-1]
		//log.Printf("客户端[%s]发送的信息是:%s", addr, cmd)
		//ls,查看文件列表
		if cmd == "ls" {
			log.Printf("客户端输入%s命令", cmd)
			send.FileDir(conn)
			continue
		}
		//download filename
		//先截取download+空格
		if len(cmd) > 9 && cmd[0:9] == "download " {
			filename := cmd[9:]
			//fmt.Println(filename)
			log.Printf("客户端:[%s]请求下载文件:[%s]\n", addr, filename)
			//拼接文件路径
			filePath := filepath.Join(config.FileDir, filename)
			send.FileCon(conn, filePath)
			continue
		}
		_, _ = conn.Write([]byte("请输入正确的命令！\n"))
	}

}

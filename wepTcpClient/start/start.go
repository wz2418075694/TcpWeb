package start

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"wepTcpClient/config"
	"wepTcpClient/receive"
)

func Start(tcpAddr string) {
	//连接客户端
	conn, err := net.Dial("tcp", tcpAddr)
	if err != nil {
		//log.Fatal("客户端连接失败！", err)
		fmt.Println("客户端连接失败!")
		return
	}
	//用defer断开连接
	defer conn.Close()

	go receive.ReFileMessage(conn)

	//file, err := os.Create("1.txt")
	//file.Close()
	//连接成功
	fmt.Println("客户端和服务器连接成功！")
	fmt.Println("键盘录入命令，回车发送！")
	fmt.Println("-------------------------")
	fmt.Println("查看目录:ls")
	fmt.Println("下载文件:download 文件名")
	fmt.Println("-------------------------")

	//bufio.NewReader创建一个读取器
	//os.Stdin系统标准输入
	reader := bufio.NewReader(os.Stdin)
	//发送上边键盘传入的数据
	fmt.Println("请输入命令>>>")
	for {
		if config.ISDownloading() {
			fmt.Println("正在下载文件,请稍后！")
			continue
		}
		//ReadString()方法读取字符串，遇到某个字符串终止
		//fmt.Print(">")
		cmd, _ := reader.ReadString('\n')
		//去掉两边的空格
		cmd = strings.TrimSpace(cmd)
		//检查命令是否是空,空的话，跳过这次循环
		if cmd == "" {
			continue
		}
		//发送命令的时候，要加个换行符，否则接收器检测不到\n，他会一直进行阻塞
		_, err := conn.Write([]byte(cmd + "\n"))
		if err != nil {
			log.Println("发送数据失败！", err)
			return
		}

	}
}

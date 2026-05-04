package start

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"wepTcpClient/config"
	"wepTcpClient/receive"
	"wepTcpClient/send"
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
	fmt.Println("上载文件:upload 文件名")
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
		//使用strings.TrimSplace()函数，去掉两边的空格
		cmd = strings.TrimSpace(cmd)
		//检查命令是否是空,空的话，跳过这次循环
		if cmd == "" {
			continue
		}
		//使用strings.HasPrefix()判断字符串的前缀知否是上传
		if strings.HasPrefix(cmd, "upload") {
			//校验通过，启动上载
			filename := cmd[7:]
			//拼接路径
			filePath := filepath.Join(config.Dir, filename)
			//启动协程不影响TCP连接。
			go send.Upload(filePath)
			continue

		}
		//下载和查看文件夹的命令就用TCP连接发送出去
		//发送命令的时候，要加个换行符，否则接收器检测不到\n，他会一直进行阻塞
		_, err := conn.Write([]byte(cmd + "\n"))
		if err != nil {
			log.Println("发送数据失败！", err)
			return
		}

	}
}

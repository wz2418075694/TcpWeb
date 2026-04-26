package receive

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"path/filepath"
	"strconv"
	"strings"
	"wepTcpClient/MD5"
	"wepTcpClient/config"
)

// 接收消息
func ReFileMessage(conn net.Conn) {
	//创建接收器
	reader := bufio.NewReader(conn)
	//循环阻塞服务器发送消息
	for {

		//解决粘包问题，正确完整的读取协议，遇到\n就终止
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		//去掉末尾的换行符
		message = strings.TrimSpace(message)
		//如果传递的是协议
		if len(message) > 5 && message[:5] == "FILE:" {
			//用Split()拆分字符串,把文件传输协议按 ：拆成4个部分,
			//FILE:文件名:文件大小:MD5
			parts := strings.Split(message, ":")
			//拿到文件名字和文件的大小
			filename := parts[1]
			filesize := parts[2]
			//将字符串转换为int64类型
			size, _ := strconv.ParseInt(filesize, 10, 64)
			serverMD5 := parts[3]
			//下载
			downloadFile(conn, filename, size)

			//计算下载后文件的MD5
			//拼接路径
			filePath := filepath.Join(config.Dir, filename)
			clientMD5, err := MD5.GetMD5(filePath)
			if err != nil {
				fmt.Println(err)
				return
			}
			flag := MD5.CheakMD5(serverMD5, clientMD5)
			if flag {
				fmt.Printf("MD5校验通过！ 服务端:[%s] 客户端:[%s]\n", serverMD5, clientMD5)
			} else {
				fmt.Printf("MD5校验失败！ 服务端:[%s] 客户端:[%s]\n", serverMD5, clientMD5)
			}

			fmt.Println("请输入命令>>>")

		}
		if len(message) >= 3 && message[:3] == "DIR" {
			PrintIDr(conn)
			fmt.Println("请输入命令>>>")

		}

	}

	//作废，没有解决粘包问题
	//创建缓冲区，buffer英语是缓冲的意思
	//buf := make([]byte, 1024)
	//for {
	//
	//	conn.Read()
	//	阻塞，将返回的数据放入buf中，n是返回的字节数,
	//		tcp中有两种情况会返回错误，1.
	//	网络出错，2.
	//	服务器断开连接
	//	n, err := conn.Read(buf)
	//	if err != nil {
	//		fmt.Println("服务器断开连接！")
	//		return
	//	}
	//	buf[:n]
	//	只将前n个字符进行转换，如果是string(buf)
	//	是转换整个篮子，这样肯定是不行的
	//
	//	取消息，消息分两部分，一部分是要下载文件，一个是传递的文件目录
	//	message := string(buf[:n])
	//	//看是否是文件传输协议FILE:filename:fileSize
	//	if message[:5] == "FILE:" {
	//		//用Split()拆分字符串,把文件传输协议按 ：拆成3个部分
	//		parts := strings.Split(message, ":")
	//		//拿到文件名字和文件的大小
	//		filename := parts[1]
	//		filesize := parts[2]
	//		size, _ := strconv.ParseInt(filesize, 10, 64)
	//		//下载
	//		downloadFile(conn, filename, size)
	//		fmt.Println("请输入命令>>>")
	//
	//	} else { //传递的是文件目录
	//		content := string(buf[:n])
	//		//打印目录
	//		fmt.Printf(content)
	//		fmt.Println("请输入命令>>>")
	//
	//	}
	//
	//}
}

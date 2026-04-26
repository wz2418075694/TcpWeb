package send

import (
	"fmt"
	"log"
	"net"
	"os"
	"wepTcpServer/config"
)

func FileDir(conn net.Conn) {

	//发送协议头，告诉客户端发送的目录
	_, _ = conn.Write([]byte(fmt.Sprintf("DIR\n")))

	//os.ReadDir()入参是目录，返回值是目录项切片，每一个元素代表一个文件或者目录
	files, err := os.ReadDir(config.FileDir)
	if err != nil {
		//log.Println("读取目录失败！")
		_, _ = conn.Write([]byte("读取目录失败！"))
		return
	}
	//遍历目录项切片，把它返回给客户端
	result := "------------------\n"
	for _, file := range files {
		//判断是否是文件
		if !file.IsDir() {
			result += file.Name() + "\n"
		}
	}
	result += "------------------\n"
	//将文件目录发送过去
	_, err = conn.Write([]byte(result))
	if err != nil {
		fmt.Println("发送失败")
	}
	//发送结束符号
	_, err = conn.Write([]byte("END\n"))
	if err != nil {
		fmt.Println("发送失败")
	}
	log.Println("目录发送完成")

}

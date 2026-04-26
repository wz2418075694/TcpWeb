package receive

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func PrintIDr(conn net.Conn) {
	fmt.Println("开始打印目录")

	//创建读取器，一行一行的读
	reader := bufio.NewReader(conn)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("服务器断开连接！")
			} else {
				fmt.Println("非EOF错误！")
			}
			return
		}
		//去掉换行
		str = strings.TrimSpace(str)

		if str == "END" {
			break
		}
		fmt.Println(str)
	}

	//for {

	//buf := make([]byte, 1024*1024)
	////接受目录内容
	//n, err := conn.Read(buf)
	//if err != nil {
	//	if err == io.EOF {
	//		fmt.Println("服务器断开连接！")
	//	}
	//	fmt.Println("非EOF的错误！")
	//}
	//if {
	//
	//}
	//content := string(buf[:n])
	//fmt.Println(content)
	//}

}

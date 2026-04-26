package send

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

//发送文件的步骤
//1.先找文件
//2.打开文件
//3.发送协议头，告诉客户端我要发文件了，发送FILE:文件名字:文件大小:MD5
//4.发送文件

func FileCon(conn net.Conn, filePath string) {

	//1.这段删了，直接在函数外面就把路径找好
	//拼接路径
	//filePath := filepath.Join(config.FileDir, filename)

	//2.
	//os.Open()打开文件,只读，记得判断错误，记得关闭文件
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("打开文件失败！")
		_, _ = conn.Write([]byte("文件不存在！"))
		return
	}
	//关闭
	defer file.Close()
	//得到文件信息
	info, _ := file.Stat()
	fileSize := info.Size()
	filename := info.Name()
	MD5, err := GetMD5(filePath)
	if err != nil {
		log.Println(err)
		return
	}

	//3.将这些信息发送给客户端,
	//fmt.Println(filename)
	//发送协议头
	_, err = conn.Write([]byte(fmt.Sprintf("FILE:%s:%d:%s\n", filename, fileSize, MD5)))
	if err != nil {
		log.Println("发送协议失败！")
		return
	}

	//4.发送内容
	//创建一个缓冲区
	buf := make([]byte, 1024*1024)
	//创建一个循环，去一直读，直到读完
	for {
		//读文件放到篮子里面，两个返回值，n是字节数，err是是否出错(文件损坏)/是否读完了
		n, err := file.Read(buf)
		if n > 0 {
			_, _ = conn.Write(buf[:n])
		}

		//文件读完，EOF是一个全局变量，在io包里面，类型是error
		if err == io.EOF {
			log.Printf("文件:[%s]发送结束！", filename)
			return
		}
	}
}

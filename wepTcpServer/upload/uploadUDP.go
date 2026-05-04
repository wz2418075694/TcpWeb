package upload

import (
	"log"
	"net"
	"os"
)

func FileUpload(conn net.PacketConn, filePath string) {

	//创建文件
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("创建文件失败！", err)
	}
	//关闭文件
	defer file.Close()
	//创建缓冲区
	buf := make([]byte, 1024)

	//写入文件,用conn.ReadFrom()读取文件，用conn.WriteTo()恢复数据
	for {
		//从缓冲区里面读取,如果没有数据会一直阻塞
		n, Addr, err := conn.ReadFrom(buf)
		//udp不会返回io.EOF
		if err != nil {
			log.Println(err)
			break
		}
		if n == 3 && string(buf[:3]) == "END" {
			_, err := conn.WriteTo([]byte("服务器接受完成"), Addr)
			if err != nil {
				log.Println("给客户端发送消息失败！", err)
			}
			break
		}
		//写入文件
		_, err2 := file.Write(buf[:n])
		if err2 != nil {
			log.Println("写入文件失败", err2)
		}
		//UDP不能使用io.Copy()这种方式，UDP不是数据流，而是数据包，必须手写结束语句
	}
	//log.Println("客户端UDP链接发送文件完成！")
}

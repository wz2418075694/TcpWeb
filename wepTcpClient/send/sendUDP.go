package send

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"wepTcpClient/MD5"
)

func Upload(filepath string) {

	//连接，一般都用这种方式连接UDP,不用net.DialUDP()
	conn, err := net.Dial("udp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println("UDP连接失败！")
	}
	//用完记得关闭连接
	defer conn.Close()

	//UDP直接发
	fmt.Println("开始上传文件！")

	//先发送协议，upload:文件名字:文件大小:文件MD5
	//打开文件
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("打开文件失败！", err)
		return
	}
	//记得关闭文件
	defer file.Close()

	//获取文件的信息
	info, _ := file.Stat()
	filename := info.Name() //名字
	filesize := info.Size() //大小 int64类型
	fileMD5, err := MD5.GetMD5(filepath)
	if err != nil {
		fmt.Println("获取MD5失败！", err)
		return
	}

	//发送协议，记得用\n收尾，防止粘包
	_, err2 := conn.Write([]byte(fmt.Sprintf("upload:%s:%d:%s\n", filename, filesize, fileMD5)))
	if err2 != nil {
		fmt.Println("协议头发送失败！", err2)
		return
	}

	//开始发送数据
	//创建缓冲区
	buf := make([]byte, 1024)
	//sendByte := int64(0)
	//直接发,不等待，不确认
	for {
		n, err := file.Read(buf) //文件读取完成了err=EOF
		if err != nil {
			if err == io.EOF {
				//发送数据结束标志
				_, _ = conn.Write([]byte("END"))
				break
			}
			//非io.EOF的错误
			fmt.Println("数据搬迁失败！", err)
			return
		}

		_, err2 := conn.Write(buf[:n])
		if err2 != nil {
			log.Println("发送数据失败！", err2)
			return
		}
		//sendByte += int64(n2)
		////上载进度
		//percentage := (float64(sendByte) / float64(filesize)) * 100
		//fmt.Printf("已上载:%.1f%%\n", percentage)
	}
	fmt.Println("数据上载完成(不保证数据完整收到)！")

}

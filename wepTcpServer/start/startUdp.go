package start

import (
	"bufio"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"wepTcpServer/MD5"
	"wepTcpServer/config"
	"wepTcpServer/upload"
)

func init() {
	//创建目录用于存放客户端上传上来的文件
	err := os.MkdirAll(config.ClientFileDir, 0755)
	if err != nil {
		log.Printf("创建客户端上传文件的目录失败！:%v", err)
	}
}

//开启UDP端口
//循环接受小片的数据
//将数据写入文件中

func StartUdp(ip string, port int) {
	udpAddr := &net.UDPAddr{
		//将字符串解析成IP类型.
		IP:   net.ParseIP(ip),
		Port: port,
	}
	//go中推荐这种写法，虽然net.Listen()可以监听UDP
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalf("UDP监听失败:%v", err)
	}
	defer conn.Close()
	//创建一个读取器，用于接受协议
	reader := bufio.NewReader(conn)
	//循环阻塞UDP消息，遇到\n为止
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		//去除掉末尾的换行符号，不然发的MD5后面有\n
		str = strings.TrimSpace(str)
		parts := strings.Split(str, ":")
		filename := parts[1]
		filesize := parts[2]
		log.Println("文件大小：", filesize)
		//将字符串转换为整形
		//fileSize, _ := strconv.ParseInt(size, 10, 64)/

		//先拼接路径
		filePath := filepath.Join(config.ClientFileDir, filename)
		upload.FileUpload(conn, filePath)
		log.Println("上载完成！")

		//验证MD5是否一致
		fileMD5 := parts[3]
		//计算上载文件的MD5，然后进行比较
		calFileMD5, err := MD5.GetMD5(filePath)
		if err != nil {
			log.Println("计算文件的MD5失败！", err)
		}
		if fileMD5 == calFileMD5 {
			log.Printf("上载文件MD5校验通过！服务端:[%s] 客户端:[%s]\n", calFileMD5, fileMD5)
		} else {
			log.Printf("上载文件MD5校验失败！服务端:[%s] 客户端:[%s]\n", calFileMD5, fileMD5)
		}

	}
}

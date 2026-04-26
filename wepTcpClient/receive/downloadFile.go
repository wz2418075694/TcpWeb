package receive

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"wepTcpClient/config"
)

// 这个函数用于下载即将过来的这个文件
// 1.创建空的文件,
// 2.创建缓冲区
// 3.读数据，循环读完
func downloadFile(conn net.Conn, filename string, filesize int64) {

	fmt.Println("开始下载文件！")
	//开启下载开关
	config.SetDownload(true)

	//1.创建
	fmt.Printf("文件名字:[%s]  文件大小:[%.2f]KB\n", filename, float64(filesize)/1024)
	//拼接
	path := filepath.Join(config.Dir, filename)
	//创建文件，用于数据存储
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("创建文件失败")
		return
	}
	//记得关闭文件，上边的创建就相当于打开文件
	defer file.Close()

	//2.篮子
	buf := make([]byte, 1024*1024)
	receive := int64(0)

	//3.接收
	//循环接受然后写入
	fmt.Println("-------------------------")
	for receive < filesize {
		//设置如果5秒中没有数据读入，就结束循环
		//_ = conn.SetReadDeadline(time.Now().Add(time.Second * 5))
		n, err := conn.Read(buf)
		//fmt.Println(n)
		//如果有数据来
		if n > 0 {
			n, err := file.Write(buf[:n])
			receive += int64(n)
			if err != nil {
				fmt.Println("文件写入失败")
				return
			}
			//下载进度
			total := 20 //
			str := ""
			result := (float64(receive) / float64(filesize)) * 100
			//算当前多少块
			current := int(result / 100 * float64(total))
			//计算字符串
			for i := 1; i <= current; i++ {
				str += "♥"
			}
			for i := current + 1; i <= total; i++ {
				str += " "
			}
			fmt.Printf("已下载:[%s] %.1f%%\n", str, result)
		}

		//这样检测数据发完了，是不行的，因为conn.Read()是在监听这条线，没有数据的时候，连接正常一直阻塞等待
		//err有两种情况 1.对方主动断开，返回err=io.EOF 2.连接异常断开，返回非EOF的错误

		//if err == io.EOF {
		//	fmt.Printf("文件[%s]下载完成\n", filename)
		//	return
		//}

		if err != nil {
			if err == io.EOF {
				fmt.Println("服务器断开连接！")
				return
			}

			fmt.Println("非EOF的错误！")
			return
		}

	}
	fmt.Println("-------------------------")

	fmt.Printf("文件:[%s]下载完成\n", filename)
	//开关关了
	config.SetDownload(false)

}

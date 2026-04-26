package main

import (
	"log"
	"os"
	"wepTcpClient/config"
	"wepTcpClient/start"
)

func init() {
	//设置日志打印格式log.LstdFlags是标准格式，Lshortfile短文件名字
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	//查看当前的工作目录
	//wd, _ := os.Getwd()
	//fmt.Printf("当前的工作目录是:%s", wd)

	//创建存储数据文件的目录
	_ = os.MkdirAll(config.Dir, 0755)
}
func main() {

	start.Start("0.0.0.0:8080")

}

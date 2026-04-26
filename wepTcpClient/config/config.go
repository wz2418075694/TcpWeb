package config

//这是一个公共的工具包

// 存储数据的目录的路径，当前的工作目录是wepTcpClient
const (
	Dir = "./contentDir"
)

var isDownloading bool = false

// 设置状态
func SetDownload(flag bool) {
	isDownloading = false
}

// 查看是否正在下载
func ISDownloading() bool {
	return isDownloading
}

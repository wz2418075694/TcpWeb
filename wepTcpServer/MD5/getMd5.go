package MD5

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

//1.打开文件
//2.创建MD5计算器
//3.把文件喂给计算器
//4.去到结果

func GetMD5(filePath string) (strMD5 string, err error) {

	//打开文件
	file, err := os.Open(filePath)
	if err != nil {
		//log.Println("打开文件失败！")
		return "", err
	}
	//记得关闭文件
	defer file.Close()

	//创建MD5计算机
	hash := md5.New()

	//把文件喂给它，这里使用io.Copy(目的，来源),它自动搬运所有的数据，它的使用方法是：
	//1.这里的MD5的计算。
	//2.TCP文件的传输
	_, err = io.Copy(hash, file)
	if err != nil {
		//log.Println("io.Copy搬运出错！")
		return "", err
	}

	//拿结果,标准写法，nil表示不需要追加额外数据
	resultByte := hash.Sum(nil)
	//把这个二进制转化为16进制的字符串
	resultHexString := hex.EncodeToString(resultByte)

	return resultHexString, nil
}

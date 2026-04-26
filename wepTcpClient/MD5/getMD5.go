package MD5

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func GetMD5(filePath string) (strMD5 string, err error) {

	//打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	//创建MD5计算器
	hash := md5.New()
	//用io.Copy()将数据搬运到MD5计算器
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	//取结果
	resultByte := hash.Sum(nil)
	//将结果转化为16进制的字符串
	resultHexString := hex.EncodeToString(resultByte)

	return resultHexString, nil
}

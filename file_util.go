package util

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

// 判断资源包是否存在
func FilePathExists(filename string) (bool, string) {
	file_path := filename
	_, err := os.Stat(file_path)
	if err == nil {
		return true, file_path
	} else {
		return false, "FilePathExists:::NotFound " + file_path
	}
}

func WriteToFile(fullname string, content string) bool {
	filePath := path.Dir(fullname)
	exists, _ := FilePathExists(filePath)
	if exists == false {
		os.MkdirAll(filePath, os.ModePerm)
	}
	bcontent := []byte(content)
	//将指定内容写入到文件中
	err := ioutil.WriteFile(fullname, bcontent, 0666)
	if err != nil {
		log.Fatalln("ioutil WriteFile error: ", err)
		return false
	}

	return true
}
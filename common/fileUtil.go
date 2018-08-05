package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/debug"
	"strings"
)

func LoadFileByName(filePath string) ([]byte, error) {
	file, err := os.Open(filePath) // For read access.
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}
	return data, err
}
func WriteFileByName(filePath string, fileContent []byte) error {
	return ioutil.WriteFile(filePath, fileContent, 0644)
}
func SafeCall(f func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(string(debug.Stack()))
		}
	}()
	f()
}

// folder tool
func EnsureFolder(path string) {
	exist, err := PathExists(path)
	if err != nil {
		fmt.Printf("get dir error %v \n", err)
		return
	}

	if exist {
		fmt.Printf("has dir %v \n", path)
	} else {
		fmt.Printf("no dir %v \n", path)
		// 创建文件夹
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed %v \n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
}
func ClearFolder(path string) {
	exist, err := PathExists(path)
	if err != nil {
		fmt.Printf("get dir error %v \n", err)
		return
	}

	if exist {
		fmt.Printf("has dir![%v]\n", path)
		os.RemoveAll(path)
	}
	fmt.Printf("no dir![%v]\n", path)
	// 创建文件夹
	err = os.Mkdir(path, os.ModePerm)
	if err != nil {
		fmt.Printf("mkdir failed![%v]\n", err)
	} else {
		fmt.Printf("mkdir success!\n")
	}
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func ParserFileNameByPath(fullpath string) (name string, suffix string) {
	fullpath = strings.Replace(fullpath, "\\", "/", -1)
	nameWithSuffix := ""
	tmpList := strings.Split(fullpath, "/")
	if len(tmpList) <= 1 {
		nameWithSuffix = fullpath
	} else {
		nameWithSuffix = tmpList[len(tmpList)-1]
	}

	fileNameList := strings.Split(nameWithSuffix, ".")
	suffix = ""
	name = nameWithSuffix

	if len(fileNameList) > 1 {
		suffix = fileNameList[len(fileNameList)-1]
		name = fileNameList[0]
	}

	return
}

package common

import (
	"io/ioutil"
	"log"
	"os"
	"runtime/debug"
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

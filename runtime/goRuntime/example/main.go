package main

import (
	"fmt"
	"github.com/Blizzardx/ConfigProtocol/runtime/goRuntime"
)

type ItemConfig struct {
	Content []*ItemConfigInfo
}
type ItemConfigInfo struct {
	Id   int32
	Name string
}

type BoxConfig struct {
	Content map[int32]*BoxConfigInfo
}
type BoxConfigInfo struct {
	Id   int32
	Name string
}

func main() {
	content := &ItemConfig{}
	err := goRuntime.LoadConfig(content)
	fmt.Println(err)
	fmt.Println(content)
}

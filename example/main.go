package main

import (
	"fmt"
	"github.com/Blizzardx/ConfigProtocol/define"
	"github.com/Blizzardx/ConfigProtocol/tool/exportTool"
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
	//content := &ItemConfig{}
	//err := goRuntime.LoadConfig(content)
	//fmt.Println(err)
	//fmt.Println(content)
	var exportList []*exportTool.ExportTarget
	exportList = append(exportList, &exportTool.ExportTarget{Name: "server", Lan: define.SupportLan_Go})
	err := exportTool.ExportFile("config/BasicItem_Common.xlsx", "output", exportList, ".bytes")
	if nil != err {
		fmt.Println(err)
	}
}

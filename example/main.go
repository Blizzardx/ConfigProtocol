package main

import (
	"fmt"
	"github.com/Blizzardx/ConfigProtocol/define"
	"github.com/Blizzardx/ConfigProtocol/example/output/Go"
	"github.com/Blizzardx/ConfigProtocol/runtime/goRuntime"
	"github.com/Blizzardx/ConfigProtocol/tool/exportTool"
)

func main() {
	//gen()
	use()
}
func gen() {
	var exportList []*exportTool.ExportTarget
	exportList = append(exportList, &exportTool.ExportTarget{Name: "server", Lan: define.SupportLan_Go})
	exportList = append(exportList, &exportTool.ExportTarget{Name: "client", Lan: define.SupportLan_Csharp})

	err := exportTool.ExportFile("config/BasicItem_Common.xlsx", "output", exportList, ".bytes")
	fmt.Println(err)
}
func use() {
	content := &config.BasicItem_Common{}
	goRuntime.SetWorkspace("output")
	err := goRuntime.LoadConfig(content)
	fmt.Println(err)
	for _, v := range content.Content {
		//fmt.Printf("%v", v)
		fmt.Println(v)
	}
}

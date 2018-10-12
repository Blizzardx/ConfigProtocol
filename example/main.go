package main

import (
	"fmt"
	"github.com/Blizzardx/ConfigProtocol/define"
	"github.com/Blizzardx/ConfigProtocol/example/output/server/Go"
	"github.com/Blizzardx/ConfigProtocol/runtime/goRuntime"
	"github.com/Blizzardx/ConfigProtocol/tool/exportTool"
)

func main() {
	gen()
	use()
}
func gen() {
	var exportList []*exportTool.ExportTarget
	exportList = append(exportList, &exportTool.ExportTarget{Name: "server", Lan: define.SupportLan_Go, OutPutSuffix: ".bytes", PackageName: "config"})
	exportList = append(exportList, &exportTool.ExportTarget{Name: "client", Lan: define.SupportLan_Csharp, OutPutSuffix: ".bytes", PackageName: "config"})
	exportList = append(exportList, &exportTool.ExportTarget{Name: "httpServer", Lan: define.SupportLan_Java, OutPutSuffix: ".bytes", PackageName: "config"})
	exportList = append(exportList, &exportTool.ExportTarget{Name: "cocosClient", Lan: define.SupportLan_Json, OutPutSuffix: ".json", PackageName: "config"})

	//err := exportTool.ExportFile("config/BasicItem_Common.xlsx", "output", exportList)
	err := exportTool.ExportDirectory("config", "output", exportList)
	fmt.Println(err)
}
func use() {
	content := &config.BasicItem_Common{}
	goRuntime.SetWorkspace("output/server")
	err := goRuntime.LoadConfig(content)
	fmt.Println(err)
	for _, v := range content.Content {
		//fmt.Printf("%v", v)
		fmt.Println(v)
	}
}

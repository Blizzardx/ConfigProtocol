package exportTool

import "github.com/Blizzardx/ConfigProtocol/define"

const codeTemplate_Csharp = `// Generated by gen-tool
// DO NOT EDIT!
package {{.PackageName}}

type {{.ConfigName}} struct {
	Content []*{{.ConfigName}}Info
}
type {{.ConfigName}}Info struct {
{{range .FieldList}}
	{{.Name}} {{.Type}}
{{end}}
}
`

type genRuntimeCodeTool_Csharp struct {
}

func (self *genRuntimeCodeTool_Csharp) GenRuntimeCode(outputPath string, provision *ConfigDefine, enumInfo []*EnumDefine) error {
	return nil
}
func (self *genRuntimeCodeTool_Csharp) Name() define.SupportLan {
	return define.SupportLan_Csharp
}
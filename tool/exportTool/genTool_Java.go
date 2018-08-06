package exportTool

import "github.com/Blizzardx/ConfigProtocol/define"

const codeTemplate_Java = `// Generated by gen-tool
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

type genRuntimeCodeTool_Java struct {
}

func (self *genRuntimeCodeTool_Java) GenRuntimeCode(outputPath string, provision *ConfigDefine, enumInfo []*EnumDefine) error {
	return nil
}
func (self *genRuntimeCodeTool_Java) Name() define.SupportLan {
	return define.SupportLan_Java
}
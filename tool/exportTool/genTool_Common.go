package exportTool

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"text/template"
)

type ConfigDefine struct {
	PackageName      string
	ConfigName       string
	MapKeyType       string
	FieldList        []*ConfigFieldDefine
	EnumInfoTemplate string
	ImportPkg        []string
}
type ConfigFieldDefine struct {
	Name   string
	Type   string
	IsList bool
}
type EnumDefine struct {
	ConfigName   string
	EnumName     string
	EnumElemList []*EnumElemDefine
}
type EnumElemDefine struct {
	ConfigName    string
	EnumName      string
	EnumElemName  string
	EnumElemValue string
}

//根据模板生成代码
func generateCode(templateStr string, model interface{}, needFormat bool) (string, error) {

	var err error

	var bf bytes.Buffer

	tpl, err := template.New("Template").Parse(templateStr)
	if err != nil {
		return "", err
	}

	err = tpl.Execute(&bf, model)
	if err != nil {
		return "", err
	}

	if needFormat {
		if err = formatCode(&bf); err != nil {
			fmt.Println("format golang code err", err)
		}
	}

	return string(bf.Bytes()), nil
}

//格式化go文件
func formatCode(bf *bytes.Buffer) error {

	fset := token.NewFileSet()

	ast, err := parser.ParseFile(fset, "", bf, parser.ParseComments)
	if err != nil {
		return err
	}

	bf.Reset()

	err = (&printer.Config{Mode: printer.TabIndent | printer.UseSpaces, Tabwidth: 8}).Fprint(bf, fset, ast)
	if err != nil {
		return err
	}

	return nil
}

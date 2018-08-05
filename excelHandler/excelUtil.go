package excelHandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/Blizzardx/ConfigProtocol/define"
	"strings"
)

func ReadExcelFile(filePath string) ([][]string, error) {

	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var content [][]string
	sheets := xlsx.GetSheetMap()
	for _, sheetName := range sheets {
		table := xlsx.GetRows(sheetName)
		content = append(content, table...)
	}
	return content, nil
}

func ParserExcelToConfigProvision(content [][]string, fileName string) (*define.ConfigInfo, error) {
	if len(content) < 2 {
		return nil, errors.New("content error ")
	}
	// parser package name
	provision := &define.ConfigInfo{TableName: fileName}
	// global define at [0][0] pos
	if len(content[0]) < 1 {
		return nil, errors.New("content length error at config " + fileName)
	}
	globalDefineStr := content[0][0]
	provision.GlobalInfo = &define.ConfigHeadInfo{}
	err := json.Unmarshal([]byte(globalDefineStr), provision.GlobalInfo)
	if nil != err {
		return nil, errors.New("error on unmarshal global content ,it must be json " + globalDefineStr + " at config " + fileName)
	}
	for _, lineELem := range content[1] {
		fieldInfo := &define.ConfigFieldInfo{}
		if lineELem != "" {
			err := json.Unmarshal([]byte(lineELem), fieldInfo)
			if nil != err {
				return nil, errors.New("error on unmarshal field info,it must be json " + lineELem + " at config " + fileName)
			}
		}
		provision.LineInfo = append(provision.LineInfo, fieldInfo)
	}
	return provision, nil
}

func FixExcelFile(content [][]string) [][]string {
	var resultContent [][]string
	for index, line := range content {
		if index < 3 {
			continue
		}
		if len(line) > 0 && line[0] == "#" {
			continue
		}

		resultContent = append(resultContent, line)
	}
	return resultContent
}

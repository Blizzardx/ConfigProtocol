package exportTool

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Blizzardx/ConfigProtocol/common"
	"github.com/Blizzardx/ConfigProtocol/pbConfig"
	"strings"
	"time"
)

type ConfigSerializer_Json struct {
}

func (self *ConfigSerializer_Json) Export(configContent *config.ConfigTable) ([]byte, error) {
	if configContent.Type == config.ConfigType_typeList {
		// list
		var jsonRoot []interface{}
		for lineIndex, line := range configContent.Content {
			lineInfo, err := self.exportLine(configContent, line, lineIndex)
			if err != nil {
				return nil, err
			}
			jsonRoot = append(jsonRoot, lineInfo)
		}
		return json.Marshal(jsonRoot)
	} else {
		// map
		jsonRoot := map[string]interface{}{}

		for lineIndex, line := range configContent.Content {
			keyInfo, lineInfo, err := self.exportLineWithKey(configContent, line, lineIndex)
			if err != nil {
				return nil, err
			}
			jsonRoot[fmt.Sprint(keyInfo)] = lineInfo
		}
		return json.Marshal(jsonRoot)
	}
}
func (self *ConfigSerializer_Json) exportLine(configContent *config.ConfigTable, configLineInfo *config.ConfigLine, lineIndex int) (interface{}, error) {
	lineInfo := map[string]interface{}{}
	for index, fieldInfo := range configContent.FieldInfoList {
		cell := configLineInfo.Content[index]
		cellName := fieldInfo.Name
		cellValue, err := self.parserCell(fieldInfo.Type, cell, fieldInfo.IsList)
		if nil != err {
			return nil, errors.New(getTipMessage(TipMessageDefine_ErrorOnExportConfigParserError, lineIndex+4, index+1, err.Error()))
		}
		lineInfo[cellName] = cellValue
	}
	return lineInfo, nil
}
func (self *ConfigSerializer_Json) exportLineWithKey(configContent *config.ConfigTable, configLineInfo *config.ConfigLine, lineIndex int) (interface{}, interface{}, error) {
	lineInfo := map[string]interface{}{}
	var keyInfo interface{} = nil

	for index, fieldInfo := range configContent.FieldInfoList {
		cell := configLineInfo.Content[index]
		cellName := fieldInfo.Name
		cellValue, err := self.parserCell(fieldInfo.Type, cell, fieldInfo.IsList)
		if nil != err {
			return nil, nil, errors.New(getTipMessage(TipMessageDefine_ErrorOnExportConfigParserError, lineIndex+4, index+1, err.Error()))
		}
		lineInfo[cellName] = cellValue
		if configContent.KeyFieldName == fieldInfo.Name {
			keyInfo = cellValue
		}
	}
	if nil == keyInfo {
		return nil, nil, errors.New(getTipMessage(TipMessageDefine_ErrorOnExportConfigToJsonKeyNotFound, lineIndex+4, configContent.KeyFieldName))
	}
	return keyInfo, lineInfo, nil
}
func (self *ConfigSerializer_Json) parserCell(cellType config.FieldType, cellContent string, isList bool) (interface{}, error) {
	if isList {
		var list []interface{}
		listContent := strings.Split(cellContent, "|")
		for _, cellElement := range listContent {
			value, err := self.parserCellValue(cellType, cellElement)
			if nil != err {
				return nil, err
			}
			list = append(list, value)
		}
		return list, nil
	} else {
		return self.parserCellValue(cellType, cellContent)
	}
}

func (self *ConfigSerializer_Json) parserCellValue(cellType config.FieldType, cellContent string) (interface{}, error) {

	switch cellType {
	case config.FieldType_typeInt32:
		var value int32 = 0
		err := common.Parser_int32(cellContent, &value)
		return value, err
	case config.FieldType_typeInt64:
		var value int64 = 0
		err := common.Parser_int64(cellContent, &value)
		return value, err
	case config.FieldType_typeFloat32:
		var value float32 = 0
		err := common.Parser_float32(cellContent, &value)
		return value, err
	case config.FieldType_typeFloat64:
		var value float64 = 0
		err := common.Parser_float64(cellContent, &value)
		return value, err
	case config.FieldType_typeDateTime:
		var value time.Time
		err := common.Parser_dateTime(cellContent, &value)
		return value.Unix(), err
	}
	return cellContent, nil
}

package exportTool

import (
	"errors"
	"github.com/Blizzardx/ConfigProtocol/common"
	"github.com/Blizzardx/ConfigProtocol/define"
	"github.com/Blizzardx/ConfigProtocol/pbConfig"
	"github.com/Blizzardx/ConfigProtocol/tool/excelHandler"
	"github.com/golang/protobuf/proto"
	"strconv"
	"strings"
)

type ExportTarget struct {
	Name string
	Lan  define.SupportLan
}

type ConfigRunTimeCodeGenerator interface {
	GenRuntimeCode(outputPath string, provision *ConfigDefine, enumInfo []*EnumDefine) error
	Name() define.SupportLan
}

var codeGenToolStore = map[define.SupportLan]ConfigRunTimeCodeGenerator{}
var currentConfigEnumInfoList []*config.ConfigEnumInfo

func init() {
	// register
	codeGenToolStore[define.SupportLan_Go] = &genRuntimeCodeTool_Go{}
	codeGenToolStore[define.SupportLan_Csharp] = &genRuntimeCodeTool_Csharp{}
	codeGenToolStore[define.SupportLan_Java] = &genRuntimeCodeTool_Java{}
}
func ExportFile(filePath string, outputPath string, exportTargetList []*ExportTarget, outputFileSuffix string) error {
	currentConfigEnumInfoList = nil

	// parser file name
	fileName, _ := common.ParserFileNameByPath(filePath)

	// load excel file
	content, err := excelHandler.ReadExcelFile(filePath)
	if err != nil {
		return err
	}
	// parser file header
	provision, err := excelHandler.ParserExcelToConfigProvision(content, fileName)
	if err != nil {
		return err
	}
	// begin check provision
	err = checkConfigProvisionCorrect(provision)
	if err != nil {
		return err
	}

	// begin check content
	err = checkConfigContentCorrect(provision, content)
	if nil != err {
		return err
	}

	// begin export
	for _, exportTarget := range exportTargetList {
		tmpConfig, err := doExport(outputPath, provision, content, exportTarget)
		if err != nil {
			return errors.New("error on export config at target " + exportTarget.Name + " " + err.Error())
		}
		byteContent, err := proto.Marshal(tmpConfig)

		if err != nil {
			return errors.New("error on export config at target " + exportTarget.Name + " " + err.Error())
		}

		// ensure output path
		common.EnsureFolder(outputPath)

		// do export
		common.WriteFileByName(outputPath+"/"+provision.TableName+outputFileSuffix, byteContent)
	}
	return nil
}
func checkConfigProvisionCorrect(provision *define.ConfigInfo) error {
	if provision.GlobalInfo.TableType != "list" && provision.GlobalInfo.TableType != "map" {
		return errors.New("error table type " + provision.GlobalInfo.TableType + " table type must be 'list' or 'map' ")

	}
	if provision.GlobalInfo.TableType == "map" {
		isFindKeyField := false
		for _, field := range provision.LineInfo {
			if field.FieldName == provision.GlobalInfo.TableKeyFieldName {
				isFindKeyField = true
				break
			}
		}
		if !isFindKeyField {
			return errors.New("can't find key field by name " + provision.GlobalInfo.TableKeyFieldName)
		}
	}
	// read enum
	var err error
	currentConfigEnumInfoList, err = parserEnumList(provision.GlobalInfo.GlobalEnumDefine)
	if nil != err {
		return err
	}

	var fieldNameMap = map[string]int32{}
	for _, field := range provision.LineInfo {
		if field.FieldType == "" && field.FieldName == "" {
			continue
		}

		_, error := convertStrToFieldType(field.FieldType)
		if error != nil {
			// check is field in enum list
			if !checkFieldIsInEnum(field.FieldType) {
				return error
			}
		}
		if field.FieldName == "" {
			return errors.New("field named error : " + field.FieldName)
		}
		if _, ok := fieldNameMap[field.FieldName]; ok {
			return errors.New("more then one field named " + field.FieldName)
		}
		fieldNameMap[field.FieldName] = 1
	}
	return nil
}
func checkFieldIsInEnum(fieldType string) bool {
	for _, enum := range currentConfigEnumInfoList {
		if enum.Name == fieldType {
			return true
		}
	}
	return false
}
func checkFieldIsInEnumWithName(fieldType string) (bool, *config.ConfigEnumInfo) {
	for _, enum := range currentConfigEnumInfoList {
		if enum.Name == fieldType {
			return true, enum
		}
	}
	return false, nil
}
func checkFieldIsCorrectInEnum(fieldType string, value string) bool {
	for _, enum := range currentConfigEnumInfoList {
		if enum.Name == fieldType {
			// begin check value
			var tmpValue int32 = 0
			err := common.Parser_int32(value, &tmpValue)
			if nil != err {
				return false
			}
			for _, enumElement := range enum.Value {
				if enumElement.Value == tmpValue {
					return true
				}
			}
			return false
		}
	}
	return false

}
func checkConfigContentCorrect(provision *define.ConfigInfo, content [][]string) error {
	fixedContent := excelHandler.FixExcelFile(content)
	for rowIndex, rowContent := range fixedContent {
		for colIndex, contentCell := range rowContent {
			positionMark := " at col: " + strconv.Itoa(colIndex+1) + " row: " + strconv.Itoa(rowIndex+4)
			if colIndex >= len(provision.LineInfo) {
				// check error
				return errors.New("out of line range " + positionMark)
			}
			fieldProvisionInfo := provision.LineInfo[colIndex]

			if fieldProvisionInfo.FieldType == "" && fieldProvisionInfo.FieldName == "" {
				continue
			}

			if fieldProvisionInfo.IsList {
				// parser list
				contentCellList := strings.Split(contentCell, "|")
				for _, tmpCell := range contentCellList {
					// check type check range
					err := checkFieldTypeCorrect(fieldProvisionInfo.FieldType,
						tmpCell,
						fieldProvisionInfo.FieldValueRangeLimitMin,
						fieldProvisionInfo.FieldValueRangeLimitMax)
					if nil != err {
						return errors.New("error" + positionMark + " " + err.Error())
					}
				}
			} else {
				// check type check range
				err := checkFieldTypeCorrect(fieldProvisionInfo.FieldType,
					contentCell,
					fieldProvisionInfo.FieldValueRangeLimitMin,
					fieldProvisionInfo.FieldValueRangeLimitMax)
				if nil != err {
					return errors.New("error" + positionMark + " " + err.Error())
				}
			}
		}
	}
	return nil
}
func checkFieldTypeCorrect(fieldTypeStr string, content string, minValue string, maxValue string) error {

	fieldType, err := convertStrToFieldType(fieldTypeStr)
	if err != nil {
		if checkFieldIsCorrectInEnum(fieldTypeStr, content) {
			return nil
		}
		return errors.New("error on parser enum " + fieldTypeStr + " with unknown value " + content)
	}
	switch fieldType {
	case config.FieldType_typeInt32:
		var tmpValue int32 = 0
		err := common.Parser_int32(content, &tmpValue)
		if nil != err {
			return err
		}
		err = common.CheckValueLimit_int32(tmpValue, minValue, maxValue)
		if nil != err {
			return err
		}
		return nil
	case config.FieldType_typeInt64:
		var tmpValue int64 = 0
		err := common.Parser_int64(content, &tmpValue)
		if nil != err {
			return err
		}
		err = common.CheckValueLimit_int64(tmpValue, minValue, maxValue)
		if nil != err {
			return err
		}
		return nil
	case config.FieldType_typeFloat32:
		var tmpValue float32 = 0
		err := common.Parser_float32(content, &tmpValue)
		if nil != err {
			return err
		}
		err = common.CheckValueLimit_float32(tmpValue, minValue, maxValue)
		if nil != err {
			return err
		}
		return nil
	case config.FieldType_typeFloat64:
		var tmpValue float64 = 0
		err := common.Parser_float64(content, &tmpValue)
		if nil != err {
			return err
		}
		err = common.CheckValueLimit_float64(tmpValue, minValue, maxValue)
		if nil != err {
			return err
		}
		return nil
	case config.FieldType_typeBool:
		var tmpValue bool = false
		err := common.Parser_bool(content, &tmpValue)
		if nil != err {
			return err
		}
		return nil
	case config.FieldType_typeString:
		var tmpValue string = ""
		err := common.Parser_string(content, &tmpValue)
		if nil != err {
			return err
		}
		return nil
	}
	return errors.New("unknown field type " + fieldTypeStr)
}
func checkExportTarget(exportTarget *ExportTarget, fieldTarget string) bool {
	if fieldTarget == "" {
		return true
	}
	// parser field target
	targetList := strings.Split(fieldTarget, "|")
	if len(targetList) == 0 {
		return true
	}
	for _, target := range targetList {
		if target == exportTarget.Name {
			return true
		}
	}
	return false
}
func convertStrToFieldType(fileType string) (config.FieldType, error) {
	switch fileType {
	case "int32":
		return config.FieldType_typeInt32, nil
	case "int64":
		return config.FieldType_typeInt64, nil
	case "float32":
		return config.FieldType_typeFloat32, nil
	case "float64":
		return config.FieldType_typeFloat64, nil
	case "bool":
		return config.FieldType_typeBool, nil
	case "string":
		return config.FieldType_typeString, nil
	case "class":
		return config.FieldType_typeClass, nil
	case "time":
		return config.FieldType_typeDateTime, nil
	case "color":
		return config.FieldType_typeColor, nil
	default:
		return config.FieldType_typeInt32, errors.New("unknown field type " + fileType)
	}
}
func parserEnumList(enumList []string) ([]*config.ConfigEnumInfo, error) {

	var keyNameMap = map[string]int{}
	var result []*config.ConfigEnumInfo
	for _, enum := range enumList {
		tmpElem, err := parserEnumString(enum)
		if err != nil {
			return nil, err
		}
		if _, ok := keyNameMap[tmpElem.Name]; ok {
			return nil, errors.New("error on parser enum ,name repeated " + tmpElem.Name)
		}
		keyNameMap[tmpElem.Name] = 1
		result = append(result, tmpElem)
	}
	return result, nil
}
func parserEnumString(enum string) (*config.ConfigEnumInfo, error) {
	// parser enum name
	tmpStr := strings.Split(enum, ":")
	if len(tmpStr) != 2 {
		return nil, errors.New("error on parser enum " + enum)
	}
	result := &config.ConfigEnumInfo{}
	result.Name = tmpStr[0]
	tmpStr = strings.Split(tmpStr[1], "|")
	if len(tmpStr) <= 0 {
		return nil, errors.New("error on parser enum values " + enum)
	}
	var keyNameMap = map[string]int{}
	for _, tmpElem := range tmpStr {
		tmpCell := strings.Split(tmpElem, "=")
		if len(tmpCell) != 2 {
			return nil, errors.New("error on parser enum values " + enum + " at: " + tmpElem)
		}
		// check value type
		tmpCellEnum := &config.ConfigEnumElementInfo{}
		tmpCellEnum.Name = tmpCell[0]
		if tmpCellEnum.Name == "" || tmpCellEnum.Name == " " {
			return nil, errors.New("error on parser enum values " + enum + " at: " + tmpElem + " key can't empty")
		}
		if _, ok := keyNameMap[tmpCellEnum.Name]; ok {
			return nil, errors.New("key repeated " + enum + " at: " + tmpCellEnum.Name)
		}

		keyNameMap[tmpCellEnum.Name] = 1

		err := common.Parser_int32(tmpCell[1], &tmpCellEnum.Value)
		if nil != err {
			return nil, errors.New("error on parser enum values " + enum + " at: " + tmpElem + " value must be int32 " + tmpCell[1])
		}
		result.Value = append(result.Value, tmpCellEnum)
	}
	return result, nil
}
func convertPbEnum(configName string, pbEnumInfo []*config.ConfigEnumInfo) []*EnumDefine {
	var result []*EnumDefine
	for _, pbEnum := range pbEnumInfo {
		elem := &EnumDefine{ConfigName: configName, EnumName: pbEnum.Name}
		for _, pbEnumElem := range pbEnum.Value {
			elem.EnumElemList = append(elem.EnumElemList, &EnumElemDefine{
				ConfigName:    configName,
				EnumName:      pbEnum.Name,
				EnumElemName:  pbEnumElem.Name,
				EnumElemValue: strconv.Itoa(int(pbEnumElem.Value)),
			})
		}
		result = append(result, elem)
	}
	return result
}
func doExport(outputPath string, provision *define.ConfigInfo, content [][]string, exportTarget *ExportTarget) (*config.ConfigTable, error) {
	pbConfig := &config.ConfigTable{}

	if provision.GlobalInfo.TableType == "list" {
		pbConfig.Type = config.ConfigType_typeList

	} else if provision.GlobalInfo.TableType == "map" {
		pbConfig.Type = config.ConfigType_typeMap
		pbConfig.KeyFieldName = provision.GlobalInfo.TableKeyFieldName
	}

	var ignoreColIndex = map[int]int{}
	for index, field := range provision.LineInfo {
		if field.FieldName == "" && field.FieldType == "" {
			//
			ignoreColIndex[index] = 1
			continue
		}
		if !checkExportTarget(exportTarget, field.ExportTarget) {
			//
			ignoreColIndex[index] = 1
			continue
		}
		parameter := ""
		fieldType, err := convertStrToFieldType(field.FieldType)
		if err != nil {
			// check is enum
			if ok, enumInfo := checkFieldIsInEnumWithName(field.FieldType); ok {
				fieldType = config.FieldType_typeEnum
				parameter = enumInfo.Name
			}
		}
		pbConfig.FieldInfoList = append(pbConfig.FieldInfoList, &config.ConfigFieldInfo{Name: field.FieldName, Type: fieldType, IsList: field.IsList, Parameter: parameter})
	}
	fixedContent := excelHandler.FixExcelFile(content)
	for _, rowContent := range fixedContent {
		configLine := &config.ConfigLine{}
		for colIndex, contentCell := range rowContent {
			// check colIndex is in ignore list
			if _, ok := ignoreColIndex[colIndex]; ok {
				continue
			}
			configLine.Content = append(configLine.Content, contentCell)
		}
		pbConfig.Content = append(pbConfig.Content, configLine)
	}

	pbConfig.EnumInfoList = currentConfigEnumInfoList

	enumDefine := convertPbEnum(provision.TableName, currentConfigEnumInfoList)

	define := &ConfigDefine{PackageName: "config", ConfigName: provision.TableName}
	for index, field := range provision.LineInfo {
		if _, ok := ignoreColIndex[index]; ok {
			continue
		}
		define.FieldList = append(define.FieldList, &ConfigFieldDefine{Name: field.FieldName, Type: field.FieldType, IsList: field.IsList})
	}
	define.MapKeyType = ""
	if provision.GlobalInfo.TableType == "map" {
		for _, tmpField := range provision.LineInfo {
			if tmpField.FieldName == provision.GlobalInfo.TableKeyFieldName {
				define.MapKeyType = tmpField.FieldType
				break
			}
		}
	}
	// gen runtime code
	err := codeGenToolStore[exportTarget.Lan].GenRuntimeCode(outputPath, define, enumDefine)
	if nil != err {
		return nil, err
	}
	return pbConfig, nil
}

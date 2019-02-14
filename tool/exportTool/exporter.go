package exportTool

import (
	"errors"
	"github.com/Blizzardx/ConfigProtocol/common"
	"github.com/Blizzardx/ConfigProtocol/define"
	"github.com/Blizzardx/ConfigProtocol/tool/excelHandler"
	"image/color"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type ExportTarget struct {
	Name         string
	Lan          define.SupportLan
	OutPutSuffix string
	PackageName  string
}
type LoadedConfigInfo struct {
	Content   [][]string
	Provision *define.ConfigProvisionInfo
}
type ConfigRunTimeCodeGenerator interface {
	GenRuntimeCode(outputPath string, provision *ConfigDefine, enumInfo []*EnumDefine) error
	Name() define.SupportLan
}
type ConfigExportSerializer interface {
	Export(configContent *define.ConfigTable) ([]byte, error)
}

var codeGenToolStore = map[define.SupportLan]ConfigRunTimeCodeGenerator{}
var configSerializerStore = map[define.SupportLan]ConfigExportSerializer{}
var currentConfigEnumInfoList []*define.ConfigEnumInfo
var loadedConfigFileInfoStore = map[string]*LoadedConfigInfo{}
var workSpace string = ""

//导出所有文件，指定文件夹内的所有文件
func ExportDirectory(directory string, outputPath string, exportTargetList []*ExportTarget) error {
	loadedConfigFileInfoStore = map[string]*LoadedConfigInfo{}
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Println("error on check directory  " + err.Error())
		return err
	}
	errStr := ""
	workSpace = directory

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		err := doExportFile(directory+"/"+file.Name(), outputPath, exportTargetList)
		if err != nil {
			errStr += directory + "/" + file.Name() + " : " + err.Error()
		}
	}
	if errStr == "" {
		return nil
	}
	return errors.New(errStr)
}

//导出单个文件
func ExportFile(filePath string, outputPath string, exportTargetList []*ExportTarget) error {
	loadedConfigFileInfoStore = map[string]*LoadedConfigInfo{}

	workSpace = common.ParserFileDirectoryByFullPath(filePath)

	err := doExportFile(filePath, outputPath, exportTargetList)
	if nil != err {
		return errors.New(filePath + " : " + err.Error())
	}
	return nil
}

//初始化
func init() {
	// register
	codeGenToolStore[define.SupportLan_Go] = &genRuntimeCodeTool_Go{}
	codeGenToolStore[define.SupportLan_Csharp] = &genRuntimeCodeTool_Csharp{}
	codeGenToolStore[define.SupportLan_Java] = &genRuntimeCodeTool_Java{}

	commonSerializer := &ConfigSerializer_Common{}
	configSerializerStore[define.SupportLan_Go] = commonSerializer
	configSerializerStore[define.SupportLan_Csharp] = commonSerializer
	configSerializerStore[define.SupportLan_Java] = commonSerializer
	configSerializerStore[define.SupportLan_Json] = &ConfigSerializer_Json{}
}

//文件描述检查
func checkConfigProvisionCorrect(provision *define.ConfigProvisionInfo) error {
	if provision.GlobalInfo.TableType != "list" && provision.GlobalInfo.TableType != "map" {
		return errors.New(getTipMessage(TipMessageDefine_ErrorFieldType, provision.GlobalInfo.TableType))
	}
	if provision.GlobalInfo.TableType == "map" {
		isFindKeyField := false
		for _, field := range provision.LineInfo {
			if field.FieldName == provision.GlobalInfo.TableKeyFieldName {
				isFindKeyField = true
				if field.IsList {
					return errors.New(getTipMessage(TipMessageDefine_ErrorFieldKeyCannotBeList, provision.GlobalInfo.TableKeyFieldName))
				}
				break
			}
		}
		if !isFindKeyField {
			return errors.New(getTipMessage(TipMessageDefine_ErrorFieldKeyNotFound, provision.GlobalInfo.TableKeyFieldName))
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
			return errors.New(getTipMessage(TipMessageDefine_ErrorFieldError, field.FieldName))
		}
		if _, ok := fieldNameMap[field.FieldName]; ok {
			return errors.New(getTipMessage(TipMessageDefine_ErrorFieldRepeated, field.FieldName))
		}
		fieldNameMap[field.FieldName] = 1
	}
	return nil
}

//文件内容检查
func checkConfigContentCorrect(provision *define.ConfigProvisionInfo, content [][]string) error {
	fixedContent := excelHandler.FixExcelFile(content)
	for rowIndex, rowContent := range fixedContent {
		for colIndex, contentCell := range rowContent {
			positionMark := getTipMessage(TipMessageDefine_PositionMark, rowIndex+4, colIndex+1)
			if colIndex >= len(provision.LineInfo) {
				// check error
				return errors.New(getTipMessage(TipMessageDefine_ErrorFormate, positionMark, getTipMessage(TipMessageDefine_OutofLineRange)))
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
						return errors.New(getTipMessage(TipMessageDefine_ErrorFormate, positionMark, err.Error()))
					}
				}
			} else {
				// check type check range
				err := checkFieldTypeCorrect(fieldProvisionInfo.FieldType,
					contentCell,
					fieldProvisionInfo.FieldValueRangeLimitMin,
					fieldProvisionInfo.FieldValueRangeLimitMax)
				if nil != err {
					return errors.New(getTipMessage(TipMessageDefine_ErrorFormate, positionMark, err.Error()))
				}
			}
		}
	}
	return nil
}

//属性检查
func getConfigFileInfo(filePath string, fileName string) ([][]string, *define.ConfigProvisionInfo, error) {
	if v, ok := loadedConfigFileInfoStore[filePath]; ok {
		if v.Provision != nil && v.Content != nil {
			return v.Content, v.Provision, nil
		}
	}

	// load excel file
	content, err := excelHandler.ReadExcelFile(filePath)
	if err != nil {
		return nil, nil, err
	}
	// parser file header
	provision, err := excelHandler.ParserExcelToConfigProvision(content, fileName)
	if err != nil {
		return nil, nil, err
	}
	loadedConfigFileInfoStore[filePath] = &LoadedConfigInfo{Content: content, Provision: provision}
	return content, provision, err
}
func checkFieldTypeCorrect(fieldTypeStr string, content string, minValue string, maxValue string) error {

	fieldType, err := convertStrToFieldType(fieldTypeStr)
	if err != nil {
		if checkFieldIsCorrectInEnum(fieldTypeStr, content) {
			return nil
		}
		return err
	}
	switch fieldType {
	case define.FieldType_typeInt32:
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
	case define.FieldType_typeInt64:
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
	case define.FieldType_typeFloat32:
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
	case define.FieldType_typeFloat64:
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
	case define.FieldType_typeBool:
		var tmpValue bool = false
		err := common.Parser_bool(content, &tmpValue)
		if nil != err {
			return err
		}
		return nil
	case define.FieldType_typeString:
		var tmpValue string = ""
		err := common.Parser_string(content, &tmpValue)
		if nil != err {
			return err
		}
		return nil
	case define.FieldType_typeDateTime:
		var tmpValue time.Time
		err := common.Parser_dateTime(content, &tmpValue)
		if nil != err {
			return err
		}
		err = common.CheckValueLimit_dateTime(tmpValue, minValue, maxValue)
		if nil != err {
			return err
		}
		return nil
	case define.FieldType_typeColor:
		var tmpValue color.RGBA
		err := common.Parser_color(content, &tmpValue)
		if nil != err {
			return err
		}
		return nil
	}
	return errors.New(getTipMessage(TipMessageDefine_UnknownFieldType, fieldTypeStr))
}
func convertStrToFieldType(fileType string) (define.FieldType, error) {
	switch fileType {
	case "int32":
		return define.FieldType_typeInt32, nil
	case "int64":
		return define.FieldType_typeInt64, nil
	case "float32":
		return define.FieldType_typeFloat32, nil
	case "float64":
		return define.FieldType_typeFloat64, nil
	case "bool":
		return define.FieldType_typeBool, nil
	case "string":
		return define.FieldType_typeString, nil
	case "class":
		return define.FieldType_typeClass, nil
	case "time":
		return define.FieldType_typeDateTime, nil
	case "color":
		return define.FieldType_typeColor, nil
	default:
		return define.FieldType_typeInt32, errors.New(getTipMessage(TipMessageDefine_UnknownFieldType, fileType))
	}
}

//导出设置检查
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

//检查引用
func checkReference(content [][]string, provision *define.ConfigProvisionInfo) error {
	errorStr := ""
	for index, v := range provision.LineInfo {
		if v.ReferenceTableName != "" {
			// check this reference
			fileName, fieldName, e := parserReferenceInfo(v.ReferenceTableName)
			if e != nil {
				errorStr += e.Error()
				continue
			}
			for colIndex, lineContent := range content {
				if index < 0 || index >= len(lineContent) {
					errorStr += getTipMessage(TipMessageDefine_ReferenceCheckError, index, colIndex, getTipMessage(TipMessageDefine_ReferenceIndexerror))

					continue
				}
				targetCell := lineContent[index]
				err := checkIsTargetColHavValue(fileName, fieldName, targetCell)
				if nil != err {
					errorStr += getTipMessage(TipMessageDefine_ReferenceCheckError, index, colIndex, err.Error())
					continue
				}
			}
		}
	}
	return nil
}
func checkIsTargetColHavValue(fileName string, fieldName string, targetCell string) error {
	//
	targetCell = strings.TrimSpace(targetCell)
	filePath := workSpace + "/" + fileName + ".xlsx"
	content, provision, err := getConfigFileInfo(filePath, fileName)
	if err != nil {
		return err
	}
	targetColIndex := -1
	for colIndex, fieldInfo := range provision.LineInfo {
		if fieldInfo.FieldName == fieldName {
			targetColIndex = colIndex
			break
		}
	}
	if targetColIndex == -1 {
		return errors.New(getTipMessage(TipMessageDefine_ReferenceFieldNotFound, fieldName, fileName))
	}

	for _, lineContent := range content {
		if targetColIndex < 0 || targetColIndex >= len(lineContent) {
			return errors.New(getTipMessage(TipMessageDefine_ReferenceFieldNotFound, fieldName, fileName))
		}
		if strings.TrimSpace(lineContent[targetColIndex]) == targetCell {
			return nil
		}
	}
	return errors.New(getTipMessage(TipMessageDefine_ReferenceFieldNotFound, fieldName, fileName))
}
func parserReferenceInfo(info string) (configName string, fieldName string, err error) {
	list := strings.Split(info, ":")
	if len(list) != 2 {
		return "", "", errors.New(getTipMessage(TipMessageDefine_ErrorOnParserReference, info))
	}
	return list[0], list[1], nil
}

//枚举检查
func checkFieldIsInEnum(fieldType string) bool {
	for _, enum := range currentConfigEnumInfoList {
		if enum.Name == fieldType {
			return true
		}
	}
	return false
}
func checkFieldIsInEnumWithName(fieldType string) (bool, *define.ConfigEnumInfo) {
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
func parserEnumList(enumList []string) ([]*define.ConfigEnumInfo, error) {

	var keyNameMap = map[string]int{}
	var result []*define.ConfigEnumInfo
	for _, enum := range enumList {
		tmpElem, err := parserEnumString(enum)
		if err != nil {
			return nil, err
		}
		if _, ok := keyNameMap[tmpElem.Name]; ok {
			return nil, errors.New(getTipMessage(TipMessageDefine_ErrorOnParserEnumRepeated, enum, tmpElem.Name))
		}
		keyNameMap[tmpElem.Name] = 1
		result = append(result, tmpElem)
	}
	return result, nil
}
func parserEnumString(enum string) (*define.ConfigEnumInfo, error) {
	// parser enum name
	tmpStr := strings.Split(enum, ":")
	if len(tmpStr) != 2 {
		return nil, errors.New(getTipMessage(TipMessageDefine_ErrorOnParserEnumFormateError, enum))
	}
	result := &define.ConfigEnumInfo{}
	result.Name = tmpStr[0]
	tmpStr = strings.Split(tmpStr[1], "|")
	if len(tmpStr) <= 0 {
		return nil, errors.New(getTipMessage(TipMessageDefine_ErrorOnParserEnumFormateError, enum))
	}
	var keyNameMap = map[string]int{}
	for _, tmpElem := range tmpStr {
		tmpCell := strings.Split(tmpElem, "=")
		if len(tmpCell) != 2 {
			return nil, errors.New(getTipMessage(TipMessageDefine_ErrorOnParserEnumFormateError, enum))
		}
		// check value type
		tmpCellEnum := &define.ConfigEnumElementInfo{}
		tmpCellEnum.Name = tmpCell[0]
		if tmpCellEnum.Name == "" || tmpCellEnum.Name == " " {
			return nil, errors.New(getTipMessage(TipMessageDefine_ErrorOnParserEnumEmpty, enum, tmpElem))
		}
		if _, ok := keyNameMap[tmpCellEnum.Name]; ok {
			return nil, errors.New(getTipMessage(TipMessageDefine_ErrorOnParserEnumRepeated, enum, tmpCellEnum.Name))
		}

		keyNameMap[tmpCellEnum.Name] = 1

		err := common.Parser_int32(tmpCell[1], &tmpCellEnum.Value)
		if nil != err {
			return nil, errors.New(getTipMessage(TipMessageDefine_ErrorOnParserEnum, enum, tmpElem, tmpCell[1]))
		}
		result.Value = append(result.Value, tmpCellEnum)
	}
	return result, nil
}
func convertPbEnum(configName string, pbEnumInfo []*define.ConfigEnumInfo) []*EnumDefine {
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

//执行导出
func doExportFile(filePath string, outputPath string, exportTargetList []*ExportTarget) error {
	currentConfigEnumInfoList = nil

	// parser file name
	fileName, _ := common.ParserFileNameByPath(filePath)

	// load excel file & parser file header
	content, provision, err := getConfigFileInfo(filePath, fileName)

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
		realOutputPath := outputPath + "/" + exportTarget.Name

		// ensure output path
		common.EnsureFolder(realOutputPath)

		// check is need export at this target

		if !checkExportTarget(exportTarget, provision.GlobalInfo.ExportTarget) {
			//
			continue
		}

		tmpConfig, err := doExport(realOutputPath, provision, content, exportTarget)
		if err != nil {
			return errors.New(getTipMessage(TipMessageDefine_ExportWithError, exportTarget.Name, err.Error()))
		}

		byteContent, err := configSerializerStore[exportTarget.Lan].Export(tmpConfig)
		if err != nil {
			return errors.New(getTipMessage(TipMessageDefine_ExportWithError, exportTarget.Name, err.Error()))
		}

		// do export
		common.WriteFileByName(realOutputPath+"/"+provision.TableName+exportTarget.OutPutSuffix, byteContent)
	}

	// begin check reference
	err = checkReference(content, provision)

	return nil
}
func doExport(outputPath string, provision *define.ConfigProvisionInfo, content [][]string, exportTarget *ExportTarget) (*define.ConfigTable, error) {
	pbConfig := &define.ConfigTable{}

	pbConfig.ConfigName = provision.TableName
	pbConfig.PackageName = "config"

	if provision.GlobalInfo.TableType == "list" {
		pbConfig.Type = define.ConfigType_typeList

	} else if provision.GlobalInfo.TableType == "map" {
		pbConfig.Type = define.ConfigType_typeMap
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
				fieldType = define.FieldType_typeEnum
				parameter = enumInfo.Name
			}
		}
		pbConfig.FieldInfoList = append(pbConfig.FieldInfoList, &define.ConfigFieldInfo{Name: field.FieldName, Type: fieldType, IsList: field.IsList, Parameter: parameter})
	}
	fixedContent := excelHandler.FixExcelFile(content)
	for _, rowContent := range fixedContent {
		configLine := &define.ConfigLine{}
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

	if exportTarget.PackageName == "" {
		exportTarget.PackageName = "config"
	}
	define := &ConfigDefine{PackageName: exportTarget.PackageName, ConfigName: provision.TableName}
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
	if genCodeTool, ok := codeGenToolStore[exportTarget.Lan]; ok {
		err := genCodeTool.GenRuntimeCode(outputPath, define, enumDefine)
		if nil != err {
			return nil, err
		}
	}
	return pbConfig, nil
}

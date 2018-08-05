package exportTool

import (
	"errors"
	"github.com/Blizzardx/ConfigProtocol/common"
	"github.com/Blizzardx/ConfigProtocol/define"
	"github.com/Blizzardx/ConfigProtocol/pbConfig"
	"github.com/Blizzardx/ConfigProtocol/tool/excelHandler"
	"github.com/gogo/protobuf/proto"
	"strconv"
	"strings"
)

type ExportTarget struct {
	Name string
	Lan  define.SupportLan
}

func exportFile(filePath string, outputPath string, exportTargetList []*ExportTarget, outputFileSuffix string) error {
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
	if nil != checkConfigProvisionCorrect(provision) {
		return err
	}
	// begin check content
	if nil != checkConfigContentCorrect(provision, content) {
		return err
	}

	// begin export
	for _, exportTarget := range exportTargetList {
		tmpConfig := doExport(provision, content, exportTarget)
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
	var fieldNameMap = map[string]int32{}
	for _, field := range provision.LineInfo {
		if field.FieldType == "" && field.FieldName == "" {
			continue
		}

		_, error := convertStrToFieldType(field.FieldType)
		if error != nil {
			return error
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
func checkConfigContentCorrect(provision *define.ConfigInfo, content [][]string) error {
	fixedContent := excelHandler.FixExcelFile(content)
	for rowIndex, rowContent := range fixedContent {
		for colIndex, contentCell := range rowContent {
			positionMark := " at col: " + strconv.Itoa(colIndex) + " row: " + strconv.Itoa(rowIndex+4)
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
						return errors.New("error on " + positionMark + " " + err.Error())
					}
				}
			} else {
				// check type check range
				err := checkFieldTypeCorrect(fieldProvisionInfo.FieldType,
					contentCell,
					fieldProvisionInfo.FieldValueRangeLimitMin,
					fieldProvisionInfo.FieldValueRangeLimitMax)
				if nil != err {
					return errors.New("error on " + positionMark + " " + err.Error())
				}
			}
		}
	}
	return nil
}
func checkFieldTypeCorrect(fieldTypeStr string, content string, minValue string, maxValue string) error {

	fieldType, _ := convertStrToFieldType(fieldTypeStr)
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
	case config.FieldType_typeBool:
		var tmpValue bool = false
		err := common.Parser_bool(content, &tmpValue)
		if nil != err {
			return err
		}
	case config.FieldType_typeString:
		var tmpValue string = ""
		err := common.Parser_string(content, &tmpValue)
		if nil != err {
			return err
		}
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
	default:
		return config.FieldType_typeInt32, errors.New("unknown field type " + fileType)
	}
}
func doExport(provision *define.ConfigInfo, content [][]string, exportTarget *ExportTarget) *config.ConfigTable {
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
		fieldType, _ := convertStrToFieldType(field.FieldType)

		pbConfig.FieldInfoList = append(pbConfig.FieldInfoList, &config.ConfigFieldInfo{Name: field.FieldName, Type: fieldType, IsList: field.IsList})
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
	return pbConfig
}

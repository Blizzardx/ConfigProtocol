package goRuntime

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Blizzardx/ConfigProtocol/common"
	"github.com/Blizzardx/ConfigProtocol/define"
	"image/color"
	"reflect"
	"strings"
	"time"
)

var workspace = ""

func SetWorkspace(workspacePath string) {
	workspace = workspacePath
}
func LoadConfig(configStruct interface{}) error {
	//parser class name to reflect config name
	configType := reflect.TypeOf(configStruct)

	if configType.Kind() != reflect.Ptr {
		return errors.New("error type")
	}

	configType = configType.Elem()
	if configType.Kind() != reflect.Struct {
		return errors.New("error type")
	}

	structName := configType.Name()
	// load config by name
	byteContent, err := common.LoadFileByName(workspace + "/" + structName + ".json")
	if err != nil {
		return err
	}
	pbConfig := &define.ConfigTable{}
	err = json.Unmarshal(byteContent, pbConfig)
	if err != nil {
		return err
	}
	fieldNum := configType.NumField()
	tableFieldInfo := configType.Field(0)
	if fieldNum != 1 {
		return errors.New("error type")
	}
	configInstance := reflect.ValueOf(configStruct)

	// check type
	if pbConfig.Type == define.ConfigType_typeList {
		if tableFieldInfo.Type.Kind() != reflect.Slice {
			return errors.New("error type")
		}
		parserList(configInstance, configType, pbConfig)
	} else {
		if tableFieldInfo.Type.Kind() != reflect.Map {
			return errors.New("error type")
		}
		parserMap(configInstance, configType, pbConfig)
	}

	return nil
}
func LoadConfigByContent(byteContent []byte, configStruct interface{}) error {
	//parser class name to reflect config name
	configType := reflect.TypeOf(configStruct)

	if configType.Kind() != reflect.Ptr {
		return errors.New("error type")
	}

	configType = configType.Elem()
	if configType.Kind() != reflect.Struct {
		return errors.New("error type")
	}

	pbConfig := &define.ConfigTable{}
	err := json.Unmarshal(byteContent, pbConfig)
	if err != nil {
		return err
	}
	fieldNum := configType.NumField()
	tableFieldInfo := configType.Field(0)
	if fieldNum != 1 {
		return errors.New("error type")
	}
	configInstance := reflect.ValueOf(configStruct)

	// check type
	if pbConfig.Type == define.ConfigType_typeList {
		if tableFieldInfo.Type.Kind() != reflect.Slice {
			return errors.New("error type")
		}
		parserList(configInstance, configType, pbConfig)
	} else {
		if tableFieldInfo.Type.Kind() != reflect.Map {
			return errors.New("error type")
		}
		parserMap(configInstance, configType, pbConfig)
	}

	return nil
}
func parserList(configInstance reflect.Value, tableType reflect.Type, pbConfig *define.ConfigTable) reflect.Value {
	configContentInstance := reflect.New(tableType.Field(0).Type).Elem()

	lineDefineField := tableType.Field(0).Type.Elem().Elem()
	for _, rowContent := range pbConfig.Content {
		lineContentInstance := reflect.New(lineDefineField)
		for colIndex, cell := range rowContent.Content {
			parserLine(colIndex, pbConfig, lineDefineField, cell, lineContentInstance)
		}
		configContentInstance = reflect.Append(configContentInstance, lineContentInstance)
	}
	configInstance.Elem().Field(0).Set(configContentInstance)
	return configInstance
}
func parserMap(configInstance reflect.Value, tableType reflect.Type, pbConfig *define.ConfigTable) reflect.Value {
	configContentInstance := reflect.MakeMap(tableType.Field(0).Type)

	lineDefineField := tableType.Field(0).Type.Elem().Elem()
	for _, rowContent := range pbConfig.Content {
		lineContentInstance := reflect.New(lineDefineField)
		var thisLineKeyValue reflect.Value
		for colIndex, cell := range rowContent.Content {
			isKey, keyValue := parserLine(colIndex, pbConfig, lineDefineField, cell, lineContentInstance)
			if isKey {
				thisLineKeyValue = keyValue
			}
		}
		tmpKey := configContentInstance.MapIndex(thisLineKeyValue)
		if tmpKey.IsValid() {
			// error
			continue
		}
		configContentInstance.SetMapIndex(thisLineKeyValue, lineContentInstance)
	}
	configInstance.Elem().Field(0).Set(configContentInstance)
	return configInstance
}
func parserLine(colIndex int, pbConfig *define.ConfigTable, lineDefineField reflect.Type, cell string, lineContentInstance reflect.Value) (isKey bool, keyValue reflect.Value) {
	isKey = false

	var definedFiledInfo reflect.StructField
	isExist := false
	// read pb field info
	fieldInfo := pbConfig.FieldInfoList[colIndex]
	for i := 0; i < lineDefineField.NumField(); i++ {
		if lineDefineField.Field(i).Name == fieldInfo.Name {
			isExist = true
			definedFiledInfo = lineDefineField.Field(i)
			break
		}
	}
	if !isExist {
		return
	}
	if fieldInfo.IsList {
		if definedFiledInfo.Type.Kind() != reflect.Slice {
			return
		}
		listInstance := reflect.New(definedFiledInfo.Type).Elem()
		cellList := strings.Split(cell, "|")
		for _, tmpCell := range cellList {

			cellValue, err := parserCell(tmpCell, fieldInfo.Type)
			if err != nil {
				continue
			}
			if fieldInfo.Type == define.FieldType_typeEnum {
				tmpCellInstance := reflect.New(listInstance.Type().Elem()).Elem()
				tmpCellInstance.SetInt(int64(cellValue.(int32)))
				listInstance = reflect.Append(listInstance, tmpCellInstance)

			} else {
				listInstance = reflect.Append(listInstance, reflect.ValueOf(cellValue))
			}
		}
		lineContentInstance.Elem().FieldByName(definedFiledInfo.Name).Set(listInstance)
	} else {
		// check type
		if !checkType(definedFiledInfo.Type.Kind(), fieldInfo.Type) {
			return
		}
		cellValue, err := parserCell(cell, fieldInfo.Type)
		if err != nil {
			//return err
			fmt.Println(err)
			return
		}
		if fieldInfo.Type == define.FieldType_typeEnum {
			tmpCellInstance := reflect.New(lineContentInstance.Elem().FieldByName(definedFiledInfo.Name).Type()).Elem()
			tmpCellInstance.SetInt(int64(cellValue.(int32)))
			lineContentInstance.Elem().FieldByName(definedFiledInfo.Name).Set(tmpCellInstance)

		} else {
			lineContentInstance.Elem().FieldByName(definedFiledInfo.Name).Set(reflect.ValueOf(cellValue))
		}
	}

	if fieldInfo.Name == pbConfig.KeyFieldName {
		keyValue = lineContentInstance.Elem().FieldByName(definedFiledInfo.Name)
		isKey = true
	}
	return
}
func parserCell(cell string, pbType define.FieldType) (interface{}, error) {
	switch pbType {

	case define.FieldType_typeInt32:
		var tmpValue int32
		err := common.Parser_int32(cell, &tmpValue)
		return tmpValue, err
	case define.FieldType_typeInt64:
		var tmpValue int64
		err := common.Parser_int64(cell, &tmpValue)
		return tmpValue, err
	case define.FieldType_typeFloat32:
		var tmpValue float32
		err := common.Parser_float32(cell, &tmpValue)
		return tmpValue, err
	case define.FieldType_typeFloat64:
		var tmpValue float64
		err := common.Parser_float64(cell, &tmpValue)
		return tmpValue, err
	case define.FieldType_typeBool:
		var tmpValue bool
		err := common.Parser_bool(cell, &tmpValue)
		return tmpValue, err
	case define.FieldType_typeString:
		var tmpValue string
		err := common.Parser_string(cell, &tmpValue)
		return tmpValue, err
	case define.FieldType_typeEnum:
		var tmpValue int32
		err := common.Parser_int32(cell, &tmpValue)
		return tmpValue, err
	case define.FieldType_typeDateTime:
		var tmpValue time.Time
		err := common.Parser_dateTime(cell, &tmpValue)
		return tmpValue, err
	case define.FieldType_typeColor:
		var tmpValue color.RGBA
		err := common.Parser_color(cell, &tmpValue)
		return tmpValue, err
		//case config.FieldType_typeClass:
	}
	return 0, errors.New("unsupport type ")
}
func checkType(definedType reflect.Kind, pbType define.FieldType) bool {
	switch pbType {

	case define.FieldType_typeInt32:
		return definedType == reflect.Int32
	case define.FieldType_typeInt64:
		return definedType == reflect.Int64
	case define.FieldType_typeFloat32:
		return definedType == reflect.Float32
	case define.FieldType_typeFloat64:
		return definedType == reflect.Float64
	case define.FieldType_typeBool:
		return definedType == reflect.Bool
	case define.FieldType_typeString:
		return definedType == reflect.String
	case define.FieldType_typeClass:
		return definedType == reflect.Struct
	case define.FieldType_typeEnum:
		return definedType == reflect.Int32
	case define.FieldType_typeDateTime:
		return definedType == reflect.Struct
	case define.FieldType_typeColor:
		return definedType == reflect.Struct
	}
	return false
}

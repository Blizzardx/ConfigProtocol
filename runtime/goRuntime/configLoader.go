package goRuntime

import (
	"errors"
	"github.com/Blizzardx/ConfigProtocol/common"
	"github.com/Blizzardx/ConfigProtocol/pbConfig"
	"github.com/gogo/protobuf/proto"
	"reflect"
	"strings"
)

var workspace = ""

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
	byteContent, err := common.LoadFileByName(workspace + "/" + structName + ".bytes")
	if err != nil {
		return err
	}
	pbConfig := &config.ConfigTable{}
	err = proto.Unmarshal(byteContent, pbConfig)
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
	if pbConfig.Type == config.ConfigType_typeList {
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
func parserList(configInstance reflect.Value, tableType reflect.Type, pbConfig *config.ConfigTable) reflect.Value {
	configContentInstance := reflect.New(tableType.Field(0).Type)
	configInstance.Field(0).Set(configContentInstance)

	lineDefineField := tableType.Field(0).Type.Elem().Elem()
	for _, rowContent := range pbConfig.Content {
		lineContentInstance := reflect.New(lineDefineField)
		for colIndex, cell := range rowContent.Content {
			parserLine(colIndex, pbConfig, lineDefineField, cell, lineContentInstance)
		}
		configContentInstance = reflect.Append(configContentInstance, lineContentInstance)
	}
	return configInstance
}
func parserMap(configInstance reflect.Value, tableType reflect.Type, pbConfig *config.ConfigTable) reflect.Value {
	configContentInstance := reflect.New(tableType.Field(0).Type)
	configInstance.Field(0).Set(configContentInstance)

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
		if !tmpKey.IsValid() {
			// error
			continue
		}
		configContentInstance.SetMapIndex(thisLineKeyValue, lineContentInstance)
	}
	return configInstance
}
func parserLine(colIndex int, pbConfig *config.ConfigTable, lineDefineField reflect.Type, cell string, lineContentInstance reflect.Value) (isKey bool, keyValue reflect.Value) {
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
		listInstance := reflect.New(definedFiledInfo.Type)
		cellList := strings.Split(cell, "|")
		for _, tmpCell := range cellList {

			cellValue, err := parserCell(tmpCell, fieldInfo.Type)
			if err != nil {
				continue
			}
			listInstance = reflect.Append(listInstance, reflect.ValueOf(cellValue))
		}
		lineContentInstance.FieldByName(definedFiledInfo.Name).Set(listInstance)
	} else {
		// check type
		if !checkType(definedFiledInfo.Type.Kind(), fieldInfo.Type) {
			return
		}
		cellValue, err := parserCell(cell, fieldInfo.Type)
		if err != nil {
			//return err
			return
		}
		lineContentInstance.FieldByName(definedFiledInfo.Name).Set(reflect.ValueOf(cellValue))
	}

	if fieldInfo.Name == pbConfig.KeyFieldName {
		keyValue = lineContentInstance.FieldByName(definedFiledInfo.Name)
		isKey = true
	}
	return
}
func parserCell(cell string, pbType config.FieldType) (interface{}, error) {
	switch pbType {

	case config.FieldType_typeInt32:
		var tmpValue int32
		err := common.Parser_int32(cell, &tmpValue)
		return tmpValue, err
	case config.FieldType_typeInt64:
		var tmpValue int64
		err := common.Parser_int64(cell, &tmpValue)
		return tmpValue, err
	case config.FieldType_typeFloat32:
		var tmpValue float32
		err := common.Parser_float32(cell, &tmpValue)
		return tmpValue, err
	case config.FieldType_typeFloat64:
		var tmpValue float64
		err := common.Parser_float64(cell, &tmpValue)
		return tmpValue, err
	case config.FieldType_typeBool:
		var tmpValue bool
		err := common.Parser_bool(cell, &tmpValue)
		return tmpValue, err
	case config.FieldType_typeString:
		var tmpValue string
		err := common.Parser_string(cell, &tmpValue)
		return tmpValue, err
		//case config.FieldType_typeClass:
	}
	return 0, errors.New("unsupport type ")
}
func checkType(definedType reflect.Kind, pbType config.FieldType) bool {
	switch pbType {

	case config.FieldType_typeInt32:
		return definedType == reflect.Int32
	case config.FieldType_typeInt64:
		return definedType == reflect.Int64
	case config.FieldType_typeFloat32:
		return definedType == reflect.Float32
	case config.FieldType_typeFloat64:
		return definedType == reflect.Float64
	case config.FieldType_typeBool:
		return definedType == reflect.Bool
	case config.FieldType_typeString:
		return definedType == reflect.String
	case config.FieldType_typeClass:
		return definedType == reflect.Struct
	}
	return false
}

package define

// 类型
type FieldType int32

const (
	FieldType_Int32 FieldType = iota
	FieldType_Int64
	FieldType_Float32
	FieldType_Float64
	FieldType_String
	FieldType_Bool
	FieldType_Custom
)

type ConfigInfo struct {
	TableName  string
	GlobalInfo *ConfigHeadInfo
	LineInfo   []*ConfigFieldInfo
}
type ConfigHeadInfo struct {
	TableType         string                      `json:"type"` // list or map
	TableKeyFieldName string                      `json:"keyName"`
	GlobalDefine      []*ConfigConstFieldInfo     `json:"const"`
	GlobalEnumDefine  []*ConfigConstEnumFieldInfo `json:"constEnum"`
}
type ConfigConstFieldInfo struct {
	FieldType  string `json:"type"`
	FieldName  string `json:"name"`
	FieldValue string `json:"value"`
}
type ConfigConstEnumFieldInfo struct {
	FieldName  string  `json:"name"`
	FieldValue []int32 `json:"value"`
}
type ConfigEnumKeyValueInfo struct {
	FieldName  string `json:"name"`
	FieldValue int32  `json:"value"`
}
type ConfigFieldInfo struct {
	FieldType               string `json:"type"`
	FieldName               string `json:"name"`
	FieldValueRangeLimitMin string `json:"min"`
	FieldValueRangeLimitMax string `json:"max"`
	IsList                  bool   `json:"isList"`   // if is list ,split by '|'
	ReferenceTableName      string `json:"refTable"` // like itemConfig:itemId ,split by ':'
}

func ConvertStrToFieldType(fileType string) FieldType {
	switch fileType {
	case "int32":
		return FieldType_Int32
	case "int64":
		return FieldType_Int64
	case "float32":
		return FieldType_Float32
	case "float64":
		return FieldType_Float64
	case "bool":
		return FieldType_Bool
	case "string":
		return FieldType_String
	default:
		return FieldType_Custom
	}
}

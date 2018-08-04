package define

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
type ConfigFieldInfo struct {
	FieldType               string `json:"type"`
	FieldName               string `json:"name"`
	FieldValueRangeLimitMin string `json:"min"`
	FieldValueRangeLimitMax string `json:"max"`
	IsList                  bool   `json:"isList"`   // if is list ,split by '|'
	ReferenceTableName      string `json:"refTable"` // like itemConfig:itemId ,split by ':'
	ExportTarget            string `json:"exTarget"` // split by '|' like 'client|server|serverTool', 'client' ,'server' ,''(empty) mean's all target
}

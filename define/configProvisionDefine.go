package define

type ConfigInfo struct {
	TableName  string
	GlobalInfo *ConfigHeadInfo
	LineInfo   []*ConfigFieldInfo
}
type ConfigHeadInfo struct {
	TableType         string   `json:"type"` // list or map
	TableKeyFieldName string   `json:"keyName"`
	GlobalEnumDefine  []string `json:"enum"` //like quality:black=1|white=2|yellow=3
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

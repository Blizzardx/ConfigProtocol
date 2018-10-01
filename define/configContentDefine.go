package define

type ConfigType int32

const (
	ConfigType_typeList ConfigType = 0
	ConfigType_typeMap  ConfigType = 1
)

type FieldType int32

const (
	FieldType_typeInt32    FieldType = 0
	FieldType_typeInt64    FieldType = 1
	FieldType_typeFloat32  FieldType = 2
	FieldType_typeFloat64  FieldType = 3
	FieldType_typeBool     FieldType = 4
	FieldType_typeString   FieldType = 5
	FieldType_typeClass    FieldType = 6
	FieldType_typeEnum     FieldType = 7
	FieldType_typeDateTime FieldType = 8
	FieldType_typeColor    FieldType = 9
)

type ConfigFieldInfo struct {
	Name      string    `json:"name,omitempty"`
	Type      FieldType `json:"type,omitempty"`
	IsList    bool      `json:"isList,omitempty"`
	Parameter string    `json:"parameter,omitempty"`
}
type ConfigEnumElementInfo struct {
	Name  string `json:"name,omitempty"`
	Value int32  `json:"value,omitempty"`
}
type ConfigEnumInfo struct {
	Name  string                   `json:"name,omitempty"`
	Value []*ConfigEnumElementInfo `json:"value,omitempty"`
}

type ConfigLine struct {
	Content []string `json:"content,omitempty"`
}
type ConfigTable struct {
	PackageName   string             `json:"packageName,omitempty"`
	ConfigName    string             `json:"configName,omitempty"`
	Type          ConfigType         `json:"type,omitempty"`
	KeyFieldName  string             `json:"keyFieldName,omitempty"`
	FieldInfoList []*ConfigFieldInfo `json:"fieldInfoList,omitempty"`
	EnumInfoList  []*ConfigEnumInfo  `json:"enumInfoList,omitempty"`
	Content       []*ConfigLine      `json:"content,omitempty"`
}

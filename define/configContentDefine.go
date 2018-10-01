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
	Name      string    `json:"Name,omitempty"`
	Type      FieldType `json:"Type,omitempty"`
	IsList    bool      `json:"IsList,omitempty"`
	Parameter string    `json:"Parameter,omitempty"`
}
type ConfigEnumElementInfo struct {
	Name  string `json:"Name,omitempty"`
	Value int32  `json:"Value,omitempty"`
}
type ConfigEnumInfo struct {
	Name  string                   `json:"Name,omitempty"`
	Value []*ConfigEnumElementInfo `json:"Value,omitempty"`
}

type ConfigLine struct {
	Content []string `json:"Content,omitempty"`
}
type ConfigTable struct {
	PackageName   string             `json:"PackageName,omitempty"`
	ConfigName    string             `json:"ConfigName,omitempty"`
	Type          ConfigType         `json:"Type,omitempty"`
	KeyFieldName  string             `json:"KeyFieldName,omitempty"`
	FieldInfoList []*ConfigFieldInfo `json:"FieldInfoList,omitempty"`
	EnumInfoList  []*ConfigEnumInfo  `json:"EnumInfoList,omitempty"`
	Content       []*ConfigLine      `json:"Content,omitempty"`
}

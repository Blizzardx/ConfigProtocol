using System.Collections.Generic;

namespace define
{
    public enum ConfigType
    {
        TypeList = 0,
        TypeMap = 1,
    }
    public enum FieldType
    {
        TypeInt32     = 0,
        TypeInt64     = 1,
        TypeFloat32   = 2,
        TypeFloat64   = 3,
        TypeBool      = 4,
        TypeString    = 5,
        TypeClass     = 6,
        TypeEnum      = 7,
        TypeDateTime  = 8,
        TypeColor     = 9,
    }

    public class ConfigFieldInfo
    {
        public string Name;
        public FieldType Type;
        public bool IsList;
        public string Parameter;
    }

    public class ConfigEnumElementInfo
    {
        public string Name;
        public int Value;
    }

    public class ConfigEnumInfo
    {
        public string Name;
        public List<ConfigEnumElementInfo> Value;
    }

    public class ConfigLine
    {
        public List<string> Content;
    }

    public class ConfigTable
    {
        public string PackageName;
        public string ConfigName;
        public ConfigType Type;
        public string KeyFieldName;
        public List<ConfigFieldInfo> FieldInfoList;
        public List<ConfigEnumInfo> EnumInfoList;
        public List<ConfigLine> Content;
    }
}
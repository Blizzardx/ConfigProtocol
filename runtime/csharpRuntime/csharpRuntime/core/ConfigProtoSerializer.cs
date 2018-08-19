using System;
using System.Reflection;
using Config;
using Google.Protobuf;

namespace ConfigProto.core
{
    public class ConfigProtoSerializer
    {
        public static Object DeSerialize(byte[] configContent, out string error)
        {
            error = "";
            Config.ConfigTable config = new ConfigTable();

            try
            {
                config.MergeFrom(configContent);
            }
            catch (Exception e)
            {
                error = "error on parser config content " + e.Message;
                Logger.Instance.LogError(error);
                return null;
            }

            Type configTyep = ReflectionManager.Instance.GetTypeByName(config.PackageName);
            if (null == configTyep)
            {
                error = "error on get type by name " + config.PackageName;
                Logger.Instance.LogError(error);
                return null;
            }
            var configInstance = Activator.CreateInstance(configTyep);
            PropertyInfo[] pps = configInstance.GetType().GetProperties(BindingFlags.Public | BindingFlags.Instance);
            if (pps.Length != 1)
            {
                error = "error field count " + config.PackageName;
                Logger.Instance.LogError(error);
                return null;
            }
            if (config.Type == ConfigType.TypeList)
            {
                ParserList(configInstance,config, pps[0],out error);
            }
            else
            {
                ParserMap(configInstance, config, pps[0], out error);
            }
            if (error != "")
            {
                Logger.Instance.LogError(error);
            }
            return configInstance;
        }

        #region parser list
        private static void ParserList(Object configInstance,ConfigTable config,PropertyInfo propertyInfo, out string error)
        {
            error = "";
            string fieldName = config.ConfigName + "Info";
            if (propertyInfo.PropertyType.Name != "List<" + fieldName + ">")
            {
                error = "error field type ,type must be " + "List<" + fieldName + ">" + " now is " + propertyInfo.PropertyType.Name;
                return ;
            }
            var lineType = ReflectionManager.Instance.GetTypeByName(fieldName);
            if (null == lineType)
            {
                error = "cant' find type by name " + fieldName;
                return;
            }
            var content = Activator.CreateInstance(propertyInfo.PropertyType);
            // set content line 
            propertyInfo.SetValue(configInstance, content);
            var lineMethorInfo = content.GetType().GetMethod("Add");
            if (null == lineMethorInfo)
            {
                error = "cant' find method by name Add";
                return;
            }

            string parserCellErrorMsg = "";
            for (int j = 0; j < config.Content.Count; ++j)
            {
                var line = config.Content[j];

                if (line.Content.Count != config.FieldInfoList.Count)
                {
                    Logger.Instance.LogWarning("line content count != fieldInfo count");
                    continue;
                }
                
                var lineInstance = Activator.CreateInstance(lineType);
                for (int i = 0; i < config.FieldInfoList.Count; ++i)
                {
                    string posInfo = "error at pos line: " + (j + 4) + " col: " + (i + 1)+ " ";

                    string parserErrorMsg = "";
                    var fieldInfo = config.FieldInfoList[i];
                    var contentCell = line.Content[i];
                    ParserCell(lineInstance,fieldInfo, contentCell, out parserErrorMsg);
                    if ("" != parserErrorMsg)
                    {
                        parserCellErrorMsg += posInfo + parserErrorMsg;
                    }
                }
                // do list.Add(lineInstance);
                lineMethorInfo.Invoke(content, new []{lineInstance});
            }
            return ;
        }
        #endregion

        #region parser map
        private static void ParserMap(Object configInstance, ConfigTable config, PropertyInfo propertyInfo, out string error)
        {
            error = "";
            
            ConfigFieldInfo keyFieldInfo = null;
            foreach (var fieldInfo in config.FieldInfoList)
            {
                if (fieldInfo.Name == config.KeyFieldName)
                {
                    keyFieldInfo = fieldInfo;
                    break;
                }
            }
            if (null == keyFieldInfo)
            {
                error = "cant' find key fieldinfo by name " + config.KeyFieldName;
                return;
            }
            string keyFieldName = ConvertFieldInfo(keyFieldInfo.Type, out error);
            if ("" != error)
            {
                return;
            }

            string fieldName = config.ConfigName + "Info";
            if (propertyInfo.PropertyType.Name != "Dictionary<" + keyFieldName  + "," + fieldName + ">")
            {
                error = "error field type ,type must be " + "List<" + fieldName + ">" + " now is " + propertyInfo.PropertyType.Name;
                return;
            }
            var lineType = ReflectionManager.Instance.GetTypeByName(fieldName);
            if (null == lineType)
            {
                error = "cant' find type by name " + fieldName;
                return;
            }
            var content = Activator.CreateInstance(propertyInfo.PropertyType);
            // set content line 
            propertyInfo.SetValue(configInstance, content);
            var ContainsKeyMethorInfo = content.GetType().GetMethod("ContainsKey");
            if (null == ContainsKeyMethorInfo)
            {
                error = "cant' find method by name ContainsKey";
                return;
            }
            var lineMethorInfo = content.GetType().GetMethod("Add");
            if (null == lineMethorInfo)
            {
                error = "cant' find method by name Add";
                return;
            }

            string parserCellErrorMsg = "";
            for (int j = 0; j < config.Content.Count; ++j)
            {
                var line = config.Content[j];

                if (line.Content.Count != config.FieldInfoList.Count)
                {
                    Logger.Instance.LogWarning("line content count != fieldInfo count");
                    continue;
                }

                var lineInstance = Activator.CreateInstance(lineType);
                Object keyInstance = null;
                for (int i = 0; i < config.FieldInfoList.Count; ++i)
                {
                    string posInfo = "error at pos line: " + (j + 4) + " col: " + (i + 1) + " ";

                    string parserErrorMsg = "";
                    var fieldInfo = config.FieldInfoList[i];
                    var contentCell = line.Content[i];
                    var cellInstance = ParserCell(lineInstance, fieldInfo, contentCell, out parserErrorMsg);
                    if ("" != parserErrorMsg)
                    {
                        parserCellErrorMsg += posInfo + parserErrorMsg;
                        continue;
                    }
                    if (fieldInfo.Name == config.KeyFieldName)
                    {
                        keyInstance = cellInstance;
                    }
                }
                if (null == keyInstance)
                {
                    parserCellErrorMsg += "error on parser may key at line " + (j + 4);
                    continue;
                }
                bool res = (bool)ContainsKeyMethorInfo.Invoke(content, new object[] { keyInstance });
                if (res)
                {
                    parserCellErrorMsg += "key already in map " + keyInstance + " at line: " + (j + 4);
                    continue;
                }

                // do map.Add(lineInstance);
                lineMethorInfo.Invoke(content, new[] { keyInstance, lineInstance });
            }
            return;
        }
        #endregion

        #region parser tool
        private static Object ParserCell(Object lineInstance, ConfigFieldInfo fieldInfo, string cell, out string errorMsg)
        {
            errorMsg = "";
            PropertyInfo propertyInfo = lineInstance.GetType().GetProperty(fieldInfo.Name);
            if (null == propertyInfo)
            {
                errorMsg = "cant' find property by name " + fieldInfo.Name;
                return null;
            }
            Object cellInstance = null;
            // do parser cell
            if (fieldInfo.IsList)
            {
                var cellList = cell.Split('|');
                cellInstance = Activator.CreateInstance(propertyInfo.PropertyType);
                var lineMethorInfo = cellInstance.GetType().GetMethod("Add");
                if (null == lineMethorInfo)
                {
                    errorMsg = "cant' find method by name Add at field " + fieldInfo.Name;
                    return null;
                }
                for (int i = 0; i < cellList.Length; ++i)
                {
                    var tmpCell = DoParserCell(cellList[i], fieldInfo, out errorMsg);
                    if ("" != errorMsg)
                    {
                        return null;
                    }
                    // do cellInstance.Add(tmpCell);
                    lineMethorInfo.Invoke(cellInstance, new[] { tmpCell });
                }
            }
            else
            {
                cellInstance = DoParserCell(cell, fieldInfo, out errorMsg);

            }
            propertyInfo.SetValue(lineInstance, cellInstance);
            
            return cellInstance;
        }
        private static Object DoParserCell(string cell, ConfigFieldInfo fieldInfo, out string errorMsg)
        {
            errorMsg = "";
            switch (fieldInfo.Type)
            {
                case FieldType.TypeInt32:
                    return ParserTool.Parser_Int32(cell, out errorMsg);
                case FieldType.TypeInt64:
                    return ParserTool.Parser_Int64(cell, out errorMsg);
                case FieldType.TypeFloat32:
                    return ParserTool.Parser_Float32(cell, out errorMsg);
                case FieldType.TypeFloat64:
                    return ParserTool.Parser_Float64(cell, out errorMsg);
                case FieldType.TypeBool:
                    return ParserTool.Parser_Bool(cell, out errorMsg);
                case FieldType.TypeString:
                    return ParserTool.Parser_String(cell, out errorMsg);
                case FieldType.TypeEnum:
                    return ParserTool.Parser_Int32(cell, out errorMsg);
                case FieldType.TypeDateTime:
                    return ParserTool.Parser_DateTime(cell, out errorMsg);
                case FieldType.TypeColor:
                    break;
                case FieldType.TypeClass:
                    break;
            }
            errorMsg = "error on parser " + cell + " unknown type " + fieldInfo.Type;
            return null;
        }
        private static string ConvertFieldInfo(FieldType fieldName,out string errorMsg)
        {
            errorMsg = "";
            switch (fieldName)
            {
                case FieldType.TypeInt32:
                    return "int";
                case FieldType.TypeInt64:
                    return "long";
                case FieldType.TypeFloat32:
                    return "float";
                case FieldType.TypeFloat64:
                    return "double";
                case FieldType.TypeBool:
                    return "bool";
                case FieldType.TypeString:
                    return "string";
                case FieldType.TypeEnum:
                    return "int";
                case FieldType.TypeDateTime:
                    return "DateTime";
                case FieldType.TypeColor:
                    return "Color";
                case FieldType.TypeClass:
                    break;
            }
            errorMsg = "unkonw field type " + fieldName;
            return "";
        }
        #endregion

    }
}

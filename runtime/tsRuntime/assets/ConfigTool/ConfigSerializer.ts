export enum ConfigType
{
    TypeList = 0,
    TypeMap = 1,
}
export enum FieldType
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

export class ConfigFieldInfo
{
    public Name : string;
    public Type : FieldType ;
    public  IsList : boolean;
    public  Parameter : string;
}

export class ConfigEnumElementInfo
{
    public  Name : string;
    public  Value : number;
}

export class ConfigEnumInfo
{
    public Name : string ;
    public Value : Array<ConfigEnumElementInfo> ;
}

export class ConfigLine
{
    public Content : Array<string> ;
}

export class ConfigTable
{
    public  PackageName : string;
    public  ConfigName : string;
    public Type : ConfigType;
    public KeyFieldName : string;
    public FieldInfoList : Array<ConfigFieldInfo>;
    public EnumInfoList : Array<ConfigEnumInfo>;
    public Content : Array<ConfigLine>;
}

export  class ConfigSerializer {
    private constructor(){
    }
    private static g_Instance : ConfigSerializer  = null;
    public static GetInstance(){
        if(null == this.g_Instance){
            this.g_Instance = new ConfigSerializer();
            this.g_Instance.Init();
        }
        return this.g_Instance;
    }

    private Init(){
    }
    public test(){
    }
    public DeSerialize(configContent:string):any{
        try {
            let configTable = JSON.parse(configContent);
            let resultTable : ConfigTable = configTable;
            this.FixUndefinedField(resultTable);
            return resultTable;
        }catch (e) {
            console.error("error on parser config " + e);
        }
        return null;
    }
    private CreateObjectByTable(configTable : ConfigTable):any{
        let configContainer = new Object();

        configContainer["Content"] = new Object();

        return configContainer;
    }
    private FixUndefinedField(configTable : ConfigTable){
        configTable.PackageName = this.FixUndefined(configTable.PackageName,"");
        configTable.ConfigName = this.FixUndefined(configTable.ConfigName,"");
        configTable.Type = this.FixUndefined( configTable.Type,ConfigType.TypeList);
        configTable.KeyFieldName = this.FixUndefined( configTable.KeyFieldName,"");
        configTable.FieldInfoList = this.FixUndefined(configTable.FieldInfoList,new Array<ConfigFieldInfo>());
        configTable.EnumInfoList = this.FixUndefined(configTable.EnumInfoList,new Array<ConfigEnumInfo>());
        configTable.Content = this.FixUndefined(configTable.Content,new Array<ConfigLine>());

        //field info list
        for(var i=0;i<configTable.FieldInfoList.length;i++){
            let finalFieldInfo = configTable.FieldInfoList[i];
            finalFieldInfo.Name = this.FixUndefined(finalFieldInfo.Name,"");
            finalFieldInfo.Type = this.FixUndefined(finalFieldInfo.Type,FieldType.TypeInt32);
            finalFieldInfo.IsList = this.FixUndefined(finalFieldInfo.IsList,false);
            finalFieldInfo.Parameter = this.FixUndefined(finalFieldInfo.Parameter,"");
        }

        // enum info list
        for(var i=0;i<configTable.EnumInfoList.length;i++){
            let finalEnumInfo =configTable.EnumInfoList[i];
            finalEnumInfo.Name = this.FixUndefined(finalEnumInfo.Name,"");
            finalEnumInfo.Value = this.FixUndefined(finalEnumInfo.Value,new Array<ConfigEnumElementInfo>());

            for(var j=0;j<finalEnumInfo.Value.length;j++){
                let finalEnumValueInfo = finalEnumInfo.Value[j];
                finalEnumValueInfo.Name = this.FixUndefined(finalEnumValueInfo.Name,"");
                finalEnumValueInfo.Value = this.FixUndefined(finalEnumValueInfo.Value,0);
            }
        }

        return configTable;
    }
    private DeSerializeToTemplateConfig(configTableObj : any):ConfigTable{
        let configTable = new ConfigTable();
        configTable.PackageName = this.GetValueByName(configTableObj,"PackageName","");
        configTable.ConfigName = this.GetValueByName(configTableObj,"ConfigName","");
        configTable.Type = this.GetValueByName(configTableObj,"Type",ConfigType.TypeList);
        configTable.KeyFieldName = this.GetValueByName(configTableObj,"KeyFieldName","");
        configTable.FieldInfoList = new Array<ConfigFieldInfo>();
        configTable.EnumInfoList = new Array<ConfigEnumInfo>();
        configTable.Content = new Array<ConfigLine>();

        let fieldInfoList = configTableObj["FieldInfoList"];
        let enumInfoList = configTableObj["EnumInfoList"];
        let contentList = configTableObj["Content"];

        //field info list
        if(fieldInfoList){
            for(var fieldKey in fieldInfoList){
                let fieldInfo = fieldInfoList[fieldKey];
                let finalFieldInfo = new ConfigFieldInfo();
                finalFieldInfo.Name = this.GetValueByName(fieldInfo,"Name","");
                finalFieldInfo.Type = this.GetValueByName(fieldInfo,"Type",FieldType.TypeInt32);
                finalFieldInfo.IsList = this.GetValueByName(fieldInfo,"IsList",false);
                finalFieldInfo.Parameter = this.GetValueByName(fieldInfo,"Parameter","");
                configTable.FieldInfoList.push(finalFieldInfo);
            }
        }

        // enum info list
        if(enumInfoList){
            for(var enumKey in enumInfoList){
                let enumInfo = enumInfoList[enumKey];
                let finalEnumInfo = new ConfigEnumInfo();
                finalEnumInfo.Name = this.GetValueByName(enumInfo,"Name","");
                finalEnumInfo.Value = new Array<ConfigEnumElementInfo>();
                let enumValueList = enumInfo["Value"];
                if(enumValueList){
                    for(var enumValueKey in enumValueList){
                        let enumValueInfo = enumValueList[enumValueKey];
                        let finalEnumValueInfo = new ConfigEnumElementInfo();
                        finalEnumValueInfo.Name = this.GetValueByName(enumValueInfo,"Name","");
                        finalEnumValueInfo.Value = this.GetValueByName(enumValueInfo,"Value",0);
                        finalEnumInfo.Value.push(finalEnumValueInfo);
                    }
                }
                configTable.EnumInfoList.push(finalEnumInfo);
            }
        }

        // content
        if(contentList){
            for(var contentKey in contentList){
                let contentInfo = contentList[contentKey];
                let finalContentInfo = new ConfigLine();
                finalContentInfo.Content = new Array<string>();
                let lineList = contentInfo["Content"];
                if(lineList){
                    for(let lineKey in lineList){
                        let cell = lineList[lineKey];
                        finalContentInfo.Content.push(cell);
                    }
                }
                configTable.Content.push(contentInfo);
            }
        }
        return configTable;
    }
    private GetValueByName(content : any,name : string,defaultValue : any):any{
        let value = content[name];
        if(!value){
            return defaultValue;
        }
        return value;
    }
    private FixUndefined(value : any,defaultValue : any):any{
        if(!value){
            return defaultValue;
        }
        return value;
    }
}

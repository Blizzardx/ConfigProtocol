package core.define;


import java.util.ArrayList;
import java.util.List;

public class ConfigTable
{
    public ConfigTable(){
        this.Type = ConfigType.typeList;
    }

    public String PackageName;
    public String getPackageName(){return this.PackageName;}
    public void setPackageName(String value){this.PackageName = value;}

    public String ConfigName;
    public String getConfigName(){return this.ConfigName;}
    public void setConfigName(String value){this.ConfigName = value;}

    public ConfigType Type;
    public ConfigType getType(){return this.Type;}
    public void setType(ConfigType value){this.Type = value;}

    public String KeyFieldName;
    public String getKeyFieldName(){return this.KeyFieldName;}
    public void setKeyFieldName(String value){this.KeyFieldName = value;}

    public List<ConfigFieldInfo> FieldInfoList;
    public ConfigFieldInfo getFieldInfoList(int index){return this.FieldInfoList.get(index);}
    public int getFieldInfoListCount(){return this.FieldInfoList.size();}
    public void setFieldInfoList(ArrayList<ConfigFieldInfo> value){this.FieldInfoList = value;}

    public ArrayList<ConfigEnumInfo> EnumInfoList;
    public  ArrayList<ConfigEnumInfo> getEnumInfoList(){return this.EnumInfoList;}
    public void setEnumInfoList( ArrayList<ConfigEnumInfo> value){this.EnumInfoList = value;}

    public ArrayList<ConfigLine> Content;
    public ConfigLine getContent(int index){return this.Content.get(index);}
    public void setContent(ArrayList<ConfigLine> value){this.Content = value;}
    public int getContentCount(){return this.Content.size();}

}
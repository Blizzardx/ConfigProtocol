package core.define;


import java.util.ArrayList;

public class ConfigEnumInfo
{
    public String Name;
    public String getName(){return this.Name;}
    public void setName(String value){this.Name = value;}

    public ArrayList<ConfigEnumElementInfo> Value;
}
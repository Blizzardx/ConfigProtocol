package core.define;


public class ConfigFieldInfo
{
    public ConfigFieldInfo(){
        this.Type = FieldType.typeInt32;
    }
    public String Name;
    public String getName(){return this.Name;}
    public void setName(String value){this.Name = value;}

    public FieldType Type;
    public FieldType getType(){return this.Type;}
    public void setType(FieldType value){this.Type = value;}

    public boolean IsList;
    public boolean getIsList(){return this.IsList;}
    public void setIsList(boolean value){this.IsList = value;}

    public String Parameter;
    public String getParameter(){return this.Parameter;}
    public void setParameter(String value){this.Parameter = value;}

}
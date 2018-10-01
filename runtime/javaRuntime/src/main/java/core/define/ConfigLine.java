package core.define;


import java.util.ArrayList;

public class ConfigLine
{
    public ArrayList<String> Content;
    public int getContentCount(){return this.Content.size();}
    public String getContent(int index){return this.Content.get(index);}
}
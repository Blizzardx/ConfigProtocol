package config;

import java.awt.*;
import java.util.*;

enum BasicItem_CommonQuality
{
    yelow(1) ,
    blue(2),
    red(3),
    ;
    private int index;
    private BasicItem_CommonQuality(int i){
        index = i;
    }
    public int getIndex(){
        return index;
    }
}


public class BasicItem_Common
{
    public ArrayList<BasicItem_CommonInfo> Content;
}
 class BasicItem_CommonInfo
{

    public Integer Id;

    public String Icon;

    public Integer Quality;

    public Integer Price;

    public ArrayList<Double> LimitNum;

    public Long Acauire;

    public String ConsumeItem;

    public Integer ConsumeCoin;

    public Integer FormatIndex;

    public BasicItem_CommonQuality Quality1;

    public Date ShowTime;

    public Date EndTime;

    public Color TextColor;

}
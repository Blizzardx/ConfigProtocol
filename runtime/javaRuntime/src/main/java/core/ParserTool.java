package core;


import config.Config;

import java.awt.*;
import java.lang.reflect.Method;
import java.text.DateFormat;
import java.util.Date;

public class ParserTool {

    public static Object Parser_Enum(String content,String enumName,CommonError errorMsg){
        int enumValue = Parser_Int32(content,errorMsg);
        if(errorMsg.errorMsg != ""){
            return null;
        }
        try {
            Class enumClass = Class. forName(enumName);
            Method getEnumMethod = enumClass.getMethod("getEnum",  int.class);

            return getEnumMethod.invoke(null,enumValue);
        }catch (Exception e){
            errorMsg.errorMsg = "error on parser cell " + content + " must be enum "+e.getMessage();
            return null;
        }
    }
    public static Integer Parser_Int32(String content, CommonError errorMsg)
    {
        errorMsg.errorMsg = "";
        try {
            Integer res = Integer.parseInt(content);
            return res;
        }catch (Exception e){
            errorMsg.errorMsg = "error on parser cell " + content + " must be int32 "+e.getMessage();
            return 0;
        }
    }
    public static Long Parser_Int64(String content, CommonError errorMsg)
    {
        errorMsg.errorMsg = "";
        try {
            Long res = Long.parseLong(content);
            return res;
        }catch (Exception e){
            errorMsg.errorMsg = "error on parser cell " + content + " must be int64 "+e.getMessage();
            return 0L;
        }
    }
    public static float Parser_Float32(String content, CommonError errorMsg)
    {
        errorMsg.errorMsg = "";
        try {
            Float res = Float.parseFloat(content);
            return res;
        }catch (Exception e){
            errorMsg.errorMsg = "error on parser cell " + content + " must be float32 "+e.getMessage();
            return 0;
        }
    }
    public static double Parser_Float64(String content, CommonError errorMsg)
    {
        errorMsg.errorMsg = "";
        try {
            Double res = Double.parseDouble(content);
            return res;
        }catch (Exception e){
            errorMsg.errorMsg = "error on parser cell " + content + " must be float64 "+e.getMessage();
            return 0;
        }
    }
    public static boolean Parser_Bool(String content, CommonError errorMsg)
    {
        errorMsg.errorMsg = "";
        if (content == "true")
        {
            return true;
        }
        else if (content == "false")
        {
            return false;
        }
        errorMsg.errorMsg =  "error on parser cell " + content + " must be true or false ";
        return false;
    }
    public static String  Parser_String(String content, CommonError errorMsg)
    {
        errorMsg.errorMsg = "";
        return content;
    }
    public static Color Parser_Color(String content, CommonError errorMsg)
    {
        errorMsg.errorMsg = "";
        if (content.startsWith("#"))
        {
            content = content.substring(1);
        }
        if (content.length() != 6 && content.length() != 8)
        {
            errorMsg.errorMsg = "error on parser color ";
            return null;
        }
        long color64 = 0;
        try
        {
            color64 = Long.parseLong(content, 16);
            if (content.length() == 8)
            {
                int red = (int) (color64 >> 24);
                int green = (int) ((color64 & 0x00FF0000) >> 16);
                int blue = (int) ((color64 & 0x0000FF00) >> 8);
                int alpha = (int) (color64 & 0x000000FF);
                return new Color(red, green, blue, alpha);
            }
            else
            {
                int red = (int) (color64 >> 16);
                int green = (int) ((color64 & 0x00FF00) >> 8);
                int blue = (int) (color64 & 0x0000FF);
                return new Color(red, green, blue);
            }
        }
        catch (Exception e)
        {
            errorMsg.errorMsg = "error on parser color " + e.getMessage();
            return null;
        }
    }
    public static Date Parser_DateTime(String content, CommonError errorMsg)
    {
        errorMsg.errorMsg = "";
        try {
            Date dateValue = DateFormat.getDateTimeInstance().parse(content);
            return dateValue;
        }catch (Exception e){
            errorMsg.errorMsg = "error on parser cell " + content + " must be  datetime "+e.getMessage();
            return null;
        }
    }
}

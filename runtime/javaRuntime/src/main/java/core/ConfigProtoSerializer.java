package core;

import config.Config;

import java.lang.reflect.Field;
import java.lang.reflect.Method;

public class ConfigProtoSerializer {
    public static void DeSerialize(byte[] content) throws Exception {
        config.Config.ConfigTable configContent = config.Config.ConfigTable.parseFrom(content);

        Class configType = null;
        Class configLineType = null;
        Object configInstance = null;

        try {
            configType = Class. forName(configContent.getPackageName()+"." + configContent.getConfigName());
            configLineType = Class. forName(configContent.getPackageName()+"." + configContent.getConfigName()+"Info");

            configInstance = configType.newInstance();
            Field configTableField = configInstance.getClass().getDeclaredField("Content");
            Object configContentInstance = configTableField.getType().newInstance();
            configTableField.set(configInstance, configContentInstance);

            //get add method
            Method addMethod = null;
            if(configContent.getType() == Config.ConfigType.typeList){
                addMethod = configTableField.getType().getMethod("add", Object.class);
            }else{
                addMethod =  configTableField.getType().getMethod("put", String.class, String.class);
            }

            for(int row=0;row<configContent.getContentCount();++row){
                config.Config.ConfigLine rowContent = configContent.getContent(row);
                Object lineInstance = configLineType.newInstance();

                for(int col =0 ;col <rowContent.getContentCount();++col){
                    if(col >= configContent.getFieldInfoListCount()){
                        System.out.println("col is out of range row: " + row + " col: " + col);
                        continue;
                    }
                    config.Config.ConfigFieldInfo fieldElem = configContent.getFieldInfoList(col);
                    CommonError errorMsg = new CommonError();
                    errorMsg.errorMsg = "";
                    Field cellFieldElement = lineInstance.getClass().getDeclaredField(fieldElem.getName());
                    if(null == cellFieldElement){
                        System.out.println("not found field " + fieldElem.getName());
                        continue;
                    }
                    Object cellInstance = parserField(configContent,cellFieldElement.getType(),fieldElem,rowContent.getContent(col),errorMsg);
                    if(errorMsg.errorMsg != ""){
                        System.out.println(errorMsg.errorMsg + " row: " + row + " col: " + col);
                        continue;
                    }
                    cellFieldElement.set(lineInstance, cellInstance);
                }

                if(configContent.getType() == Config.ConfigType.typeList){
                    addMethod.invoke(configContentInstance,lineInstance);

                }else{
                    addMethod.invoke(configContentInstance,lineInstance);
                }
            }


        } catch (ClassNotFoundException e) {
            e.printStackTrace();
        }
        System.out.print(configInstance);
    }
    private static Object parserField(config.Config.ConfigTable configContent,Class fieldType,config.Config.ConfigFieldInfo fieldInfo,String cell,CommonError error){
        try {
            if(fieldInfo.getIsList()){
                Object cellInstance = fieldType.newInstance();
                //get add method
                Method addMethod =fieldType.getMethod("add", Object.class);

                String[] childList = cell.split("|");
                for(int i=0;i<childList.length;++i){
                    Object cellElementInstance = doParserCell(configContent,childList[i],fieldInfo,error);
                    addMethod.invoke(cellInstance,cellElementInstance);
                }
                return cellInstance;
            }else{
                Object cellInstance = doParserCell(configContent,cell,fieldInfo,error);
                return cellInstance;
            }
        }catch (Exception e){
            error.errorMsg = "error on create field instance " + e.getMessage();
            return null;
        }
    }
    private static Object doParserCell(config.Config.ConfigTable configContent,String cell, config.Config.ConfigFieldInfo fieldInfo,CommonError errorMsg)
    {
        errorMsg.errorMsg = "";
        switch (fieldInfo.getType())
        {
            case typeInt32:
                return ParserTool.Parser_Int32(cell, errorMsg);
            case typeInt64:
                return ParserTool.Parser_Int64(cell,  errorMsg);
            case typeFloat32:
                return ParserTool.Parser_Float32(cell,  errorMsg);
            case typeFloat64:
                return ParserTool.Parser_Float64(cell,  errorMsg);
            case typeBool:
                return ParserTool.Parser_Bool(cell,  errorMsg);
            case typeString:
                return ParserTool.Parser_String(cell,  errorMsg);
            case typeEnum:
                String enumName = configContent.getPackageName()+"." + configContent.getConfigName()+fieldInfo.getParameter();
                return ParserTool.Parser_Enum(cell,  enumName, errorMsg);
            case typeDateTime:
                return ParserTool.Parser_DateTime(cell,  errorMsg);
            case typeColor:
                return ParserTool.Parser_Color(cell,  errorMsg);
            case typeClass:
                break;
        }
        errorMsg.errorMsg = "error on parser " + cell + " unknown type " + fieldInfo.getType();
        return null;
    }
}

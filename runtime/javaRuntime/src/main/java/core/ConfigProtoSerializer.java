package core;

import config.Config;

import java.lang.reflect.Field;
import java.lang.reflect.Method;

public class ConfigProtoSerializer {
    public static void DeSerialize(byte[] content) throws Exception {
        config.Config.ConfigTable configContent = config.Config.ConfigTable.parseFrom(content);

        Class configType = null;
        Object configInstance = null;

        try {
            configType = Class. forName(configContent.getPackageName()+"." + configContent.getConfigName());
            configInstance = configType.newInstance();
            Field f = configInstance.getClass().getDeclaredField("Content");
            Object lineContent = f.getType().newInstance();
            f.set(configInstance, lineContent);

            if(configContent.getType() == Config.ConfigType.typeList){
//                Method setReadOnly = t.getClass().getMethod("setReadOnly", String.class);
//                String s ="test2";
//                setReadOnly.invoke(t,s);
//                System.out.println(t.getReadOnly());
            }else{

            }

        } catch (ClassNotFoundException e) {
            e.printStackTrace();
        }
        System.out.print(configInstance);
    }
}

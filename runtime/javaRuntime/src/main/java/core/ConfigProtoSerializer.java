package core;

public class ConfigProtoSerializer {
    public static void DeSerialize(byte[] content) throws Exception {
        config.Config.ConfigTable configContent = config.Config.ConfigTable.parseFrom(content);

        Class configInstance = null;
        try {
            configInstance = Class. forName(configContent.getPackageName()+"." + configContent.getConfigName());

        } catch (ClassNotFoundException e) {
            e.printStackTrace();
        }
        System.out.print(configContent);
    }
}

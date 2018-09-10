package core;

public class ConfigProtoSerializer {
    public static void DeSerialize(byte[] content) throws Exception {
        config.Config.ConfigTable configContent = config.Config.ConfigTable.parseFrom(content);
        System.out.print(configContent);
    }
}

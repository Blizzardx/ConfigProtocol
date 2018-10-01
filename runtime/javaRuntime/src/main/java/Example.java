import core.ConfigProtoSerializer;

import java.io.File;
import java.io.FileInputStream;

public class Example {
    public static void main(String[] args){

        File file = new File("../../example/output/httpServer/BasicItem_Common.bytes");
        Long filelength = file.length();
        FileInputStream in = null;
        byte[] filecontent = new byte[filelength.intValue()];
        try {
            in = new FileInputStream(file);
            in.read(filecontent);
            in.close();
        } catch (Exception e) {
            e.printStackTrace();
        }

        try {
            Object configContent = ConfigProtoSerializer.DeSerialize(filecontent);
            System.out.println(configContent);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}

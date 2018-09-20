import core.ConfigProtoSerializer;

import java.io.File;
import java.io.FileInputStream;

public class Example {
    public static void main(String[] args){

        File file = new File("../../example/output/BasicItem_Common.bytes");
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
            ConfigProtoSerializer.DeSerialize(filecontent);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}

package config;

public enum BasicItem_CommonQuality
{
    yelow(1) ,
    blue(2),
    red(3),
    ;
    private final int value;

    private BasicItem_CommonQuality(int value) {
        this.value = value;
    }

    public int getValue() {
        return value;
    }

    public static BasicItem_CommonQuality getEnum(int value){
        for(int i=0;i<BasicItem_CommonQuality.values().length;++i){
            if(BasicItem_CommonQuality.values()[i].getValue() == value){
                return BasicItem_CommonQuality.values()[i];
            }
        }
        return null;
    }
}
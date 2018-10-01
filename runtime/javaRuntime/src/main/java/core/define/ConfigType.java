package core.define;


public enum ConfigType
{
    typeList(0),
    typeMap(1),
    ;

    private final int value;
    private ConfigType(int value) {
        this.value = value;
    }
    private ConfigType() {
        this.value = 0;
    }
}
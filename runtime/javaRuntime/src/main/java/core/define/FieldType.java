package core.define;


public enum FieldType {
    typeInt32(0),
    typeInt64(1),
    typeFloat32(2),
    typeFloat64(3),
    typeBool(4),
    typeString(5),
    typeClass(6),
    typeEnum(7),
    typeDateTime(8),
    typeColor(9),
    ;

    private final int value;

    FieldType(int value ) {
        this.value = value;
    }
    FieldType( ) {
        this.value = 0;
    }
}
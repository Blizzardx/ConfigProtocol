namespace ConfigProto.core
{
    public class Color
    {
        public Color(byte r, byte g, byte b, byte a = 0)
        {
            R = r;
            G = g;
            B = b;
            A = a;
        }

        public Color()
        {
            R = 0;
            G = 0;
            B = 0;
            A = 0;
        }
        public byte R;
        public byte G;
        public byte B;
        public byte A;
    }
}

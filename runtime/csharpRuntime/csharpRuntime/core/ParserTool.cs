
using System;
using System.Security.Policy;

namespace ConfigProto.core
{
    public class ParserTool
    {
        public static int Parser_Int32(string content, out string errorMsg)
        {
            errorMsg = "";
            int res = 0;
            if (!int.TryParse(content, out res))
            {
                errorMsg = "error on parser cell " + content + " must be int32 ";
                return 0;
            }
            return res;
        }
        public static long Parser_Int64(string content, out string errorMsg)
        {
            errorMsg = "";
            long res = 0;
            if (!long.TryParse(content, out res))
            {
                errorMsg = "error on parser cell " + content + " must be int64 ";
                return 0;
            }
            return res;
        }
        public static float Parser_Float32(string content, out string errorMsg)
        {
            errorMsg = "";
            float res = 0;
            if (!float.TryParse(content, out res))
            {
                errorMsg = "error on parser cell " + content + " must be float32 ";
                return 0;
            }
            return res;
        }
        public static double Parser_Float64(string content, out string errorMsg)
        {
            errorMsg = "";
            double res = 0;
            if (!double.TryParse(content, out res))
            {
                errorMsg = "error on parser cell " + content + " must be float64 ";
                return 0;
            }
            return res;
        }
        public static bool Parser_Bool(string content, out string errorMsg)
        {
            errorMsg = "";
            if (content == "true")
            {
                return true;
            }
            else if (content == "false")
            {
                return false;
            }
            errorMsg =  "error on parser cell " + content + " must be true or false ";
            return false;
        }
        public static string Parser_String(string content, out string errorMsg)
        {
            errorMsg = "";
            return content;
        }
        public static Color Parser_Color(string content, out string errorMsg)
        {
            errorMsg = "";
            if (content.StartsWith("#"))
            {
                content = content.Substring(1);
            }
            if (content.Length != 6 && content.Length != 8)
            {
                errorMsg = "error on parser color";
                return null;
            }
            long color64 = Parser_Int64(content, out errorMsg);
            if("" != errorMsg)
            {
                errorMsg = "error on parser color";
                return null;
            }
            if (content.Length == 8)
            {
                byte red = (byte) (color64 >> 24);
                byte green = (byte) ((color64 & 0x00FF0000) >> 16);
                byte blue = (byte) ((color64 & 0x0000FF00) >> 8);
                byte alpha = (byte) (color64 & 0x000000FF);
                return new Color(red, green, blue, alpha);
            }
            else
            {
                byte red = (byte) (color64 >> 16);
                byte green = (byte) ((color64 & 0x00FF00) >> 8);
                byte blue = (byte) (color64 & 0x0000FF);
                return new Color(red, green, blue);
            }
            return null;
        }
        public static DateTime Parser_DateTime(string content, out string errorMsg)
        {
            errorMsg = "";
            DateTime dateValue;
            if (!DateTime.TryParse(content, out dateValue))
            {
                errorMsg = "error on parser cell " + content + " must be  datetime ";
                return dateValue;
            }
            return dateValue;
        }
        public static void CheckValueLimit_Int32(int value, string limitMin, string limitMax, out string errorMsg)
        {
            errorMsg = "";

            int min;
            int max;
            if (limitMin != "")
            {
                min = Parser_Int32(limitMin, out errorMsg);
                if (errorMsg != ""){return;}
                if (value < min)
                {
                    errorMsg = "error on check limit ";
                    return;
                }
            }
            if (limitMax != "")
            {
                max = Parser_Int32(limitMax, out errorMsg);
                if (errorMsg != ""){return;}
                if (value > max)
                {
                    errorMsg = "error on check limit ";
                }
            }
            return;
        }
        public static void CheckValueLimit_Int64(long value, string limitMin, string limitMax, out string errorMsg)
        {
            errorMsg = "";

            long min;
            long max;
            if (limitMin != "")
            {
                min = Parser_Int64(limitMin, out errorMsg);
                if (errorMsg != ""){return;}
                if (value < min)
                {
                    errorMsg = "error on check limit ";
                    return;
                }
            }
            if (limitMax != "")
            {
                max = Parser_Int64(limitMax, out errorMsg);
                if (errorMsg != ""){return;}
                if (value > max)
                {
                    errorMsg = "error on check limit ";
                }
            }
            return;
        }
        public static void CheckValueLimit_Float32(float value, string limitMin, string limitMax, out string errorMsg)
        {
            errorMsg = "";

            float min;
            float max;
            if (limitMin != "")
            {
                min = Parser_Float32(limitMin, out errorMsg);
                if (errorMsg != ""){return;}
                if (value < min)
                {
                    errorMsg = "error on check limit ";
                    return;
                }
            }
            if (limitMax != "")
            {
                max = Parser_Float32(limitMax, out errorMsg);
                if (errorMsg != ""){return;}
                if (value > max)
                {
                    errorMsg = "error on check limit ";
                }
            }
            return;
        }
        public static void CheckValueLimit_Float64(double value, string limitMin, string limitMax, out string errorMsg)
        {
            errorMsg = "";

            double min;
            double max;
            if (limitMin != "")
            {
                min = Parser_Float64(limitMin, out errorMsg);
                if (errorMsg != ""){return;}
                if (value < min)
                {
                    errorMsg = "error on check limit ";
                    return;
                }
            }
            if (limitMax != "")
            {
                max = Parser_Float64(limitMax, out errorMsg);
                if (errorMsg != ""){return;}
                if (value > max)
                {
                    errorMsg = "error on check limit ";
                }
            }
            return;
        }
    }
}

using System;
using config;
using ConfigProto.core;
using LitJson;

namespace csharpRuntime
{
    public class LitJsonSe : IJsonSerializer
    {
        public T DeSerialize<T>(byte[] content)
        {
            string str = System.Text.Encoding.UTF8.GetString(content);
            T config = JsonMapper.ToObject<T>(str);
            return config;
        }
    }

    class Program
    { 
        static void Main(string[] args)
        {
            var configContent = System.IO.File.ReadAllBytes("../../../../../example/output/client/BasicItem_Common.bytes");
            string errorMsg = "";
            var res = ConfigProtoSerializer.DeSerialize<BasicItem_Common>(new LitJsonSe(), configContent, out errorMsg);
            Console.WriteLine(errorMsg);
           
            Console.ReadLine();
        }
    }
}

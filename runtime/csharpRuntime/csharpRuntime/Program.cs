using System;
using config;
using ConfigProto.core;

namespace csharpRuntime
{
    class Program
    { 
        static void Main(string[] args)
        {
            var configContent = System.IO.File.ReadAllBytes("../../../../../example/output/BasicItem_Common.bytes");
            string errorMsg = "";
            var res = ConfigProtoSerializer.DeSerialize(configContent, out errorMsg) as BasicItem_Common;
            Console.WriteLine(errorMsg);
           
            Console.ReadLine();
        }
    }
}

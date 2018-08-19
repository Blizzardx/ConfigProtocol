using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using ConfigProto.core;

namespace csharpRuntime
{
    class Program
    {
        static void Main(string[] args)
        {
            var configContent = System.IO.File.ReadAllBytes("../../../../../example/output/BasicItem_Common.bytes");
            string errorMsg = "";
            ConfigProtoSerializer.DeSerialize(configContent, out errorMsg);
            Console.WriteLine(errorMsg);
            Console.ReadLine();
        }
    }
}

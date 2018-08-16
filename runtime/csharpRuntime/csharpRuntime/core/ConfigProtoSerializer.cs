using System;
using Config;
using Google.Protobuf;

namespace ConfigProto.core
{
    public class ConfigProtoSerializer
    {
        public Object DeSerialize(byte[] configContent,out string error)
        {
            error = "";
            Config.ConfigTable config = new ConfigTable();

            try
            {
                config.MergeFrom(configContent);
            }
            catch (Exception e)
            {
                error = "error on parser config content " + e.Message;
                return null;
            }
            
            
            return null;
        }
    }
}

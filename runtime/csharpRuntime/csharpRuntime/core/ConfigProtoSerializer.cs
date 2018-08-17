using System;
using System.Reflection;
using Config;
using Google.Protobuf;

namespace ConfigProto.core
{
    public class ConfigProtoSerializer
    {
        public Object DeSerialize(byte[] configContent, out string error)
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
                Logger.Instance.LogError(error);
                return null;
            }

            Type configTyep = ReflectionManager.Instance.GetTypeByName(config.PackageName);
            if (null == configTyep)
            {
                error = "error on get type by name " + config.PackageName;
                Logger.Instance.LogError(error);
                return null;
            }
            var configInstance = Activator.CreateInstance(configTyep);

            return configInstance;
        }
        public  PropertyInfo[] GetPropertyInfos(Type type)
        {
            return type.GetProperties(BindingFlags.Public | BindingFlags.Instance);
        }
        public  void AutoMapping<S, T>(S s, T t)
        {
            
            PropertyInfo[] pps = GetPropertyInfos(s.GetType());
            Type target = t.GetType();

            foreach (var pp in pps)
            {
                PropertyInfo targetPP = target.GetProperty(pp.Name);
                object value = pp.GetValue(s, null);

                if (targetPP != null && value != null)
                {
                    targetPP.SetValue(t, value, null);
                }
            }
        }
    }
}

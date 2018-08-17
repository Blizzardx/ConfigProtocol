using System;
using System.Collections.Generic;
using System.Reflection;

namespace ConfigProto.core
{
    public class ReflectionManager
    {
        private static readonly ReflectionManager m_Instance = new ReflectionManager();

        public static ReflectionManager Instance
        {
            get { return m_Instance; }
        }

        private Dictionary<string, Type> m_ClassFindMap;
        private List<Type> m_ClassList;

        private bool m_bIsInit;

        private ReflectionManager()
        {
            //
            CheckInit();
        }

        public void CheckInit()
        {
            if (m_bIsInit)
            {
                return;
            }
            m_ClassFindMap = new Dictionary<string, Type>();
            Assembly assem = Assembly.GetAssembly(typeof (ReflectionManager));
            m_ClassList = new List<Type>(assem.GetTypes());

            for (int i = 0; i < m_ClassList.Count; ++i)
            {
                var elem = m_ClassList[i];
                if (!m_ClassFindMap.ContainsKey(elem.Name))
                {
                    m_ClassFindMap.Add(elem.Name, elem);
                }
            }
        }

        public Type GetTypeByName(string name)
        {
            Type res = null;
            if (!m_ClassFindMap.TryGetValue(name, out res))
            {
                // 
                Logger.Instance.LogError("error on get type by name " + name);
            }
            return res;
        }

        public List<Type> GetTypeByBase(Type baseType)
        {
            List<Type> resList = new List<Type>();
            for (int i = 0; i < m_ClassList.Count; ++i)
            {
                var elem = m_ClassList[i];
                if (!elem.IsInterface && !elem.IsAbstract)
                {
                    if (elem.BaseType == baseType)
                    {
                        // add to list
                        resList.Add(elem);
                    }
                    else if ((baseType.IsInterface || baseType.IsAbstract) && baseType.IsAssignableFrom(elem))
                    {
                        // add to list
                        resList.Add(elem);
                    }
                }
            }
            return resList;
        }
    }
}
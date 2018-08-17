
namespace ConfigProto.core
{
    public interface ILoger
    {
        void Log(string content);
        void LogError(string content);
        void LogWarning(string content);
    }

    public class DefaultLogger : ILoger
    {
        public void Log(string content)
        {
            System.Console.WriteLine(content);
        }

        public void LogError(string content)
        {
            System.Console.WriteLine(content);
        }

        public void LogWarning(string content)
        {
            System.Console.WriteLine(content);
        }
    }
    public class Logger
    {
        private static readonly Logger m_Instance = new Logger();

        public static Logger Instance
        {
            get { return m_Instance; }
        }

        private ILoger m_Logger = new DefaultLogger();

        public void SetLogger(ILoger logger)
        {
            m_Logger = logger;
        }
        public void Log(string content) { m_Logger.Log(content); }
        public void LogError(string content) { m_Logger.LogError(content); }
        public void LogWarning(string content) { m_Logger.LogWarning(content); }
    }
}

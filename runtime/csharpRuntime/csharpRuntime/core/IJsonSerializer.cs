
namespace ConfigProto.core
{
    public interface IJsonSerializer
    {
        T DeSerialize<T>(byte[] content);
    }
}

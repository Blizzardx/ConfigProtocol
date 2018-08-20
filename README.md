# ConfigProtocol

用途描述
本库用来定义 配置文件协议，通过定义好的excel文件的文件头，来生成对应的装载类。同时提供了运行的的自动装载和非运行时的生成工具支持

使用方法
1 先定义好 配置文件 excel 。
2 通过tool 下边的 exportTool，打包配置文件，生成 序列化好的配置文件，以及对应语言的装载类
3 根据自己项目的使用语言（java，c#，golang） ，拷贝 runtime下的运行时库。里边提供了运行时的配置文件装载、

// golang
运行时库在 runtime/goRuntime 下，不过运行时库需要依赖本项目，不需要拷贝，直接依赖本项目即可
使用方法 调用 goRuntime.LoadConfigByContent 进行反序列化

// c#
运行时库在 runtime/csharpRuntime 下，需要拷贝 runtime/csharpRuntime/csharpRuntime/core 下边的所有文件，到本地目录
使用方法，调用 ConfigProtoSerializer.DeSerialize 进行反序列化
同时c#版本支持自定义log 接口 需要通过 Logger.Instance.SetLogger(ILogger) 设置

// java

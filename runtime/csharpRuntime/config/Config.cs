// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: config.proto
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Config {

  /// <summary>Holder for reflection information generated from config.proto</summary>
  public static partial class ConfigReflection {

    #region Descriptor
    /// <summary>File descriptor for config.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static ConfigReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Cgxjb25maWcucHJvdG8SBmNvbmZpZyIdCgpDb25maWdMaW5lEg8KB2NvbnRl",
            "bnQYASADKAkiNAoNQ29uZmlnQ29udGVudBIjCgdjb250ZW50GAEgAygLMhIu",
            "Y29uZmlnLkNvbmZpZ0xpbmUiYwoPQ29uZmlnRmllbGRJbmZvEgwKBG5hbWUY",
            "ASABKAkSHwoEdHlwZRgCIAEoDjIRLmNvbmZpZy5GaWVsZFR5cGUSDgoGaXNM",
            "aXN0GAMgASgIEhEKCXBhcmFtZXRlchgEIAEoCSI0ChVDb25maWdFbnVtRWxl",
            "bWVudEluZm8SDAoEbmFtZRgBIAEoCRINCgV2YWx1ZRgCIAEoBSJMCg5Db25m",
            "aWdFbnVtSW5mbxIMCgRuYW1lGAEgASgJEiwKBXZhbHVlGAIgAygLMh0uY29u",
            "ZmlnLkNvbmZpZ0VudW1FbGVtZW50SW5mbyLIAQoLQ29uZmlnVGFibGUSIAoE",
            "dHlwZRgBIAEoDjISLmNvbmZpZy5Db25maWdUeXBlEhQKDGtleUZpZWxkTmFt",
            "ZRgCIAEoCRIuCg1maWVsZEluZm9MaXN0GAMgAygLMhcuY29uZmlnLkNvbmZp",
            "Z0ZpZWxkSW5mbxIsCgxlbnVtSW5mb0xpc3QYBCADKAsyFi5jb25maWcuQ29u",
            "ZmlnRW51bUluZm8SIwoHY29udGVudBgFIAMoCzISLmNvbmZpZy5Db25maWdM",
            "aW5lKqcBCglGaWVsZFR5cGUSDQoJdHlwZUludDMyEAASDQoJdHlwZUludDY0",
            "EAESDwoLdHlwZUZsb2F0MzIQAhIPCgt0eXBlRmxvYXQ2NBADEgwKCHR5cGVC",
            "b29sEAQSDgoKdHlwZVN0cmluZxAFEg0KCXR5cGVDbGFzcxAGEgwKCHR5cGVF",
            "bnVtEAcSEAoMdHlwZURhdGVUaW1lEAgSDQoJdHlwZUNvbG9yEAkqJwoKQ29u",
            "ZmlnVHlwZRIMCgh0eXBlTGlzdBAAEgsKB3R5cGVNYXAQAWIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::Config.FieldType), typeof(global::Config.ConfigType), }, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Config.ConfigLine), global::Config.ConfigLine.Parser, new[]{ "Content" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Config.ConfigContent), global::Config.ConfigContent.Parser, new[]{ "Content" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Config.ConfigFieldInfo), global::Config.ConfigFieldInfo.Parser, new[]{ "Name", "Type", "IsList", "Parameter" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Config.ConfigEnumElementInfo), global::Config.ConfigEnumElementInfo.Parser, new[]{ "Name", "Value" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Config.ConfigEnumInfo), global::Config.ConfigEnumInfo.Parser, new[]{ "Name", "Value" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Config.ConfigTable), global::Config.ConfigTable.Parser, new[]{ "Type", "KeyFieldName", "FieldInfoList", "EnumInfoList", "Content" }, null, null, null)
          }));
    }
    #endregion

  }
  #region Enums
  public enum FieldType {
    [pbr::OriginalName("typeInt32")] TypeInt32 = 0,
    [pbr::OriginalName("typeInt64")] TypeInt64 = 1,
    [pbr::OriginalName("typeFloat32")] TypeFloat32 = 2,
    [pbr::OriginalName("typeFloat64")] TypeFloat64 = 3,
    [pbr::OriginalName("typeBool")] TypeBool = 4,
    [pbr::OriginalName("typeString")] TypeString = 5,
    [pbr::OriginalName("typeClass")] TypeClass = 6,
    [pbr::OriginalName("typeEnum")] TypeEnum = 7,
    [pbr::OriginalName("typeDateTime")] TypeDateTime = 8,
    [pbr::OriginalName("typeColor")] TypeColor = 9,
  }

  public enum ConfigType {
    [pbr::OriginalName("typeList")] TypeList = 0,
    [pbr::OriginalName("typeMap")] TypeMap = 1,
  }

  #endregion

  #region Messages
  public sealed partial class ConfigLine : pb::IMessage<ConfigLine> {
    private static readonly pb::MessageParser<ConfigLine> _parser = new pb::MessageParser<ConfigLine>(() => new ConfigLine());
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<ConfigLine> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Config.ConfigReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ConfigLine() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ConfigLine(ConfigLine other) : this() {
      content_ = other.content_.Clone();
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ConfigLine Clone() {
      return new ConfigLine(this);
    }

    /// <summary>Field number for the "content" field.</summary>
    public const int ContentFieldNumber = 1;
    private static readonly pb::FieldCodec<string> _repeated_content_codec
        = pb::FieldCodec.ForString(10);
    private readonly pbc::RepeatedField<string> content_ = new pbc::RepeatedField<string>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<string> Content {
      get { return content_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as ConfigLine);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(ConfigLine other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if(!content_.Equals(other.content_)) return false;
      return true;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      hash ^= content_.GetHashCode();
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      content_.WriteTo(output, _repeated_content_codec);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      size += content_.CalculateSize(_repeated_content_codec);
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(ConfigLine other) {
      if (other == null) {
        return;
      }
      content_.Add(other.content_);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            input.SkipLastField();
            break;
          case 10: {
            content_.AddEntriesFrom(input, _repeated_content_codec);
            break;
          }
        }
      }
    }

  }

  public sealed partial class ConfigContent : pb::IMessage<ConfigContent> {
    private static readonly pb::MessageParser<ConfigContent> _parser = new pb::MessageParser<ConfigContent>(() => new ConfigContent());
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<ConfigContent> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Config.ConfigReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ConfigContent() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ConfigContent(ConfigContent other) : this() {
      content_ = other.content_.Clone();
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ConfigContent Clone() {
      return new ConfigContent(this);
    }

    /// <summary>Field number for the "content" field.</summary>
    public const int ContentFieldNumber = 1;
    private static readonly pb::FieldCodec<global::Config.ConfigLine> _repeated_content_codec
        = pb::FieldCodec.ForMessage(10, global::Config.ConfigLine.Parser);
    private readonly pbc::RepeatedField<global::Config.ConfigLine> content_ = new pbc::RepeatedField<global::Config.ConfigLine>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<global::Config.ConfigLine> Content {
      get { return content_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as ConfigContent);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(ConfigContent other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if(!content_.Equals(other.content_)) return false;
      return true;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      hash ^= content_.GetHashCode();
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      content_.WriteTo(output, _repeated_content_codec);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      size += content_.CalculateSize(_repeated_content_codec);
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(ConfigContent other) {
      if (other == null) {
        return;
      }
      content_.Add(other.content_);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            input.SkipLastField();
            break;
          case 10: {
            content_.AddEntriesFrom(input, _repeated_content_codec);
            break;
          }
        }
      }
    }

  }

  public sealed partial class ConfigFieldInfo : pb::IMessage<ConfigFieldInfo> {
    private static readonly pb::MessageParser<ConfigFieldInfo> _parser = new pb::MessageParser<ConfigFieldInfo>(() => new ConfigFieldInfo());
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<ConfigFieldInfo> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Config.ConfigReflection.Descriptor.MessageTypes[2]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ConfigFieldInfo() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ConfigFieldInfo(ConfigFieldInfo other) : this() {
      name_ = other.name_;
      type_ = other.type_;
      isList_ = other.isList_;
      parameter_ = other.parameter_;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ConfigFieldInfo Clone() {
      return new ConfigFieldInfo(this);
    }

    /// <summary>Field number for the "name" field.</summary>
    public const int NameFieldNumber = 1;
    private string name_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Name {
      get { return name_; }
      set {
        name_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "type" field.</summary>
    public const int TypeFieldNumber = 2;
    private global::Config.FieldType type_ = 0;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Config.FieldType Type {
      get { return type_; }
      set {
        type_ = value;
      }
    }

    /// <summary>Field number for the "isList" field.</summary>
    public const int IsListFieldNumber = 3;
    private bool isList_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool IsList {
      get { return isList_; }
      set {
        isList_ = value;
      }
    }

    /// <summary>Field number for the "parameter" field.</summary>
    public const int ParameterFieldNumber = 4;
    private string parameter_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Parameter {
      get { return parameter_; }
      set {
        parameter_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as ConfigFieldInfo);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(ConfigFieldInfo other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Name != other.Name) return false;
      if (Type != other.Type) return false;
      if (IsList != other.IsList) return false;
      if (Parameter != other.Parameter) return false;
      return true;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Name.Length != 0) hash ^= Name.GetHashCode();
      if (Type != 0) hash ^= Type.GetHashCode();
      if (IsList != false) hash ^= IsList.GetHashCode();
      if (Parameter.Length != 0) hash ^= Parameter.GetHashCode();
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Name.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(Name);
      }
      if (Type != 0) {
        output.WriteRawTag(16);
        output.WriteEnum((int) Type);
      }
      if (IsList != false) {
        output.WriteRawTag(24);
        output.WriteBool(IsList);
      }
      if (Parameter.Length != 0) {
        output.WriteRawTag(34);
        output.WriteString(Parameter);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Name.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Name);
      }
      if (Type != 0) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) Type);
      }
      if (IsList != false) {
        size += 1 + 1;
      }
      if (Parameter.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Parameter);
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(ConfigFieldInfo other) {
      if (other == null) {
        return;
      }
      if (other.Name.Length != 0) {
        Name = other.Name;
      }
      if (other.Type != 0) {
        Type = other.Type;
      }
      if (other.IsList != false) {
        IsList = other.IsList;
      }
      if (other.Parameter.Length != 0) {
        Parameter = other.Parameter;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            input.SkipLastField();
            break;
          case 10: {
            Name = input.ReadString();
            break;
          }
          case 16: {
            type_ = (global::Config.FieldType) input.ReadEnum();
            break;
          }
          case 24: {
            IsList = input.ReadBool();
            break;
          }
          case 34: {
            Parameter = input.ReadString();
            break;
          }
        }
      }
    }

  }

  public sealed partial class ConfigEnumElementInfo : pb::IMessage<ConfigEnumElementInfo> {
    private static readonly pb::MessageParser<ConfigEnumElementInfo> _parser = new pb::MessageParser<ConfigEnumElementInfo>(() => new ConfigEnumElementInfo());
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<ConfigEnumElementInfo> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Config.ConfigReflection.Descriptor.MessageTypes[3]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ConfigEnumElementInfo() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ConfigEnumElementInfo(ConfigEnumElementInfo other) : this() {
      name_ = other.name_;
      value_ = other.value_;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ConfigEnumElementInfo Clone() {
      return new ConfigEnumElementInfo(this);
    }

    /// <summary>Field number for the "name" field.</summary>
    public const int NameFieldNumber = 1;
    private string name_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Name {
      get { return name_; }
      set {
        name_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "value" field.</summary>
    public const int ValueFieldNumber = 2;
    private int value_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int Value {
      get { return value_; }
      set {
        value_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as ConfigEnumElementInfo);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(ConfigEnumElementInfo other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Name != other.Name) return false;
      if (Value != other.Value) return false;
      return true;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Name.Length != 0) hash ^= Name.GetHashCode();
      if (Value != 0) hash ^= Value.GetHashCode();
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Name.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(Name);
      }
      if (Value != 0) {
        output.WriteRawTag(16);
        output.WriteInt32(Value);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Name.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Name);
      }
      if (Value != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Value);
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(ConfigEnumElementInfo other) {
      if (other == null) {
        return;
      }
      if (other.Name.Length != 0) {
        Name = other.Name;
      }
      if (other.Value != 0) {
        Value = other.Value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            input.SkipLastField();
            break;
          case 10: {
            Name = input.ReadString();
            break;
          }
          case 16: {
            Value = input.ReadInt32();
            break;
          }
        }
      }
    }

  }

  public sealed partial class ConfigEnumInfo : pb::IMessage<ConfigEnumInfo> {
    private static readonly pb::MessageParser<ConfigEnumInfo> _parser = new pb::MessageParser<ConfigEnumInfo>(() => new ConfigEnumInfo());
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<ConfigEnumInfo> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Config.ConfigReflection.Descriptor.MessageTypes[4]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ConfigEnumInfo() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ConfigEnumInfo(ConfigEnumInfo other) : this() {
      name_ = other.name_;
      value_ = other.value_.Clone();
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ConfigEnumInfo Clone() {
      return new ConfigEnumInfo(this);
    }

    /// <summary>Field number for the "name" field.</summary>
    public const int NameFieldNumber = 1;
    private string name_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Name {
      get { return name_; }
      set {
        name_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "value" field.</summary>
    public const int ValueFieldNumber = 2;
    private static readonly pb::FieldCodec<global::Config.ConfigEnumElementInfo> _repeated_value_codec
        = pb::FieldCodec.ForMessage(18, global::Config.ConfigEnumElementInfo.Parser);
    private readonly pbc::RepeatedField<global::Config.ConfigEnumElementInfo> value_ = new pbc::RepeatedField<global::Config.ConfigEnumElementInfo>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<global::Config.ConfigEnumElementInfo> Value {
      get { return value_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as ConfigEnumInfo);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(ConfigEnumInfo other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Name != other.Name) return false;
      if(!value_.Equals(other.value_)) return false;
      return true;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Name.Length != 0) hash ^= Name.GetHashCode();
      hash ^= value_.GetHashCode();
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Name.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(Name);
      }
      value_.WriteTo(output, _repeated_value_codec);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Name.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Name);
      }
      size += value_.CalculateSize(_repeated_value_codec);
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(ConfigEnumInfo other) {
      if (other == null) {
        return;
      }
      if (other.Name.Length != 0) {
        Name = other.Name;
      }
      value_.Add(other.value_);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            input.SkipLastField();
            break;
          case 10: {
            Name = input.ReadString();
            break;
          }
          case 18: {
            value_.AddEntriesFrom(input, _repeated_value_codec);
            break;
          }
        }
      }
    }

  }

  public sealed partial class ConfigTable : pb::IMessage<ConfigTable> {
    private static readonly pb::MessageParser<ConfigTable> _parser = new pb::MessageParser<ConfigTable>(() => new ConfigTable());
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<ConfigTable> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Config.ConfigReflection.Descriptor.MessageTypes[5]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ConfigTable() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ConfigTable(ConfigTable other) : this() {
      type_ = other.type_;
      keyFieldName_ = other.keyFieldName_;
      fieldInfoList_ = other.fieldInfoList_.Clone();
      enumInfoList_ = other.enumInfoList_.Clone();
      content_ = other.content_.Clone();
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ConfigTable Clone() {
      return new ConfigTable(this);
    }

    /// <summary>Field number for the "type" field.</summary>
    public const int TypeFieldNumber = 1;
    private global::Config.ConfigType type_ = 0;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Config.ConfigType Type {
      get { return type_; }
      set {
        type_ = value;
      }
    }

    /// <summary>Field number for the "keyFieldName" field.</summary>
    public const int KeyFieldNameFieldNumber = 2;
    private string keyFieldName_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string KeyFieldName {
      get { return keyFieldName_; }
      set {
        keyFieldName_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "fieldInfoList" field.</summary>
    public const int FieldInfoListFieldNumber = 3;
    private static readonly pb::FieldCodec<global::Config.ConfigFieldInfo> _repeated_fieldInfoList_codec
        = pb::FieldCodec.ForMessage(26, global::Config.ConfigFieldInfo.Parser);
    private readonly pbc::RepeatedField<global::Config.ConfigFieldInfo> fieldInfoList_ = new pbc::RepeatedField<global::Config.ConfigFieldInfo>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<global::Config.ConfigFieldInfo> FieldInfoList {
      get { return fieldInfoList_; }
    }

    /// <summary>Field number for the "enumInfoList" field.</summary>
    public const int EnumInfoListFieldNumber = 4;
    private static readonly pb::FieldCodec<global::Config.ConfigEnumInfo> _repeated_enumInfoList_codec
        = pb::FieldCodec.ForMessage(34, global::Config.ConfigEnumInfo.Parser);
    private readonly pbc::RepeatedField<global::Config.ConfigEnumInfo> enumInfoList_ = new pbc::RepeatedField<global::Config.ConfigEnumInfo>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<global::Config.ConfigEnumInfo> EnumInfoList {
      get { return enumInfoList_; }
    }

    /// <summary>Field number for the "content" field.</summary>
    public const int ContentFieldNumber = 5;
    private static readonly pb::FieldCodec<global::Config.ConfigLine> _repeated_content_codec
        = pb::FieldCodec.ForMessage(42, global::Config.ConfigLine.Parser);
    private readonly pbc::RepeatedField<global::Config.ConfigLine> content_ = new pbc::RepeatedField<global::Config.ConfigLine>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<global::Config.ConfigLine> Content {
      get { return content_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as ConfigTable);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(ConfigTable other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Type != other.Type) return false;
      if (KeyFieldName != other.KeyFieldName) return false;
      if(!fieldInfoList_.Equals(other.fieldInfoList_)) return false;
      if(!enumInfoList_.Equals(other.enumInfoList_)) return false;
      if(!content_.Equals(other.content_)) return false;
      return true;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Type != 0) hash ^= Type.GetHashCode();
      if (KeyFieldName.Length != 0) hash ^= KeyFieldName.GetHashCode();
      hash ^= fieldInfoList_.GetHashCode();
      hash ^= enumInfoList_.GetHashCode();
      hash ^= content_.GetHashCode();
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Type != 0) {
        output.WriteRawTag(8);
        output.WriteEnum((int) Type);
      }
      if (KeyFieldName.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(KeyFieldName);
      }
      fieldInfoList_.WriteTo(output, _repeated_fieldInfoList_codec);
      enumInfoList_.WriteTo(output, _repeated_enumInfoList_codec);
      content_.WriteTo(output, _repeated_content_codec);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Type != 0) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) Type);
      }
      if (KeyFieldName.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(KeyFieldName);
      }
      size += fieldInfoList_.CalculateSize(_repeated_fieldInfoList_codec);
      size += enumInfoList_.CalculateSize(_repeated_enumInfoList_codec);
      size += content_.CalculateSize(_repeated_content_codec);
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(ConfigTable other) {
      if (other == null) {
        return;
      }
      if (other.Type != 0) {
        Type = other.Type;
      }
      if (other.KeyFieldName.Length != 0) {
        KeyFieldName = other.KeyFieldName;
      }
      fieldInfoList_.Add(other.fieldInfoList_);
      enumInfoList_.Add(other.enumInfoList_);
      content_.Add(other.content_);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            input.SkipLastField();
            break;
          case 8: {
            type_ = (global::Config.ConfigType) input.ReadEnum();
            break;
          }
          case 18: {
            KeyFieldName = input.ReadString();
            break;
          }
          case 26: {
            fieldInfoList_.AddEntriesFrom(input, _repeated_fieldInfoList_codec);
            break;
          }
          case 34: {
            enumInfoList_.AddEntriesFrom(input, _repeated_enumInfoList_codec);
            break;
          }
          case 42: {
            content_.AddEntriesFrom(input, _repeated_content_codec);
            break;
          }
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code

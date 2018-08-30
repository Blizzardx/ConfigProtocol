package exportTool

import "github.com/Blizzardx/ConfigProtocol/common"

type TipMessageDefine int32

const (
	TipMessageDefine_UnknownFieldType TipMessageDefine = iota
	TipMessageDefine_ErrorOnParserEnum
	TipMessageDefine_ErrorOnParserEnumRepeated
	TipMessageDefine_ErrorOnParserEnumEmpty
	TipMessageDefine_ErrorOnParserEnumFormateError
	TipMessageDefine_PositionMark
	TipMessageDefine_ErrorFormate
	TipMessageDefine_OutofLineRange
	TipMessageDefine_ErrorFieldRepeated
	TipMessageDefine_ErrorFieldError
	TipMessageDefine_ErrorFieldKeyNotFound
	TipMessageDefine_ErrorFieldKeyCannotBeList
	TipMessageDefine_ErrorFieldType
	TipMessageDefine_ExportWithError
	TipMessageDefine_ErrorOnParserReference
	TipMessageDefine_ReferenceFieldNotFound
	TipMessageDefine_ReferenceCheckError
	TipMessageDefine_ReferenceIndexerror
)

func getTipMessage(msgType TipMessageDefine, args ...interface{}) string {
	res := ""
	switch msgType {
	case TipMessageDefine_UnknownFieldType:
		res = "未知的属性类型 {0}"
	case TipMessageDefine_ErrorOnParserEnum:
		res = "解析枚举 {0} 的时候发生错误 {1} {2} 必须是int32类型"
	case TipMessageDefine_ErrorOnParserEnumRepeated:
		res = "解析枚举 {0} 的时候发生错误,枚举类型重复 {1} "
	case TipMessageDefine_ErrorOnParserEnumEmpty:
		res = "解析枚举 {0} 的时候发生错误,枚举类型不能为空 {1} "
	case TipMessageDefine_ErrorOnParserEnumFormateError:
		res = "解析枚举 {0} 的时候发生错误,枚举类型描述格式不正确,参考 quality:black=1|white=2|yellow=3"
	case TipMessageDefine_PositionMark:
		res = "在 {0}：行 {1}：列 "
	case TipMessageDefine_ErrorFormate:
		res = "{0} 发生错误，错误内容：{1} "
	case TipMessageDefine_OutofLineRange:
		res = "下标长度超过了原本内容的行数"
	case TipMessageDefine_ErrorFieldRepeated:
		res = "属性名字 {0} 重复"
	case TipMessageDefine_ErrorFieldError:
		res = "属性名字 {0} 不合法，属性名字不能为空"
	case TipMessageDefine_ErrorFieldKeyNotFound:
		res = "定义了map的主键 属性名字 {0} 未能在配置表中找到"
	case TipMessageDefine_ErrorFieldKeyCannotBeList:
		res = "定义了map的主键 属性名字 {0} ,但是在配置表中，这一列被定义成了数组，主键不能是数组"
	case TipMessageDefine_ErrorFieldType:
		res = "文件类型定义不合法，{0} 必须是 list 或者 map"
	case TipMessageDefine_ExportWithError:
		res = "导出配置文件 {0} 时发生错误，错误内容: {1}"
	case TipMessageDefine_ErrorOnParserReference:
		res = "解析引用 {0} 的时候发生错误,引用描述格式不正确,参考 itemConfig:itemId"
	case TipMessageDefine_ReferenceFieldNotFound:
		res = "属性名字 {0} 未能在配置表 {1} 中找到"
	case TipMessageDefine_ReferenceCheckError:
		res = "解析引用 时发生错误，在 {0}：行 {1}：列 ,错误内容: {2}"
	case TipMessageDefine_ReferenceIndexerror:
		res = "索引位置错误"
	}
	return common.StringFormate(res, args...)
}

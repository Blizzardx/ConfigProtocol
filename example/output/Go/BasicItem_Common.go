// Generated by gen-tool
// DO NOT EDIT!
package config

type BasicItem_CommonQuality int32

const (
	BasicItem_CommonQuality_yelow BasicItem_CommonQuality = 1

	BasicItem_CommonQuality_blue BasicItem_CommonQuality = 2

	BasicItem_CommonQuality_red BasicItem_CommonQuality = 3
)

type BasicItem_Common struct {
	Content map[int32]*BasicItem_CommonInfo
}
type BasicItem_CommonInfo struct {
	Id int32

	Icon string

	Quality int32

	Price int32

	LimitNum int32

	Acauire int32

	ConsumeItem string

	ConsumeCoin int32

	FormatIndex int32

	Quality1 BasicItem_CommonQuality
}

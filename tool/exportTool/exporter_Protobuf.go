package exportTool

import (
	"github.com/Blizzardx/ConfigProtocol/pbConfig"
	"github.com/golang/protobuf/proto"
)

type ConfigSerializer_Protobuf struct {
}

func (self *ConfigSerializer_Protobuf) Export(configContent *config.ConfigTable) ([]byte, error) {
	protoBufByteContent, err := proto.Marshal(configContent)
	return protoBufByteContent, err
}

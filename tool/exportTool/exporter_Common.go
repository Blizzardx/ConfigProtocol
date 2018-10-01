package exportTool

import (
	"encoding/json"
	"github.com/Blizzardx/ConfigProtocol/define"
)

type ConfigSerializer_Common struct {
}

func (self *ConfigSerializer_Common) Export(configContent *define.ConfigTable) ([]byte, error) {
	protoBufByteContent, err := json.Marshal(configContent)
	return protoBufByteContent, err
}

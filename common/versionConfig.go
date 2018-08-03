package common

type VersionConfig struct {
	Sign     string                  `xml:"sign"`
	FileList []*VersionConfigElement `xml:"fileList"`
}
type VersionConfigElement struct {
	FilePath string `xml:"filePath"`
	Sign     string `xml:"sign"`
}
